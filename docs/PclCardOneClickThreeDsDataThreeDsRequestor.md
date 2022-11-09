# PclCardOneClickThreeDsDataThreeDsRequestor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeDSRequestorAuthenticationInfo** | [***PclCardOneClickThreeDsDataThreeDsRequestorThreeDsRequestorAuthenticationInfo**](PCL_CARDOneClick_threeDsData_threeDSRequestor_threeDSRequestorAuthenticationInfo.md) |  | [optional] [default to null]
**ThreeDSRequestorChallengeInd** | **string** | Indicates whether a challenge is requested for this transaction. (This field is optional, For 01-PA, a Merchant may have concerns about the transaction, and request a challenge; For 02-NPA, a challenge may be necessary when adding a new card to a wallet.) Length: 2 characters 01 &#x3D; No preference – merchant do not have preferences whether a challenge (client authentication) should be performed 02 &#x3D; No challenge requested – merchant doesn’t want challenge (client authentication) to be performed 03 &#x3D; Challenge requested by merchant – merchant want challenge (client authentication) to be performed 04 &#x3D; Challenge requested (Mandate) – local requirements demand a challenge (client authentication) to be performed, e.g. in case of soft decline | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

