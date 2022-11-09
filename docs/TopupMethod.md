# TopupMethod

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentChannel** | **string** | ID of payment channel for selected payment method  | [default to null]
**PaymentChannelLogoUrl** | **string** |  | [optional] [default to null]
**Currency** | **string** | Currency code in ISO 4217 alphabetic code | [default to null]
**CryptoCurrencies** | **[]string** |  | [optional] [default to null]
**Topup** | [***TopupMethodTopup**](TopupMethod_topup.md) |  | [default to null]
**Fee** | [***TopupMethodFee**](TopupMethod_fee.md) |  | [optional] [default to null]
**MinAmount** | **string** | Transaction amount  | [optional] [default to null]
**MaxAmount** | **string** | Transaction amount  | [optional] [default to null]
**LastUsedDetails** | [**map[string]Object**](.md) | Additional data stored as key/value pairs | [optional] [default to null]
**IsLast** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

