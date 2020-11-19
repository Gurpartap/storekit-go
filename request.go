package storekit

// ReceiptRequest is the JSON you submit with the request to the App Store.
//
// To receive a decoded receipt for validation, send a request with the encoded
// receipt data to the App Store. For auto-renewable subscriptions, include the
// app password and optionally an exclusion flag. Send this JSON data using the
// HTTP POST request method.
type ReceiptRequest struct {
	// ReceiptData contains the base64 encoded receipt data. Required.
	ReceiptData []byte `json:"receipt-data"`

	// Password is your app’s shared secret (a hexadecimal string). Required.
	//
	// Use this field only for receipts that contain auto-renewable subscriptions.
	Password string `json:"password"`

	// ExcludeOldTransactions is only used for app receipts that contain
	// auto-renewable or non-renewing subscriptions.
	//
	// Set this value to true for the response to include only the latest renewal transaction
	// for any subscriptions.
	ExcludeOldTransactions bool `json:"exclude-old-transactions,omitempty"`
}
