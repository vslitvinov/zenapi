# PclCardOneClickThreeDsDataThreeDsRequestorThreeDsRequestorAuthenticationInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeDSReqAuthData** | **string** | Data that documents and supports a specific authentication process. Length: maximum 2048 bytes, JSON Data Type: String, Value accepted: Any | [optional] [default to null]
**ThreeDSReqAuthMethod** | **string** | Mechanism used by the Cardholder to authenticate to the Merchant. Values accepted: 01 &#x3D; No Merchant authentication occurred (i.e. cardholder “logged in” as guest) 02 &#x3D; Login to the cardholder account at the Merchant system using Merchant’s own credentials 03 &#x3D; Login to the cardholder account at the Merchant system using federated ID 04 &#x3D; Login to the cardholder account at the Merchant system using issuer credentials 05 &#x3D; Login to the cardholder account at the Merchant system using third-party authentication 06 &#x3D; Login to the cardholder account at the Merchant system using FIDO Authenticator  | [optional] [default to null]
**ThreeDSReqAuthTimestamp** | **string** | Date and time in UTC of the cardholder authentication. Format accepted: YYYYMMDDHHMM | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

