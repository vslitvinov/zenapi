# PclHostedCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeDsData** | [***PclCardOneClickThreeDsData**](PCL_CARDOneClick_threeDsData.md) |  | [optional] [default to null]
**ReturnUrl** | **string** | Return URL after successul payment. (Will override defualt shop settings) | [default to null]
**CardToken** | **string** | Card token created in the process of saving credit card | [optional] [default to null]
**Descriptor** | **string** | Text that will appear on customer Bank Statement. Can be used only for Credit Card payments. | [default to null]
**SaveCard** | **bool** | Defines whether card should be saved for future purchases/recurring payments | [optional] [default to null]
**ReturnVerifyUrl** | **string** | Return URL after 3DS authentiaction | [default to null]
**ScaExemptions** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

