package storekit

// InAppOwnershipType is the relationship of the user with the family-shared
// purchase to which they have access.
//
// When family members benefit from a shared subscription, App Store updates
// their receipt to include the family-shared purchase. Use the value of
// in_app_ownership_type to understand whether a transaction belongs to the
// purchaser or a family member who benefits. This field appears in the App
// Store server notifications unified receipt
// (unified_receipt.Latest_receipt_info) and in transaction receipts
// (responseBody.Latest_receipt_info). For more information about Family
// Sharing, see Supporting Family Sharing in Your App.
type InAppOwnershipType string

const (
	InAppOwnershipTypeFamilyShared InAppOwnershipType = "FAMILY_SHARED"
	InAppOwnershipTypePurchased    InAppOwnershipType = "PURCHASED"
)

// LatestReceiptInfo is an array that contains all in-app purchase transactions.
// https://developer.apple.com/documentation/appstorereceipts/responsebody/latest_receipt_info
type LatestReceiptInfo struct {
	// The time Apple customer support canceled a transaction, in a date-time format
	// similar to the ISO 8601. This field is only present for refunded
	// transactions.
	CancellationDate string `json:"cancellation_date,omitempty"`

	// The time Apple customer support canceled a transaction, or the time an
	// auto-renewable subscription plan was upgraded, in UNIX epoch time format, in
	// milliseconds. This field is only present for refunded transactions. Use this
	// time format for processing dates.
	// information.
	CancellationDateMs int64 `json:"cancellation_date_ms,string,omitempty"`

	// The time Apple customer support canceled a transaction, in the Pacific Time
	// zone. This field is only present for refunded transactions.
	CancellationDatePst string `json:"cancellation_date_pst,omitempty"`

	// The reason for a refunded transaction. When a customer cancels a transaction,
	// the App Store gives them a refund and provides a value for this key. A value
	// of “1” indicates that the customer canceled their transaction due to an
	// actual or perceived issue within your app. A value of “0” indicates that the
	// transaction was canceled for another reason; for example, if the customer
	// made the purchase accidentally.
	CancellationReason string `json:"cancellation_reason,omitempty"`

	// The time a subscription expires or when it will renew, in a date-time format
	// similar to the ISO 8601.
	ExpiresDate string `json:"expires_date,omitempty"`

	// The time a subscription expires or when it will renew, in UNIX epoch time
	// format, in milliseconds. Use this time format for processing dates.
	ExpiresDateMs int64 `json:"expires_date_ms,string,omitempty"`

	// The time a subscription expires or when it will renew, in the Pacific Time
	// zone.
	ExpiresDatePst string `json:"expires_date_pst,omitempty"`

	// A value that indicates whether the user is the purchaser of the product, or
	// is a family member with access to the product through Family Sharing.
	// Possible Values:
	//  - FAMILY_SHARED: The transaction belongs to a family member who
	// benefits from service.
	//  - PURCHASED: The transaction belongs to the purchaser.
	InAppOwnershipType InAppOwnershipType `json:"in_app_ownership_type,omitempty"`

	// An indicator of whether an auto-renewable subscription is in the introductory
	// price period.
	IsInIntroOfferPeriod string `json:"is_in_intro_offer_period,omitempty"`

	// An indicator of whether a subscription is in the free trial period.
	IsTrialPeriod string `json:"is_trial_period,omitempty"`

	// An indicator that a subscription has been canceled due to an upgrade. This
	// field is only present for upgrade transactions.
	IsUpgraded string `json:"is_upgraded,omitempty"`

	// The reference name of a subscription offer that you configured in App Store
	// Connect. This field is present when a customer redeemed a subscription offer
	// code. For more information about offer codes, see [Set Up Offer Codes](https://help.apple.com/app-store-connect/#/dev6a098e4b1),
	// and [Implementing Offer Codes in Your App](https://developer.apple.com/documentation/storekit/in-app_purchase/subscriptions_and_offers/implementing_offer_codes_in_your_app).
	OfferCodeRefName string `json:"offer_code_ref_name,omitempty"`

	// The time of the original app purchase, in a date-time format similar to ISO
	// 8601.
	OriginalPurchaseDate string `json:"original_purchase_date,omitempty"`

	// The time of the original app purchase, in UNIX epoch time format, in
	// milliseconds. Use this time format for processing dates. For an
	// auto-renewable subscription, this value indicates the date of the
	// subscription's initial purchase. The original purchase date applies to all
	// product types and remains the same in all transactions for the same product
	// ID. This value corresponds to the original transaction’s transactionDate
	// property in StoreKit.
	OriginalPurchaseDateMs int64 `json:"original_purchase_date_ms,string,omitempty"`

	// The time of the original app purchase, in the Pacific Time zone.
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst,omitempty"`

	// The transaction identifier of the original purchase.
	OriginalTransactionId string `json:"original_transaction_id,omitempty"`

	// The unique identifier of the product purchased. You provide this value when
	// creating the product in App Store Connect, and it corresponds to the
	// productIdentifier property of the SKPayment object stored in the transaction's
	// payment property.
	ProductId string `json:"product_id,omitempty"`

	// The identifier of the subscription offer redeemed by the user.
	PromotionalOfferId string `json:"promotional_offer_id,omitempty"`

	// The time the App Store charged the user's account for a purchased or restored
	// product, or the time the App Store charged the user’s account for a subscription
	// purchase or renewal after a lapse, in a date-time format similar to ISO 8601.
	PurchaseDate string `json:"purchase_date,omitempty"`

	// For consumable, non-consumable, and non-renewing subscription products, the time
	// the App Store charged the user's account for a purchased or restored product, in
	// the UNIX epoch time format, in milliseconds. For auto-renewable subscriptions, the
	// time the App Store charged the user’s account for a subscription purchase or
	// renewal after a lapse, in the UNIX epoch time format, in milliseconds. Use this
	// time format for processing dates.
	PurchaseDateMs int64 `json:"purchase_date_ms,string,omitempty"`

	// The time the App Store charged the user's account for a purchased or restored
	// product, or the time the App Store charged the user’s account for a subscription
	// purchase or renewal after a lapse, in the Pacific Time zone.
	PurchaseDatePst string `json:"purchase_date_pst,omitempty"`

	// The number of consumable products purchased. This value corresponds to the
	// quantity property of the SKPayment object stored in the transaction's payment
	// property. The value is usually “1” unless modified with a mutable payment. The
	// maximum value is 10.
	Quantity int `json:"quantity,string,omitempty"`

	// The identifier of the subscription group to which the subscription belongs. The
	// value for this field is identical to the subscriptionGroupIdentifier property in
	// SKProduct.
	SubscriptionGroupIdentifier string `json:"subscription_group_identifier,omitempty"`

	// A unique identifier for a transaction such as a purchase, restore, or renewal.
	TransactionId string `json:"transaction_id,omitempty"`

	// A unique identifier for purchase events across devices, including
	// subscription-renewal events. This value is the primary key for identifying
	// subscription purchases.
	WebOrderLineItemId string `json:"web_order_line_item_id,omitempty"`
}
