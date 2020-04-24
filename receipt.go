package storekit

type Receipt struct {
	// BundleIdentifier is the app’s bundle identifier.
	//
	// This corresponds to the value of CFBundleIdentifier in the Info.plist
	// file. Use this value to validate if the receipt was indeed generated for
	// your app.
	BundleIdentifier string `json:"bundle_id,omitempty"`

	// AppVersion is the app’s version number.
	//
	// This corresponds to the value of CFBundleVersion (in iOS) or
	// CFBundleShortVersionString (in macOS) in the Info.plist.
	AppVersion string `json:"application_version,omitempty"`

	// InAppPurchaseReceipt is the receipt for an in-app purchase.
	//
	// SET of in-app purchase receipt attributes
	// array of in-app purchase receipts
	//
	// In the JSON file, the value of this key is an array containing all
	// in-app purchase receipts based on the in-app purchase transactions
	// present in the input base-64 receipt-data. For receipts containing
	// auto-renewable subscriptions, check the value of the latest_receipt_info
	// key to get the status of the most recent renewal.
	//
	//	In the ASN.1 file, there are multiple fields that all have type 17,
	// each of which contains a single in-app purchase receipt.
	//
	// 	Note: An empty array is a valid receipt.
	// 	The in-app purchase receipt for a consumable product is added to the
	// receipt when the purchase is made. It is kept in the receipt until your
	// app finishes that transaction. After that point, it is removed from the
	// receipt the next time the receipt is updated - for example, when the
	// user makes another purchase or if your app explicitly refreshes the
	// receipt.
	//
	// 	The in-app purchase receipt for a non-consumable product,
	// auto-renewable subscription, non-renewing subscription, or free
	// subscription remains in the receipt indefinitely.
	InAppPurchaseReceipt []InAppPurchaseReceipt `json:"in_app,omitempty"`

	// OriginalApplicationVersion is the version of the app that was originally
	// purchased.
	//
	// This corresponds to the value of CFBundleVersion (in iOS) or
	// CFBundleShortVersionString (in macOS) in the Info.plist file when the
	// purchase was originally made.
	//
	// In the sandbox environment, the value of this field is always “1.0”.
	OriginalApplicationVersion string `json:"original_application_version,omitempty"`

	// ReceiptCreationDate is the date when the app receipt was created.
	//
	// When validating a receipt, use this date to validate the receipt’s
	// signature.
	ReceiptCreationDate string `json:"creation_date,omitempty"` // time.Time

	// ReceiptExpirationDate is the date that the app receipt expires.
	//
	// This key is present only for apps purchased through the Volume Purchase
	// Program. If this key is not present, the receipt does not expire.
	//
	// When validating a receipt, compare this date to the current date to
	// determine whether the receipt is expired. Do not try to use this date to
	// calculate any other information, such as the time remaining before
	// expiration.
	ReceiptExpirationDate string `json:"expiration_date,omitempty"` // time.Time
}

type ReceiptRequest struct {
	// ReceiptData contains the base64 encoded receipt data.
	ReceiptData []byte `json:"receipt-data"`

	// Password is your app’s shared secret (a hexadecimal string).
	//
	// Only used for receipts that contain auto-renewable subscriptions.
	Password string `json:"password,omitempty"`

	// ExcludeOldTransactions is only used for iOS7 style app receipts that
	// contain auto-renewable or non-renewing subscriptions.
	//
	// If value is true, response includes only the latest renewal transaction
	// for any subscriptions.
	ExcludeOldTransactions bool `json:"exclude-old-transactions"`
}

