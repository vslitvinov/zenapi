# TransactionDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**RedirectUrl** | **string** | Redirect URL provided by PSP. If \&quot;actions\&quot; states it, buyer should be redirected to this URL to proceed with payment | [optional] [default to null]
**ImageUrl** | **string** | Image URL after successul payment initialization (e.g., for QR code payment) | [optional] [default to null]
**MerchantAction** | [***OneOfTransactionDetailsMerchantAction**](OneOfTransactionDetailsMerchantAction.md) |  | [optional] [default to null]
**MerchantTransactionId** | **string** | Merchant transaction ID. Should be unique in merchant store. | [default to null]
**Amount** | **string** | Transaction amount  | [default to null]
**Currency** | **string** | Currency code in ISO 4217 alphabetic code | [default to null]
**Fee** | [***TransactionFee**](Transaction fee.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [default to null]
**ModifiedAt** | [**time.Time**](time.Time.md) |  | [default to null]
**Type_** | **string** | Type of transaction | [default to null]
**Status** | **string** | Transaction status | [default to null]
**PaymentChannel** | **string** | ID of payment channel for selected payment method  | [default to null]
**Actions** | [***CombinedPaymentMethodsInformation**](Combined payment methods information.md) |  | [optional] [default to null]
**FraudFields** | [***AnyOfTransactionDetailsFraudFields**](AnyOfTransactionDetailsFraudFields.md) |  | [optional] [default to null]
**RejectCode** | **string** |  | [optional] [default to null]
**RejectReason** | **string** |  | [optional] [default to null]
**Refunds** | [**[]DetailsAboutRefundTransaction**](Details about refund transaction.md) |  | [optional] [default to null]
**Meta** | [***ExternalTransactionDetailsMeta**](ExternalTransactionDetails_meta.md) |  | [optional] [default to null]
**Customer** | [***TransactionsCustomer**](transactions_customer.md) |  | [optional] [default to null]
**CardInfo** | [***TransactionCardInfo**](Transaction card info.md) |  | [optional] [default to null]
**BillingAddress** | [***TransactionsBillingAddress**](transactions_billingAddress.md) |  | [optional] [default to null]
**ShippingAddress** | [***TransactionsShippingAddress**](transactions_shippingAddress.md) |  | [optional] [default to null]
**Items** | [**[]TransactionsItems**](transactions_items.md) | If items are provided sum of items amount should be equal transaction amount | [optional] [default to null]
**VerifyReturnmac** | **string** |  | [optional] [default to null]
**Cashback** | [***ExternalTransactionDetailsCashback**](ExternalTransactionDetails_cashback.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

