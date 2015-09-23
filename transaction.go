package braintree

import (
	"github.com/lionelbarrow/braintree-go/nullable"
	"time"
)

type TransactionStatusHistory struct {
	Timestamp *time.Time `xml:"timestamp,omitempty"`
	Status    string     `xml:"status"`
	Amount    *Decimal   `xml:"amount"`
	User      string     `xml:"user,omitempty"`
	Source    string     `xml:"transaction-source,omitempty"`
}

type SubscriptionBillingPeriod struct {
	StartsAt string `xml:"billing-period-start-date,omitempty"`
	EndsAt   string `xml:"billing-period-end-date,omitempty"`
}

type Transaction struct {
	XMLName               string                     `xml:"transaction"`
	Id                    string                     `xml:"id,omitempty"`
	CustomerID            string                     `xml:"customer-id,omitempty"`
	Status                string                     `xml:"status,omitempty"`
	Type                  string                     `xml:"type,omitempty"`
	Amount                *Decimal                   `xml:"amount"`
	OrderId               string                     `xml:"order-id,omitempty"`
	PaymentMethodToken    string                     `xml:"payment-method-token,omitempty"`
	PaymentMethodNonce    string                     `xml:"payment-method-nonce,omitempty"`
	MerchantAccountId     string                     `xml:"merchant-account-id,omitempty"`
	PlanId                string                     `xml:"plan-id,omitempty"`
	CreditCard            *CreditCard                `xml:"credit-card,omitempty"`
	Customer              *Customer                  `xml:"customer,omitempty"`
	BillingAddress        *Address                   `xml:"billing,omitempty"`
	ShippingAddress       *Address                   `xml:"shipping,omitempty"`
	Options               *TransactionOptions        `xml:"options,omitempty"`
	ServiceFeeAmount      *Decimal                   `xml:"service-fee-amount,attr,omitempty"`
	CreatedAt             *time.Time                 `xml:"created-at,omitempty"`
	UpdatedAt             *time.Time                 `xml:"updated-at,omitempty"`
	DisbursementDetails   *DisbursementDetails       `xml:"disbursement-details,omitempty"`
	RefundIds             *[]string                  `xml:"refund-ids>item,omitempty"`
	RefundedTransactionId *string                    `xml:"refunded-transaction-id,omitempty"`
	StatusHistory         []TransactionStatusHistory `xml:"status-history>status-event,omitempty"`
	SubscriptionID        string                     `xml:"subscription-id,omitempty"`
	BillingPeriod         SubscriptionBillingPeriod  `xml:"subscription,omitempty"`
	AddOnList             []AddOn                    `xml:"add-ons>add-on,omitempty"`
	DiscountList          []Discount                 `xml:"discounts>discount,omitempty"`
}

// TODO: not all transaction fields are implemented yet, here are the missing fields (add on demand)
//
// <transaction>
//   <currency-iso-code>USD</currency-iso-code>
//   <settlement-batch-id>2013-10-08_49grybq7pbtsnvsr</settlement-batch-id>
//   <custom-fields>
//   </custom-fields>
//   <avs-error-response-code nil="true"></avs-error-response-code>
//   <avs-postal-code-response-code>I</avs-postal-code-response-code>
//   <avs-street-address-response-code>I</avs-street-address-response-code>
//   <cvv-response-code>I</cvv-response-code>
//   <gateway-rejection-reason nil="true"></gateway-rejection-reason>
//   <processor-authorization-code>YCSBWR</processor-authorization-code>
//   <processor-response-code>1000</processor-response-code>
//   <processor-response-text>Approved</processor-response-text>
//   <voice-referral-number nil="true"></voice-referral-number>
//   <purchase-order-number nil="true"></purchase-order-number>
//   <tax-amount nil="true"></tax-amount>
//   <tax-exempt type="boolean">false</tax-exempt>
//   <add-ons type="array"/>
//   <discounts type="array"/>
//   <descriptor>
//     <name nil="true"></name>
//     <phone nil="true"></phone>
//   </descriptor>
//   <recurring type="boolean">true</recurring>
//   <channel nil="true"></channel>
//   <escrow-status nil="true"></escrow-status>
// </transaction>

type Transactions struct {
	Transaction []*Transaction `xml:"transaction"`
}

type TransactionOptions struct {
	SubmitForSettlement              bool `xml:"submit-for-settlement,omitempty"`
	StoreInVault                     bool `xml:"store-in-vault,omitempty"`
	AddBillingAddressToPaymentMethod bool `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `xml:"store-shipping-address-in-vault,omitempty"`
}

type TransactionSearchResult struct {
	XMLName           string              `xml:"credit-card-transactions"`
	CurrentPageNumber *nullable.NullInt64 `xml:"current-page-number"`
	PageSize          *nullable.NullInt64 `xml:"page-size"`
	TotalItems        *nullable.NullInt64 `xml:"total-items"`
	Transactions      []*Transaction      `xml:"transaction"`
}
