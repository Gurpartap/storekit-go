package storekit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"unicode"
)

const (
	sandboxReceiptVerificationURL    = "https://sandbox.itunes.apple.com/verifyReceipt"
	productionReceiptVerificationURL = "https://buy.itunes.apple.com/verifyReceipt"
)

type client struct {
	verificationURL    string
	autofixEnvironment bool
}

// NewVerificationClient defaults to production verification URL with auto fix
// enabled.
//
// Auto fix automatically handles the incompatible receipt environment error. It
// subsequently gets disabled after the first attempt to avoid unexpected
// looping.
func NewVerificationClient() *client {
	return &client{
		verificationURL:    productionReceiptVerificationURL,
		autofixEnvironment: true,
	}
}

// OnProductionEnv sets the client to use sandbox URL for verification.
func (c *client) OnSandboxEnv() *client {
	c.verificationURL = sandboxReceiptVerificationURL
	return c
}

// OnProductionEnv sets the client to use production URL for verification.
func (c *client) OnProductionEnv() *client {
	c.verificationURL = productionReceiptVerificationURL
	return c
}

// WithoutEnvAutoFix disables automatic handling of incompatible receipt
// environment error.
func (c *client) WithoutEnvAutoFix() *client {
	c.autofixEnvironment = false
	return c
}

func (c *client) Verify(ctx context.Context, req *ReceiptRequest) ([]byte, *ReceiptResponse, error) {
post:
	body, err := c.post(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	resp := &ReceiptResponse{}
	err = json.Unmarshal(
		bytes.Map(func(r rune) rune {
			if unicode.IsControl(r) {
				return -1
			}
			return r
		}, body),
		resp,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("could not unmarshal app store response: %w", err)
	}

	if c.autofixEnvironment {
		// Auto fix but only once.
		c.autofixEnvironment = false

		switch resp.Status {
		case ReceiptResponseStatusSandboxReceiptSentToProduction:
			// On a 21007 status, retry the request in the sandbox environment (only if the
			// current environment is production – to avoid unexpected loop).
			//
			// These are receipts from Apple review team.
			if c.isProduction() {
				c.verificationURL = sandboxReceiptVerificationURL
				goto post
			}
		case ReceiptResponseStatusProductionReceiptSentToSandbox:
			// On a 21008 status, retry the request in the production environment (only if
			// the current environment is sandbox – to avoid unexpected loop).
			if c.isSandbox() {
				c.verificationURL = productionReceiptVerificationURL
				goto post
			}
		default:
			// TODO: Retry at least once when an App Store internal error occurs here:
			// 	if resp.Status >= 21100 && resp.Status <= 21199 {
			// 		if resp.IsRetryable {
			// 			goto post
			// 		}
			// 	}
			break
		}
	}

	return body, resp, nil
}

func (c *client) post(ctx context.Context, receiptRequest *ReceiptRequest) ([]byte, error) {
	// Prepare request:

	reqJSON, err := json.Marshal(receiptRequest)
	if err != nil {
		return nil, fmt.Errorf("could not marshal receipt request: %w", err)
	}

	// Dial the App Store server:

	buf := bytes.NewReader(reqJSON)

	req, err := http.NewRequest(http.MethodPost, c.verificationURL, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO: Handle this error (and probably retry at least once):
		//       Post https://sandbox.itunes.apple.com/verifyReceipt: read tcp 10.1.11.101:36372->17.154.66.159:443: read: connection reset by peer
		return nil, fmt.Errorf("could not connect to app store server: %w", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("app store http error (%s)", r.Status)
	}

	// Parse response:

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read app store response: %w", err)
	}

	return body, nil
}

func (c *client) isSandbox() bool {
	return c.verificationURL == sandboxReceiptVerificationURL
}

func (c *client) isProduction() bool {
	return c.verificationURL == productionReceiptVerificationURL
}
