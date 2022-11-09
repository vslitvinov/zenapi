# PclCardOneClickThreeDsDataMerchantRiskIndicator

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryEmailAddress** | **string** | For Electronic delivery, the email address to which the merchandise was delivered. Length: maximum 254 characters | [optional] [default to null]
**DeliveryTimeframe** | **string** | Indicates the merchandise delivery timeframe. Values accepted: 01 &#x3D; Electronic Delivery 02 &#x3D; Same day shipping 03 &#x3D; Overnight shipping 04 &#x3D; Two-day or more shipping | [optional] [default to null]
**GiftCardAmount** | **string** | For prepaid or gift card purchase, the purchase amount total of prepaid or gift card(s) in major units (for example, USD 123.45 is 123). Length: maximum 15 characters | [optional] [default to null]
**GiftCardCount** | **string** | For prepaid or gift card purchase, total count of individual prepaid or gift cards/codes purchased. Length: 2 characters | [optional] [default to null]
**GiftCardCurr** | **string** | For prepaid or gift card purchase, the currency code of the card as defined in ISO 4217 other than those listed in Table A.6. in EMVCO Protocol and Core Functions Specification Length: 3 characters | [optional] [default to null]
**PreOrderDate** | **string** | For a pre-ordered purchase, the expected date that the merchandise will be available. Format accepted: YYYYMMDD | [optional] [default to null]
**PreOrderPurchaseInd** | **string** | Indicates whether Cardholder is placing an order for merchandise with a future availability or release date. Values accepted: 01 &#x3D; Merchandise available 02 &#x3D; Future availability | [optional] [default to null]
**ReorderItemsInd** | **string** | Indicates whether the cardholder is reordering previously purchased merchandise. Values accepted:01 &#x3D; First time ordered 02 &#x3D; Reordered | [optional] [default to null]
**ShipIndicator** | **string** | Indicates shipping method chosen for the transaction. Merchants must choose the Shipping Indicator code that most accurately describes the cardholder&#x27;s specific transaction, not their general business. If one or more items are included in the sale, use the Shipping Indicator code for the physical goods, or if all digital goods, use the Shipping Indicator code that describes the most expensive item. Values accepted: 01 &#x3D; Ship to cardholder’s billing address 02 &#x3D; Ship to another verified address on file with merchant 03 &#x3D; Ship to address that is different than the cardholder’s billing address 04 &#x3D; “Ship to Store” / Pick-up at local store (Store address shall be populated in shipping address fields) 05 &#x3D; Digital goods (includes online services, electronic gift cards and redemption codes) 06 &#x3D; Travel and Event tickets, not shipped 07 &#x3D; Other (for example, Gaming, digital services not shipped, emedia subscriptions, etc.) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