const (
	ReceiptResponseStatusUnknown                        int = -1 // undocumented but -1 happens
	ReceiptResponseStatusOK                                 = 0
	ReceiptResponseStatusAppStoreCannotRead                 = 21000
	ReceiptResponseStatusDataMalformed                      = 21002
	ReceiptResponseStatusNotAuthenticated                   = 21003
	ReceiptResponseStatusSharedSecretDoesNotMatch           = 21004
	ReceiptResponseStatusReceiptServerUnavailable           = 21005
	ReceiptResponseStatusValidButSubscriptionExpired        = 21006
	ReceiptResponseStatusSandboxReceiptSentToProduction     = 21007
	ReceiptResponseStatusProductionReceiptSentToSandbox     = 21008
	ReceiptResponseStatusCouldNotBeAuthorized               = 21010
)

type ReceiptResponse struct {
	// Status is either 0 if the receipt is valid, or one of the error codes
	// listed.
	//
	// For iOS 6 style transaction receipts, the status code reflects the
	// status of the specific transaction’s receipt.
	//
	// For iOS 7 style app receipts, the status code is reflects the status of
	// the app receipt as a whole. For example, if you send a valid app receipt
	// that contains an expired subscription, the response is 0 because the
	// receipt as a whole is valid.
	//
	// 21000: The App Store could not read the JSON object you provided.
	// 21002: The data in the receipt-data property was malformed or missing.
	// 21003: The receipt could not be authenticated.
	// 21004: The shared secret you provided does not match the shared secret
	//        on file for your account.
	// 21005: The receipt server is not currently available.
	// 21006: This receipt is valid but the subscription has expired. When this
	//        status code is returned to your server, the receipt data is also
	//        decoded and returned as part of the response. Only returned for
	//        iOS 6 style transaction receipts for auto-renewable
	//        subscriptions.
	// 21007: This receipt is from the test environment, but it was sent to the
	//        production environment for verification. Send it to the test
	//        environment instead.
	// 21008: This receipt is from the production environment, but it was sent
	//        to the test environment for verification. Send it to the
	//        production environment instead.
	// 21010: This receipt could not be authorized. Treat this the same as if a
	//        purchase was never made.
	// 21100-21199: Internal data access error.
	Status int `json:"status,omitempty"`

	// Undocumented in Apple docs but useful when using the same client for
	// both production and sandbox use.
	Environment string `json:"environment,omitempty"`

	// Receipt that was sent for verification.
	Receipt Receipt `json:"receipt,omitempty"`

	// LatestReceipt is only returned for receipts containing auto-renewable
	// subscriptions.
	//
	// For iOS 6 style transaction receipts, this is the base-64 encoded
	// receipt for the most recent renewal.
	//
	// For iOS 7 style app receipts, this is the latest base-64 encoded app
	// receipt.
	LatestReceipt []byte `json:"latest_receipt,omitempty"`

	// LatestReceiptInfo is only returned for receipts containing
	// auto-renewable subscriptions.
	//
	// For iOS 6 style transaction receipts, this is the JSON representation of
	// the receipt for the most recent renewal.
	//
	// For iOS 7 style app receipts, the value of this key is an array
	// containing all in-app purchase transactions. This excludes transactions
	// for a consumable product that have been marked as finished by your app.
	LatestReceiptInfo []InAppPurchaseReceipt `json:"latest_receipt_info,omitempty"`

	// LatestExpiredReceiptInfo is only returned for iOS 6 style transaction
	// receipts, for an auto-renewable subscription. The JSON representation of
	// the receipt for the expired subscription.
	LatestExpiredReceiptInfo InAppPurchaseReceipt `json:"latest_expired_receipt_info,omitempty"`

	// PendingRenewalInfo is only returned for iOS 7 style app receipts
	// containing auto-renewable subscriptions. In the JSON file, the value of
	// this key is an array where each element contains the pending renewal
	// information for each auto-renewable subscription identified by the
	// Product Identifier. A pending renewal may refer to a renewal that is
	// scheduled in the future or a renewal that failed in the past for some
	// reason.
	PendingRenewalInfo []InAppPurchaseReceipt `json:"pending_renewal_info,omitempty"`

	// IsRetryable provides retry validation for this receipt. Only applicable
	// to status codes 21100-21199
	IsRetryable bool `json:"is-retryable,omitempty"`
}
