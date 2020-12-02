package storekit

// NotificationType is the type that describes the in-app purchase event for
// which the App Store sent the notification.
//
// You receive and can react to server notifications in real time for the
// subscription and refund events that these notification type values describe.
// The notification_type appears in the responseBody.
//
// https://developer.apple.com/documentation/appstoreservernotifications/notification_type
type NotificationType string

const (
	// Indicates that either Apple customer support canceled the subscription or the
	// user upgraded their subscription. The cancellation_date key contains the date
	// and time of the change.
	NotificationTypeCancel NotificationType = "CANCEL"

	// Indicates that the customer made a change in their subscription plan that
	// takes effect at the next renewal. The currently active plan isn’t affected.
	NotificationTypeDidChangeRenewalPref NotificationType = "DID_CHANGE_RENEWAL_PREF"

	// Indicates a change in the subscription renewal status. In the JSON response,
	// check auto_renew_status_change_date_ms to know the date and time of the last
	// status update. Check auto_renew_status to know the current renewal status.
	NotificationTypeDidChangeRenewalStatus NotificationType = "DID_CHANGE_RENEWAL_STATUS"

	// Indicates a subscription that failed to renew due to a billing issue. Check
	// is_in_billing_retry_period to know the current retry status of the
	// subscription. Check grace_period_expires_date to know the new service
	// expiration date if the subscription is in a billing grace period.
	NotificationTypeDidFailToRenew NotificationType = "DID_FAIL_TO_RENEW"

	// Indicates a successful automatic renewal of an expired subscription that
	// failed to renew in the past. Check expires_date to determine the next renewal
	// date and time.
	NotificationTypeDidRecover NotificationType = "DID_RECOVER"

	// Indicates that a customer’s subscription has successfully auto-renewed for a
	// new transaction period.
	NotificationTypeDidRenew NotificationType = "DID_RENEW"

	// Occurs at the user’s initial purchase of the subscription. Store
	// latest_receipt on your server as a token to verify the user’s subscription
	// status at any time by validating it with the App Store.
	NotificationTypeInitialBuy NotificationType = "INITIAL_BUY"

	// Indicates the customer renewed a subscription interactively, either by using
	// your app’s interface, or on the App Store in the account’s Subscriptions
	// settings. Make service available immediately.
	NotificationTypeInteractiveRenewal NotificationType = "INTERACTIVE_RENEWAL"

	// Indicates that App Store has started asking the customer to consent to your
	// app’s subscription price increase. In the
	// unified_receipt.Pending_renewal_info object, the price_consent_status value
	// is 0, indicating that App Store is asking for the customer’s consent, and
	// hasn't received it. The subscription won’t auto-renew unless the user agrees
	// to the new price. When the customer agrees to the price increase, the system
	// sets price_consent_status to 1. Check the receipt using verifyReceipt to view
	// the updated price-consent status.
	NotificationTypePriceIncreaseConsent NotificationType = "PRICE_INCREASE_CONSENT"

	// Indicates that App Store successfully refunded a transaction. The
	// cancellation_date_ms contains the timestamp of the refunded transaction. The
	// original_transaction_id and product_id identify the original transaction and
	// product. The cancellation_reason contains the reason.
	NotificationTypeRefund NotificationType = "REFUND"
)

// UnifiedReceipt is an object that contains information about the most-recent,
// in-app purchase transactions for the app.
//
// https://developer.apple.com/documentation/appstoreservernotifications/unified_receipt
type UnifiedReceipt struct {
	// The environment for which App Store generated the receipt.
	// Possible values: Sandbox, Production
	Environment string `json:"environment,omitempty"`

	// The latest Base64-encoded app receipt.
	LatestReceipt []byte `json:"latest_receipt,omitempty"`

	// An array that contains the latest 100 in-app purchase transactions of the
	// decoded value in latest_receipt. This array excludes transactions for
	// consumable products your app has marked as finished. The contents of this
	// array are identical to those in LatestReceiptInfo in the
	// verifyReceipt endpoint response for receipt validation.
	LatestReceiptInfo []LatestReceiptInfo `json:"latest_receipt_info,omitempty"`

	// An array where each element contains the pending renewal information for each
	// auto-renewable subscription identified in product_id. The contents of this
	// array are identical to those in PendingRenewalInfo in the
	// verifyReceipt endpoint response for receipt validation.
	PendingRenewalInfo []PendingRenewalInfo `json:"pending_renewal_info,omitempty"`

	// The status code, where 0 indicates that the notification is valid.
	Status int `json:"status,omitempty"`
}

// Notification is the JSON data sent in the server notification from the App
// Store.
//
// Use the information in the response body to react quickly to changes in your
// users’ subscription states. The fields available in any one notification sent
// to your server are dependent on the notification_type, which indicates the
// event that triggered the notification.
//
// https://developer.apple.com/documentation/appstoreservernotifications/responsebody
type Notification struct {
	// An identifier that App Store Connect generates and the App Store uses to
	// uniquely identify the auto-renewable subscription that the user’s
	// subscription renews. Treat this value as a 64-bit integer.
	AutoRenewAdamId string `json:"auto_renew_adam_id,omitempty"`

	// The product identifier of the auto-renewable subscription that the user’s
	// subscription renews.
	AutoRenewProductId string `json:"auto_renew_product_id,omitempty"`

	// The current renewal status for an auto-renewable subscription product. Note
	// that these values are different from those of the auto_renew_status in the
	// receipt.
	// Possible values: true, false
	AutoRenewStatus string `json:"auto_renew_status,omitempty"`

	// The time at which the user turned on or off the renewal status for an
	// auto-renewable subscription, in a date-time format similar to the ISO 8601
	// standard.
	AutoRenewStatusChangeDate string `json:"auto_renew_status_change_date,omitempty"`

	// The time at which the user turned on or off the renewal status for an
	// auto-renewable subscription, in UNIX epoch time format, in milliseconds. Use
	// this time format to process dates.
	AutoRenewStatusChangeDateMs string `json:"auto_renew_status_change_date_ms,omitempty"`

	// The time at which the user turned on or off the renewal status for an
	// auto-renewable subscription, in the Pacific time zone.
	AutoRenewStatusChangeDatePst string `json:"auto_renew_status_change_date_pst,omitempty"`

	// The environment for which App Store generated the receipt.
	// Possible values: Sandbox, PROD
	Environment string `json:"environment,omitempty"`

	// The reason a subscription expired. This field is only present for an expired
	// auto-renewable subscription. See expiration_intent for more information.
	ExpirationIntent ExpirationIntent `json:"expiration_intent,omitempty"`

	// The subscription event that triggered the notification.
	NotificationType NotificationType `json:"notification_type,omitempty"`

	// The same value as the shared secret you submit in the password field of the
	// requestBody when validating receipts.
	Password string `json:"password,omitempty"`

	// An object that contains information about the most-recent, in-app purchase
	// transactions for the app.
	UnifiedReceipt UnifiedReceipt `json:"unified_receipt,omitempty"`

	// A string that contains the app bundle ID.
	Bid string `json:"bid,omitempty"`

	// A string that contains the app bundle version.
	Bvrs string `json:"bvrs,omitempty"`
}
