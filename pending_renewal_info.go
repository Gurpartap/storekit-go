package storekit

// AutoRenewStatus is returned in the JSON response, in the
// responseBody.Pending_renewal_info array.
//
// The value for this field should not be interpreted as the customer’s
// subscription status. You can use this value to display an alternative
// subscription product in your app, such as a lower-level subscription plan to
// which the user can downgrade from their current plan. Consider presenting an
// attractive upgrade or downgrade offer in the same subscription group, if the
// auto_renew_status value is “0”. See Engineering Subscriptions from WWDC 2018
// for more information.
//
// https://developer.apple.com/documentation/appstorereceipts/auto_renew_status
type AutoRenewStatus string

const (
	// The customer has turned off automatic renewal for the subscription.
	AutoRenewStatusOff AutoRenewStatus = "0"

	// The subscription will renew at the end of the current subscription period.
	AutoRenewStatusOn AutoRenewStatus = "1"
)

// BillingRetryStatus indicates whether Apple is attempting to renew an expired
// subscription automatically.
//
// If the customer’s subscription failed to renew because the App Store was
// unable to complete the transaction, this value reflects whether the App Store
// is still trying to renew the subscription. The subscription retry flag is
// solely indicative of whether a subscription is in a billing retry state. Use
// this value in conjunction with expiration_intent, expires_date, and
// transaction_id for more insight.
//
// You can use this field to:
//
// - Determine that the user has been billed successfully, if this field has been
// removed and there is a new transaction with a future expires_date. Inform the
// user that there may be an issue with their billing information, if the value
// is “1”. For example, an expired credit card or insufficient balance could
// prevent this customer's account from being billed.
//
// - Implement a grace period to improve recovery, if the value is “1” and the
// expires_date is in the past. A grace period is free or limited subscription
// access while a subscriber is in a billing retry state. See Engineering
// Subscriptions from WWDC 2018 for more information.
//
// https://developer.apple.com/documentation/appstorereceipts/is_in_billing_retry_period
type BillingRetryStatus string

const (
	// The App Store has stopped attempting to renew the subscription.
	BillingRetryStatusStoppedAttemptingRenewal BillingRetryStatus = "0"

	// The App Store is attempting to renew the subscription.
	BillingRetryStatusAttemptingRenewal BillingRetryStatus = "1"
)

// ExpirationIntent is the reason a subscription expired.
type ExpirationIntent string

const (
	// The customer voluntarily canceled their subscription.
	ExpirationIntentVoluntarilyCancelled ExpirationIntent = "1"

	// Billing error for example, the customer's payment information was no longer
	// valid.
	ExpirationIntentBillingIssue ExpirationIntent = "2"

	// The customer did not agree to a recent price increase.
	ExpirationIntentDidNotAcceptPriceIncrease ExpirationIntent = "3"

	// The product was not available for purchase at the time of renewal.
	ExpirationIntentProductNotAvailable ExpirationIntent = "4"

	// Unknown error.
	ExpirationIntentUnknown ExpirationIntent = "5"
)

type PriceConsentStatus string

const (
	// The App Store hasn't yet requested consent from the customer.
	PriceConsentStatusNotRequested PriceConsentStatus = ""

	// The App Store is asking for the customer's consent, and hasn't received it.
	PriceConsentStatusAwaitingConsent PriceConsentStatus = "0"

	// The App Store has received customer's consent.
	PriceConsentStatusConsented PriceConsentStatus = "1"
)

// PendingRenewalInfo is an array of elements that refers to auto-renewable
// subscription renewals that are open or failed in the past.
// https://developer.apple.com/documentation/appstorereceipts/responsebody/pending_renewal_info
type PendingRenewalInfo struct {
	// The current renewal preference for the auto-renewable subscription. The value
	// for this key corresponds to the productIdentifier property of the product
	// that the customer’s subscription renews. This field is only present if the
	// user downgrades or crossgrades to a subscription of a different duration for
	// the subsequent subscription period.
	AutoRenewProductId string `json:"auto_renew_product_id,omitempty"`

	// The current renewal status for the auto-renewable subscription.
	AutoRenewStatus AutoRenewStatus `json:"auto_renew_status,omitempty"`

	// The reason a subscription expired. This field is only present for a receipt
	// that contains an expired auto-renewable subscription.
	ExpirationIntent ExpirationIntent `json:"expiration_intent,omitempty"`

	// The time at which the grace period for subscription renewals expires, in a
	// date-time format similar to the ISO 8601.
	GracePeriodExpiresDate string `json:"grace_period_expires_date,omitempty"`

	// The time at which the grace period for subscription renewals expires, in UNIX
	// epoch time format, in milliseconds. This key is only present for apps that
	// have Billing Grace Period enabled and when the user experiences a billing
	// error at the time of renewal. Use this time format for processing dates.
	GracePeriodExpiresDateMs int64 `json:"grace_period_expires_date_ms,string,omitempty"`

	// The time at which the grace period for subscription renewals expires, in the
	// Pacific Time zone.
	GracePeriodExpiresDatePst string `json:"grace_period_expires_date_pst,omitempty"`

	// A flag that indicates Apple is attempting to renew an expired subscription
	// automatically. This field is only present if an auto-renewable subscription
	// is in the billing retry state.
	IsInBillingRetryPeriod BillingRetryStatus `json:"is_in_billing_retry_period,omitempty"`

	// The reference name of a subscription offer that you configured in App Store
	// Connect. This field is present when a customer redeemed a subscription offer
	// code. For more information about offer codes, see [Set Up Offer Codes](https://help.apple.com/app-store-connect/#/dev6a098e4b1),
	// and [Implementing Offer Codes in Your App](https://developer.apple.com/documentation/storekit/in-app_purchase/subscriptions_and_offers/implementing_offer_codes_in_your_app).
	OfferCodeRefName string `json:"offer_code_ref_name,omitempty"`

	// The transaction identifier of the original purchase.
	OriginalTransactionId string `json:"original_transaction_id,omitempty"`

	// The price consent status for a subscription price increase. This field is
	// only present if the customer was notified of the price increase. The default
	// value is "0" and changes to "1" if the customer consents.
	PriceConsentStatus PriceConsentStatus `json:"price_consent_status,omitempty"`

	// The unique identifier of the product purchased. You provide this value when
	// creating the product in App Store Connect, and it corresponds to the
	// productIdentifier property of the SKPayment object stored in the
	// transaction's payment property.
	ProductId string `json:"product_id,omitempty"`
}
