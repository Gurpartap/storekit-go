package storekit

type InAppPurchaseReceipt struct {
	// Quantity is the number of items purchased.
	//
	// This value corresponds to the quantity property of the SKPayment object
	// stored in the transaction’s payment property.
	Quantity int `json:"quantity,string,omitempty"`

	// ProductIdentifier is the product identifier of the item that was
	// purchased.
	//
	// This value corresponds to the productIdentifier property of the
	// SKPayment object stored in the transaction’s payment property.
	ProductIdentifier string `json:"product_id,omitempty"`

	// TransactionIdentifier is the transaction identifier of the item that was
	// purchased.
	//
	// This value corresponds to the transaction’s transactionIdentifier
	// property.
	//
	// For a transaction that restores a previous transaction, this value is
	// different from the transaction identifier of the original purchase
	// transaction. In an auto-renewable subscription receipt, a new value for
	// the transaction identifier is generated every time the subscription
	// automatically renews or is restored on a new device.
	TransactionIdentifier string `json:"transaction_id,omitempty"`

	// OriginalTransactionIdentifier is the transaction identifier of the
	// original transaction for a transaction that restores a previous
	// transaction. Otherwise, identical to the transaction identifier.
	//
	// This value corresponds to the original transaction’s
	// transactionIdentifier property.
	//
	// This value is the same for all receipts that have been generated for a
	// specific subscription. This value is useful for relating together
	// multiple iOS 6 style transaction receipts for the same individual
	// customer’s subscription.
	OriginalTransactionIdentifier string `json:"original_transaction_id,omitempty"`

	// PurchaseDate is the date and time that the item was purchased.
	//
	// This value corresponds to the transaction’s transactionDate property.
	//
	// For a transaction that restores a previous transaction, the purchase
	// date is the same as the original purchase date. Use Original Purchase
	// Date to get the date of the original transaction.
	//
	// In an auto-renewable subscription receipt, the purchase date is the date
	// when the subscription was either purchased or renewed (with or without a
	// lapse). For an automatic renewal that occurs on the expiration date of
	// the current period, the purchase date is the start date of the next
	// period, which is identical to the end date of the current period.
	PurchaseDate   string `json:"purchase_date,omitempty"`           // time.Time
	PurchaseDateMs int64  `json:"purchase_date_ms,string,omitempty"` // time.Time

	// OriginalPurchaseDate is the date of the original transaction for a
	// transaction that restores a previous transaction.
	//
	// This value corresponds to the original transaction’s transactionDate
	// property.
	//
	// In an auto-renewable subscription receipt, this indicates the beginning
	// of the subscription period, even if the subscription has been renewed.
	OriginalPurchaseDate   string `json:"original_purchase_date,omitempty"`           // time.Time
	OriginalPurchaseDateMs int64  `json:"original_purchase_date_ms,string,omitempty"` // time.Time

	// SubscriptionExpirationDate is the expiration date for the subscription,
	// expressed as the number of milliseconds since
	// January 1, 1970, 00:00:00 GMT.
	//
	// This key is only present for auto-renewable subscription receipts. Use
	// this value to identify the date when the subscription will renew or
	// expire, to determine if a customer should have access to content or
	// service. After validating the latest receipt, if the subscription
	// expiration date for the latest renewal transaction is a past date, it is
	// safe to assume that the subscription has expired.
	SubscriptionExpirationDate   string `json:"expires_date,omitempty"`           // time.Time
	SubscriptionExpirationDateMs int64  `json:"expires_date_ms,string,omitempty"` // time.Time

	// SubscriptionExpirationIntent is for an expired subscription, the reason
	// for the subscription expiration.
	//
	// “1” - Customer canceled their subscription.
	//
	// “2” - Billing error; for example customer’s payment information was no
	//       longer valid.
	//
	// “3” - Customer did not agree to a recent price increase.
	//
	// “4” - Product was not available for purchase at the time of renewal.
	//
	// “5” - Unknown error.
	//
	// This key is only present for a receipt containing an expired
	// auto-renewable subscription. You can use this value to decide whether to
	// display appropriate messaging in your app for customers to resubscribe.
	SubscriptionExpirationIntent int `json:"expiration_intent,string,omitempty"`

	// SubscriptionRetryFlag is for an expired subscription, whether or not
	// Apple is still attempting to automatically renew the subscription.
	//
	// “1” - App Store is still attempting to renew the subscription.
	//
	// “0” - App Store has stopped attempting to renew the subscription.
	//
	// This key is only present for auto-renewable subscription receipts. If
	// the customer’s subscription failed to renew because the App Store was
	// unable to complete the transaction, this value will reflect whether or
	// not the App Store is still trying to renew the subscription.
	SubscriptionRetryFlag int `json:"is_in_billing_retry_period,string,omitempty"`

	// SubscriptionTrialPeriod is for a subscription, whether or not it is in
	// the free trial period.
	//
	// This key is only present for auto-renewable subscription receipts. The
	// value for this key is "true" if the customer’s subscription is currently
	// in the free trial period, or "false" if not.
	//
	// Note: If a previous subscription period in the receipt has the value
	// “true” for either the is_trial_period or the is_in_intro_offer_period
	// key, the user is not eligible for a free trial or introductory price
	// within that subscription group.
	SubscriptionTrialPeriod string `json:"is_trial_period,omitempty"`

	// SubscriptionIntroductoryPricePeriod is for an auto-renewable
	// subscription, whether or not it is in the introductory price period.
	//
	// This key is only present for auto-renewable subscription receipts. The
	// value for this key is "true" if the customer’s subscription is currently
	// in an introductory price period, or "false" if not.
	//
	// Note: If a previous subscription period in the receipt has the value
	// “true” for either the is_trial_period or the is_in_intro_offer_period
	// key, the user is not eligible for a free trial or introductory price
	// within that subscription group.
	SubscriptionIntroductoryPricePeriod string `json:"is_in_intro_offer_period,omitempty"`

	// CancellationDate is for a transaction that was canceled by Apple
	// customer support, the time and date of the cancellation. For an
	// auto-renewable subscription plan that was upgraded, the time and date of
	// the upgrade transaction.
	//
	// Treat a canceled receipt the same as if no purchase had ever been made.
	//
	// Note: A canceled in-app purchase remains in the receipt indefinitely.
	// Only applicable if the refund was for a non-consumable product, an
	// auto-renewable subscription, a non-renewing subscription, or for a free
	// subscription.
	CancellationDate   string `json:"cancellation_date,omitempty"`           // time.Time
	CancellationDateMs int64  `json:"cancellation_date_ms,string,omitempty"` // time.Time

	// CancellationReason is for a transaction that was cancelled, the reason
	// for cancellation.
	//
	// “1” - Customer canceled their transaction due to an actual or perceived
	// issue within your app.
	//
	// “0” - Transaction was canceled for another reason, for example, if the
	// customer made the purchase accidentally.
	//
	// Use this value along with the cancellation date to identify possible
	// issues in your app that may lead customers to contact Apple customer
	// support.
	CancellationReason int `json:"cancellation_reason,string,omitempty"`

	// AppItemID is a string that the App Store uses to uniquely identify the
	// application that created the transaction.
	//
	// If your server supports multiple applications, you can use this value
	// to differentiate between them.
	//
	// Apps are assigned an identifier only in the production environment, so
	// this key is not present for receipts created in the test environment.
	//
	// This field is not present for Mac apps.
	AppItemID string `json:"app_item_id,omitempty"`

	// ExternalVersionIdentifier is an arbitrary number that uniquely
	// identifies a revision of your application.
	//
	// This key is not present for receipts created in the test environment.
	// Use this value to identify the version of the app that the customer
	// bought.
	ExternalVersionIdentifier string `json:"version_external_identifier,omitempty"`

	// WebOrderLineItemID is the primary key for identifying subscription
	// purchases.
	//
	// This value is a unique ID that identifies purchase events across
	// devices, including subscription renewal purchase events.
	WebOrderLineItemID string `json:"web_order_line_item_id,omitempty"`

	// SubscriptionAutoRenewStatus is the current renewal status for the
	// auto-renewable subscription.
	//
	// “1” - Subscription will renew at the end of the current subscription
	// period.
	//
	// “0” - Customer has turned off automatic renewal for their subscription.
	//
	// This key is only present for auto-renewable subscription receipts, for
	// active or expired subscriptions. The value for this key should not be
	// interpreted as the customer’s subscription status. You can use this
	// value to display an alternative subscription product in your app, for
	// example, a lower level subscription plan that the customer can downgrade
	// to from their current plan.
	SubscriptionAutoRenewStatus int `json:"auto_renew_status,string,omitempty"`

	// SubscriptionAutoRenewPreference is the current renewal preference for
	// the auto-renewable subscription.
	//
	// This key is only present for auto-renewable subscription receipts. The
	// value for this key corresponds to the productIdentifier property of the
	// product that the customer’s subscription renews. You can use this value
	// to present an alternative service level to the customer before the
	// current subscription period ends.
	SubscriptionAutoRenewPreference string `json:"auto_renew_product_id,omitempty"`

	// SubscriptionPriceConsentStatus is the current price consent status for a
	// subscription price increase.
	//
	// “1” - Customer has agreed to the price increase. Subscription will renew
	// at the higher price.
	//
	// “0” - Customer has not taken action regarding the increased price.
	// Subscription expires if the customer takes no action before the renewal
	// date.
	//
	// This key is only present for auto-renewable subscription receipts if the
	// subscription price was increased without keeping the existing price for
	// active subscribers. You can use this value to track customer adoption of
	// the new price and take appropriate action.
	SubscriptionPriceConsentStatus int `json:"price_consent_status,string,omitempty"`
}
