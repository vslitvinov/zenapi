# PclCardAuthCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Type of transaction | [default to null]
**ScaExemptions** | **string** |  | [optional] [default to null]
**BrowserDetails** | [***PclCardOneClickBrowserDetails**](PCL_CARDOneClick_browserDetails.md) |  | [default to null]
**ThreeDsData** | [***PclCardOneClickThreeDsData**](PCL_CARDOneClick_threeDsData.md) |  | [optional] [default to null]
**AuthenticationServiceTarget** | **string** |  | [optional] [default to null]
**Subscription** | [***PclCardOneClickSubscription**](PCL_CARDOneClick_subscription.md) |  | [optional] [default to null]
**ReturnVerifyUrl** | **string** | Return URL after 3DS authentiaction | [optional] [default to null]
**SaveCard** | **bool** | Defines whether card should be saved for future purchases/recurring payments | [optional] [default to null]
**Descriptor** | **string** | Text that will appear on customer Bank Statement. Can be used only for Credit Card payments. | [default to null]
**Card** | [***PclCardOneTimeCard**](PCL_CARDOneTime_card.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

