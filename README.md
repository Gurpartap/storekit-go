# storekit-go

[![GoDoc](https://godoc.org/github.com/Gurpartap/storekit-go?status.svg)](https://godoc.org/github.com/Gurpartap/storekit-go)

Use this for verifying App Store receipts.

- [x] Battle proven technology
- [x] Blockchain free

See [GoDoc](https://godoc.org/github.com/Gurpartap/storekit-go) for detailed API response reference.

## Usage example (auto-renewing subscriptions)

```go
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/Gurpartap/storekit-go"
)

func main() {
	// Get it from https://AppsToReconnect.apple.com 🤯
	appStoreSharedSecret = os.GetEnv("APP_STORE_SHARED_SECRET")

	// Your own userID
	userID := "12345"

	// Input coming from either user device or subscription notifications
	// webhook
	receiptData := []byte("...")

	err := verifyAndSave(appStoreSharedSecret, userID, receiptData)
	if err != nil {
		fmt.Println("could not verify receipt:", err)
	}
}

func verifyAndSave(appStoreSharedSecret, userID string, receiptData []byte) error {
	// Use .OnProductionEnv() when deploying
	//
	// storekit-go automatically retries sandbox server upon incompatible
	// environment error. This is necessary because App Store Reviewer's
	// purchase requests go through the sandbox server instead of production.
	//
	// Use .WithoutEnvAutoFix() to disable automatic env switching and retrying
	// (not recommended on production)
	client := storekit.NewVerificationClient().OnSandboxEnv()

	// respBody is raw bytes of response, useful for storing, auditing, and
	// for future verification checks. resp is the same parsed and mapped to a
	// struct.
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	respBody, resp, err := client.Verify(ctx, &storekit.ReceiptRequest{
		ReceiptData:            receiptData,
		Password:               appStoreSharedSecret,
		ExcludeOldTransactions: true,
	})
	if err != nil {
		return err // code: internal error
	}

	if resp.Status != 0 {
		return errors.New(
			fmt.Sprintf("receipt rejected by App Store with status = %d", resp.Status),
		) // code: permission denied
	}

	// If receipt does not contain any active subscription info it is probably
	// a fraudulent attempt at activating subscription from a jailbroken
	// device.
	if len(resp.LatestReceiptInfo) == 0 {
		// keep it 🤫 that we know what's going on
		return errors.New("unknown error") // code: internal (instead of invalid argument)
	}

	// resp.LatestReceiptInfo works for me...
	// ... but, alternatively (as Apple devs also recommend) you can loop over
	// resp.Receipt.InAppPurchaseReceipt, and filter for the receipt with the
	// highest expiresAtMs to find the appropriate latest subscription
	// (not shown in this example).
	for _, latestReceiptInfo := range resp.LatestReceiptInfo {
		productID := latestReceiptInfo.ProductIdentifier
		expiresAtMs := latestReceiptInfo.SubscriptionExpirationDateMs
		// cancelledAtStr := latestReceiptInfo.CancellationDate

		// defensively check for necessary data ...
		// ... because StoreKit API responses can be a bit adventurous
		if productID == "" {
			return errors.New("missing product_id in the latest receipt info") // code: internal error
		}
		if expiresAtMs == 0 {
			return errors.New("missing expiry date in latest receipt info") // code: internal error
		}

		expiresAt := time.Unix(0, expiresAtMs*1000000)

		fmt.Printf(
			"userID = %s has subscribed for product_id = %s which expires_at = %s",
			userID,
			productID,
			expiresAt,
		)

		// ✅ Save or return productID, expiresAt, cancelledAt, respBody
	}
}

```
