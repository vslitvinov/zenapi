# ActionCreateObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantTransactionId** | **string** | ID of transaction provided by merchant | [default to null]
**PaymentChannel** | **string** | ID of payment channel for selected payment method  | [default to null]
**Amount** | **string** | Transaction amount  | [default to null]
**Currency** | **string** | Currency code in ISO 4217 alphabetic code | [default to null]
**CustomIpnUrl** | **string** | URL address used by ZEN to send IPN to | [optional] [default to null]
**Comment** | **string** |  | [optional] [default to null]
**FraudFields** | [***AnyOfActionCreateObjectFraudFields**](AnyOfActionCreateObjectFraudFields.md) |  | [optional] [default to null]
**Items** | [**[]TransactionsItems**](transactions_items.md) | Sum of items amount should be equal to transaction amount | [optional] [default to null]
**Customer** | [***TransactionsCustomer**](transactions_customer.md) |  | [default to null]
**PaymentSpecificData** | [***OneOfActionCreateObjectPaymentSpecificData**](OneOfActionCreateObjectPaymentSpecificData.md) |  | [optional] [default to null]
**BillingAddress** | [***TransactionsBillingAddress**](transactions_billingAddress.md) |  | [optional] [default to null]
**ShippingAddress** | [***TransactionsShippingAddress**](transactions_shippingAddress.md) |  | [optional] [default to null]
**Cashback** | [***TransactionsCashback**](transactions_cashback.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

