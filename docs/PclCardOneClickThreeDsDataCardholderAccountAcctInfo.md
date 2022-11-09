# PclCardOneClickThreeDsDataCardholderAccountAcctInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChAccAgeInd** | **string** | Length of time that the cardholder has had the account with the Merchant. Values accepted: 01 &#x3D; No account (guest check-out) 02 &#x3D; Created during this transaction 03 &#x3D; Less than 30 days 04 &#x3D; 30−60 days 05 &#x3D; More than 60 days | [optional] [default to null]
**ChAccChange** | **string** | Date that the cardholder&#x27;s account with the Merchant was last changed. Including Billing or Shipping address, new payment account, or new user(s) added. Format accepted: YYYYMMDD | [optional] [default to null]
**ChAccChangeInd** | **string** | Length of time since the cardholder&#x27;s account information with the Merchant was last changed. Including Billing or Shipping address, new payment account, or new user(s) added. Values accepted: 01 &#x3D; Changed during this transaction 02 &#x3D; Less than 30 days 03 &#x3D; 30−60 days 04 &#x3D; More than 60 days | [optional] [default to null]
**ChAccDate** | **string** | Date that the cardholder opened the account with the Merchant. Date format &#x3D; YYYYMMDD | [optional] [default to null]
**ChAccPwChange** | **string** | Date that cardholder&#x27;s account with the Merchant had a password change or account reset. Format accepted: YYYYMMDD | [optional] [default to null]
**ChAccPwChangeInd** | **string** | Indicates the length of time since the cardholder&#x27;s account with the Merchant had a password change or account reset. Values accepted: 01 &#x3D; No change 02 &#x3D; Changed during this transaction 03 &#x3D; Less than 30 days 04 &#x3D; 30−60 days 05 &#x3D; More than 60 days | [optional] [default to null]
**NbPurchaseAccount** | **float64** | Number of purchases with this cardholder account during the previous six months. | [optional] [default to null]
**PaymentAccAge** | **string** | Date that the payment account was enrolled in the cardholder&#x27;s account with the Merchant. Format accepted: YYYYMMDD | [optional] [default to null]
**PaymentAccInd** | **string** | Indicates the length of time that the payment account was enrolled in the cardholder&#x27;s account with the Merchant. Values accepted: 01 &#x3D; No account (guest check-out) 02 &#x3D; During this transaction 03 &#x3D; Less than 30 days 04 &#x3D; 30−60 days 05 &#x3D; More than 60 days | [optional] [default to null]
**ProvisionAttemptsDay** | **float64** | Number of Add Card attempts in the last 24 hours. | [optional] [default to null]
**ShipAddressUsage** | **string** | Date when the shipping address used for this transaction was first used with the Merchant. Format accepted: YYYYMMDD | [optional] [default to null]
**ShipAddressUsageInd** | **string** | Indicates when the shipping address used for this transaction was first used with the Merchant. Values accepted: 01 &#x3D; This transaction 02 &#x3D; Less than 30 days 03 &#x3D; 30−60 days 04 &#x3D; More than 60 days | [optional] [default to null]
**ShipNameIndicator** | **string** | Indicates if the Cardholder Name on the account is identical to the shipping Name used for this transaction. Values accepted: 01 &#x3D; Account Name identical to shipping Name 02 &#x3D; Account Name different than shipping Name | [optional] [default to null]
**SuspiciousAccActivity** | **string** | Indicates whether the Merchant has experienced suspicious activity (including previous fraud) on the cardholder account. Values accepted: 01 &#x3D; No suspicious activity has been observed 02 &#x3D; Suspicious activity has been observed | [optional] [default to null]
**TxnActivityDay** | **float64** | Number of transactions (successful and abandoned) for this cardholder account with the Merchant across all payment accounts in the previous 24 hours. Length: maximum 3 characters | [optional] [default to null]
**TxnActivityYear** | **float64** | Number of transactions (successful and abandoned) for this cardholder account with the Merchant across all payment accounts in the previous year. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

