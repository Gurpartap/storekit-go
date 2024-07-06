package storekit

// ReceiptResponseStatus is the status of the app receipt. The value for status
// is 0 if the receipt is valid, or a status code if there is an error. The
// status code reflects the status of the app receipt as a whole. For example,
// if you send a valid app receipt that contains an expired subscription, the
// response is 0 because the receipt is valid. Status codes 21100-21199 are
// various internal data access errors.
// https://developer.apple.com/documentation/appstorereceipts/status
type ReceiptResponseStatus int

const (
	// Undocumented but occurs.
	ReceiptResponseStatusUnknown ReceiptResponseStatus = -1

	// Receipt is valid.
	ReceiptResponseStatusOK ReceiptResponseStatus = 0

	// The request to the App Store was not made using the HTTP POST request method.
	ReceiptResponseStatusAppStoreCannotRead ReceiptResponseStatus = 21000

	// This status code is no longer sent by the App Store.
	ReceiptResponseStatusNoLongerSent ReceiptResponseStatus = 21001

	// The data in the receipt-data property was malformed or the service
	// experienced a temporary issue. Try again.
	ReceiptResponseStatusDataMalformed ReceiptResponseStatus = 21002

	// The receipt could not be authenticated.
	ReceiptResponseStatusNotAuthenticated ReceiptResponseStatus = 21003

	// The shared secret you provided does not match the shared secret on file for
	// your account.
	ReceiptResponseStatusSharedSecretDoesNotMatch ReceiptResponseStatus = 21004

	// The receipt server was temporarily unable to provide the receipt. Try again.
	ReceiptResponseStatusReceiptServerUnavailable ReceiptResponseStatus = 21005

	// This receipt is valid but the subscription has expired. When this status code
	// is returned to your server, the receipt data is also decoded and returned as
	// part of the response. Only returned for iOS 6-style transaction receipts for
	// auto-renewable subscriptions.
	ReceiptResponseStatusValidButSubscriptionExpired ReceiptResponseStatus = 21006

	// This receipt is from the test environment, but it was sent to the production
	// environment for verification.
	ReceiptResponseStatusSandboxReceiptSentToProduction ReceiptResponseStatus = 21007

	// This receipt is from the production environment, but it was sent to the test
	// environment for verification.
	ReceiptResponseStatusProductionReceiptSentToSandbox ReceiptResponseStatus = 21008

	// Internal data access error. Try again later.
	ReceiptResponseStatusBadAccess ReceiptResponseStatus = 21009

	// The user account cannot be found or has been deleted.
	ReceiptResponseStatusCouldNotBeAuthorized ReceiptResponseStatus = 21010

	// Status codes 21100-21199 are various internal data access errors.
)

// ReceiptResponse is the JSON data returned in the response from the App Store.
// https://developer.apple.com/documentation/appstorereceipts/responsebody
type ReceiptResponse struct {
	// The environment for which the receipt was generated.
	// Possible values: Sandbox, Production
	Environment string `json:"environment,omitempty"`

	// IsRetryable is an indicator that an error occurred during the request. A
	// value of 1 indicates a temporary issue; retry validation for this receipt at
	// a later time. A value of 0 indicates an unresolvable issue; do not retry
	// validation for this receipt. Only applicable to status codes 21100-21199.
	IsRetryable bool `json:"is-retryable,omitempty"`

	// The latest Base64 encoded app receipt. Only returned for receipts that
	// contain auto-renewable subscriptions.
	LatestReceipt []byte `json:"latest_receipt,omitempty"`

	// LatestReceiptInfo is an array that contains all in-app purchase transactions.
	// This excludes transactions for consumable products that have been marked as
	// finished by your app. Only returned for receipts that contain auto-renewable
	// subscriptions.
	LatestReceiptInfo []LatestReceiptInfo `json:"latest_receipt_info,omitempty"`

	// In the JSON file, an array where each element contains the pending renewal
	// information for each auto-renewable subscription identified by the
	// product_id. Only returned for app receipts that contain auto-renewable
	// subscriptions.
	PendingRenewalInfo []PendingRenewalInfo `json:"pending_renewal_info,omitempty"`

	// A JSON representation of the receipt that was sent for verification.
	Receipt Receipt `json:"receipt,omitempty"`

	// Either 0 if the receipt is valid, or a status code if there is an error. The
	// status code reflects the status of the app receipt as a whole.
	Status ReceiptResponseStatus `json:"status,omitempty"`
}
