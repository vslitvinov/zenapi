# Go API client for swagger

ZEN allows you to accept multiple global and local payment methods (including credit cards, instant bank transfers, and many more). The following document will guide you through the integration process and introduce you to the main features of the Public API..<br><br><p><strong><span data-contrast=\"auto\">1. Allowed HTTPs requests:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">PU</span></strong><strong><span data-contrast=\"auto\">T</span></strong><strong><span data-contrast=\"auto\">:</span></strong><span data-contrast=\"auto\">&nbsp;To create a resource&nbsp;</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">POST:</span></strong><span data-contrast=\"auto\">&nbsp;To update ra esource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">GET:</span></strong><span data-contrast=\"auto\">&nbsp;To get a resource or a list of resources</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">DELETE:</span></strong><span data-contrast=\"auto\">&nbsp;To delete a resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ul><p><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span><strong><span data-contrast=\"auto\">2. Description&nbsp;</span></strong><strong><span data-contrast=\"auto\">o</span></strong><strong><span data-contrast=\"auto\">f&nbsp;</span></strong><strong><span data-contrast=\"auto\">u</span></strong><strong><span data-contrast=\"auto\">sual&nbsp;</span></strong><strong><span data-contrast=\"auto\">s</span></strong><strong><span data-contrast=\"auto\">erver&nbsp;</span></strong><strong><span data-contrast=\"auto\">r</span></strong><strong><span data-contrast=\"auto\">esponses:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">200</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">--</span></strong><strong><span data-contrast=\"auto\">&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">OK</span></strong><span data-contrast=\"auto\"> - the request was successful (some API calls may return 201 instead)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">201&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Created</span></strong><span data-contrast=\"auto\"> - the request was successful and a resource was created</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">204  </span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">No Content</span></strong><span data-contrast=\"auto\"> - the request was successful, but there is no representation to return (</span><span data-contrast=\"auto\">i</span><span data-contrast=\"auto\">.e. the response is empty)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"4\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">400&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">Bad Request</span></strong><span data-contrast=\"auto\"> - the request could not be understood or was missing required parameters</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"5\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">401&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Unauthorized</span></strong><span data-contrast=\"auto\"> - authentication failed or user does not have permissions for requested operation</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">403&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\">&nbsp;</span><strong><span data-contrast=\"auto\">Forbidden </span></strong><span data-contrast=\"auto\"> - access denied</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">404&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\"> </span><strong><span data-contrast=\"auto\">Not Found</span></strong><span data-contrast=\"auto\"> - the resource was not found</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">405  </span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">Method Not Allowed</span></strong><span data-contrast=\"auto\"> - requested method is not supported for the resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li></ul><p><strong><span data-contrast=\"auto\">3.&nbsp;</span></strong><strong><span data-contrast=\"auto\">API Integration Guidelines</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><span data-contrast=\"auto\">To start using ZEN eCommerce Public API please follow these instructions:&nbsp;</span><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:6,&quot;335551620&quot;:6,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#before-you-begin%20\">https://www.zen.com/developer/api-integration#before-you-begin&nbsp;</a><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559737&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ol><p><strong><span data-contrast=\"auto\">4. IPN API Secrets</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><p><span data-contrast=\"auto\">Instant Payment Notification (IPN) informs you about transaction status updates, e.g., about changing the status from pending to rejected. </span><br /><span data-contrast=\"auto\">Information abut IPN is available in the following link:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#ipn\"><span data-contrast=\"auto\">https://www.zen.com/developer/api-integration#ipn</span></a></p> <p><strong><span data-contrast=\"auto\">4. Error codes</span></strong><p><span data-contrast=\"auto\">The full list of error codes can be found here:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#rejection-codes\">https://www.zen.com/developer/api-integration#rejection-codes</a></p>  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.zen.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PaymentProfilesApi* | [**CardDelete**](docs/PaymentProfilesApi.md#carddelete) | **Delete** /payment-profiles/card/{id} | Remove saved card
*PaymentProfilesApi* | [**CardFetch**](docs/PaymentProfilesApi.md#cardfetch) | **Get** /payment-profiles/card/{id} | Fetch saved card
*PaymentProfilesApi* | [**CardList**](docs/PaymentProfilesApi.md#cardlist) | **Get** /payment-profiles/card | Fetch saved card list
*PaymentProfilesApi* | [**CardUpdate**](docs/PaymentProfilesApi.md#cardupdate) | **Patch** /payment-profiles/card/{id} | Update credit card details
*TerminalsApi* | [**InfoMethods**](docs/TerminalsApi.md#infomethods) | **Get** /terminals | Return information about payment methods on a Terminal
*TransactionsApi* | [**CreateRefund**](docs/TransactionsApi.md#createrefund) | **Post** /transactions/refund | Process transaction refunds using initial transaction&#x27;s ID, received in POST /transaction
*TransactionsApi* | [**CreateTransaction**](docs/TransactionsApi.md#createtransaction) | **Post** /transactions | Create transaction
*TransactionsApi* | [**GetMerchantTransaction**](docs/TransactionsApi.md#getmerchanttransaction) | **Get** /transactions/merchant/{merchantTransactionId} | Get transaction details using merchant transaction ID
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /transactions/{id} | Get transaction details using ZEN ID
*TransactionsApi* | [**RenewAuthorization**](docs/TransactionsApi.md#renewauthorization) | **Post** /transactions/renewAuthorization | Create new authorization attempt on the same TRX transaction in case of expired or rejected previous authorization

## Documentation For Models

 - [ActionCreateObject](docs/ActionCreateObject.md)
 - [ActionRefundObject](docs/ActionRefundObject.md)
 - [ActionRenewAuthorizationObject](docs/ActionRenewAuthorizationObject.md)
 - [AnyOfActionCreateObjectFraudFields](docs/AnyOfActionCreateObjectFraudFields.md)
 - [AnyOfCustomerIp](docs/AnyOfCustomerIp.md)
 - [AnyOfDetailsAboutTransactionFraudFields](docs/AnyOfDetailsAboutTransactionFraudFields.md)
 - [AnyOfExternalCustomerIp](docs/AnyOfExternalCustomerIp.md)
 - [AnyOfExternalTransactionDetailsCustomerIp](docs/AnyOfExternalTransactionDetailsCustomerIp.md)
 - [AnyOfExternalTransactionDetailsFraudFields](docs/AnyOfExternalTransactionDetailsFraudFields.md)
 - [AnyOfTopupCustomerIp](docs/AnyOfTopupCustomerIp.md)
 - [AnyOfTransactionDetailsFraudFields](docs/AnyOfTransactionDetailsFraudFields.md)
 - [AnyOftransactionsCustomerIp](docs/AnyOftransactionsCustomerIp.md)
 - [Card](docs/Card.md)
 - [CardAdd](docs/CardAdd.md)
 - [CardDeleteOutput](docs/CardDeleteOutput.md)
 - [CardOneclick](docs/CardOneclick.md)
 - [CardOperation](docs/CardOperation.md)
 - [CardOutput](docs/CardOutput.md)
 - [CashbackInput](docs/CashbackInput.md)
 - [CashbackInputClient](docs/CashbackInputClient.md)
 - [CashbackItem](docs/CashbackItem.md)
 - [CashbackItemType](docs/CashbackItemType.md)
 - [CashbackOutput](docs/CashbackOutput.md)
 - [CombinedPaymentMethodsInformation](docs/CombinedPaymentMethodsInformation.md)
 - [CurrencyPair](docs/CurrencyPair.md)
 - [Customer](docs/Customer.md)
 - [DetailedInformationAboutCustomerForPayout](docs/DetailedInformationAboutCustomerForPayout.md)
 - [DetailsAboutRefundTransaction](docs/DetailsAboutRefundTransaction.md)
 - [DetailsAboutTransaction](docs/DetailsAboutTransaction.md)
 - [Direction](docs/Direction.md)
 - [Dragon](docs/Dragon.md)
 - [DragonDeleteOutput](docs/DragonDeleteOutput.md)
 - [DragonOutput](docs/DragonOutput.md)
 - [ExternalCustomer](docs/ExternalCustomer.md)
 - [ExternalTransactionDetails](docs/ExternalTransactionDetails.md)
 - [ExternalTransactionDetailsCashback](docs/ExternalTransactionDetailsCashback.md)
 - [ExternalTransactionDetailsCashbackClient](docs/ExternalTransactionDetailsCashbackClient.md)
 - [ExternalTransactionDetailsCustomer](docs/ExternalTransactionDetailsCustomer.md)
 - [ExternalTransactionDetailsMeta](docs/ExternalTransactionDetailsMeta.md)
 - [ExternalTransactionDetailsMetaThreeDs](docs/ExternalTransactionDetailsMetaThreeDs.md)
 - [Fee](docs/Fee.md)
 - [FeeOwner](docs/FeeOwner.md)
 - [FingerprintMerchantAction](docs/FingerprintMerchantAction.md)
 - [FingerprintMerchantActionData](docs/FingerprintMerchantActionData.md)
 - [FraudData](docs/FraudData.md)
 - [HealthcheckOutput](docs/HealthcheckOutput.md)
 - [HealthcheckOutputComponents](docs/HealthcheckOutputComponents.md)
 - [ImageMerchantAction](docs/ImageMerchantAction.md)
 - [ImageMerchantActionData](docs/ImageMerchantActionData.md)
 - [InformationObject](docs/InformationObject.md)
 - [InformationObject1](docs/InformationObject1.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [MaintenanceInformation](docs/MaintenanceInformation.md)
 - [MerchantAction](docs/MerchantAction.md)
 - [Meta](docs/Meta.md)
 - [Meta1](docs/Meta1.md)
 - [MetaVersionOutput](docs/MetaVersionOutput.md)
 - [MetaVersionOutputBuild](docs/MetaVersionOutputBuild.md)
 - [MetaVersionOutputGit](docs/MetaVersionOutputGit.md)
 - [MethodInformation](docs/MethodInformation.md)
 - [MethodInformationObject](docs/MethodInformationObject.md)
 - [MethodInformationObject1](docs/MethodInformationObject1.md)
 - [OneOfActionCreateObjectPaymentSpecificData](docs/OneOfActionCreateObjectPaymentSpecificData.md)
 - [OneOfActionRenewAuthorizationObjectPaymentSpecificData](docs/OneOfActionRenewAuthorizationObjectPaymentSpecificData.md)
 - [OneOfDetailsAboutTransactionMerchantAction](docs/OneOfDetailsAboutTransactionMerchantAction.md)
 - [OneOfExternalTransactionDetailsMerchantAction](docs/OneOfExternalTransactionDetailsMerchantAction.md)
 - [OneOfPlanCreatedAt](docs/OneOfPlanCreatedAt.md)
 - [OneOfPlanUpdatedAt](docs/OneOfPlanUpdatedAt.md)
 - [OneOfProductCreatedAt](docs/OneOfProductCreatedAt.md)
 - [OneOfProductUpdatedAt](docs/OneOfProductUpdatedAt.md)
 - [OneOfSubscriptionCreatedAt](docs/OneOfSubscriptionCreatedAt.md)
 - [OneOfSubscriptionUpdatedAt](docs/OneOfSubscriptionUpdatedAt.md)
 - [OneOfTransactionDetailsMerchantAction](docs/OneOfTransactionDetailsMerchantAction.md)
 - [PaymentCustomOptions](docs/PaymentCustomOptions.md)
 - [PayoutCustomOptions](docs/PayoutCustomOptions.md)
 - [PayoutInfo](docs/PayoutInfo.md)
 - [PclBitbayWithdrawal](docs/PclBitbayWithdrawal.md)
 - [PclBlik](docs/PclBlik.md)
 - [PclCardAuthCheck](docs/PclCardAuthCheck.md)
 - [PclCardAuthCheckToken](docs/PclCardAuthCheckToken.md)
 - [PclCardExternalPaymentToken](docs/PclCardExternalPaymentToken.md)
 - [PclCardOneClick](docs/PclCardOneClick.md)
 - [PclCardOneClickBrowserDetails](docs/PclCardOneClickBrowserDetails.md)
 - [PclCardOneClickCard](docs/PclCardOneClickCard.md)
 - [PclCardOneClickSubscription](docs/PclCardOneClickSubscription.md)
 - [PclCardOneClickThreeDsData](docs/PclCardOneClickThreeDsData.md)
 - [PclCardOneClickThreeDsDataCardholder](docs/PclCardOneClickThreeDsDataCardholder.md)
 - [PclCardOneClickThreeDsDataCardholderAccount](docs/PclCardOneClickThreeDsDataCardholderAccount.md)
 - [PclCardOneClickThreeDsDataCardholderAccountAcctInfo](docs/PclCardOneClickThreeDsDataCardholderAccountAcctInfo.md)
 - [PclCardOneClickThreeDsDataMerchantRiskIndicator](docs/PclCardOneClickThreeDsDataMerchantRiskIndicator.md)
 - [PclCardOneClickThreeDsDataPurchase](docs/PclCardOneClickThreeDsDataPurchase.md)
 - [PclCardOneClickThreeDsDataThreeDsRequestor](docs/PclCardOneClickThreeDsDataThreeDsRequestor.md)
 - [PclCardOneClickThreeDsDataThreeDsRequestorThreeDsRequestorAuthenticationInfo](docs/PclCardOneClickThreeDsDataThreeDsRequestorThreeDsRequestorAuthenticationInfo.md)
 - [PclCardOneTime](docs/PclCardOneTime.md)
 - [PclCardOneTimeCard](docs/PclCardOneTimeCard.md)
 - [PclCardRecurringFullData](docs/PclCardRecurringFullData.md)
 - [PclCardRecurringFullDataCard](docs/PclCardRecurringFullDataCard.md)
 - [PclCardRecurringToken](docs/PclCardRecurringToken.md)
 - [PclCardUnscheduled](docs/PclCardUnscheduled.md)
 - [PclDragon](docs/PclDragon.md)
 - [PclDragonCapture](docs/PclDragonCapture.md)
 - [PclDragonUser](docs/PclDragonUser.md)
 - [PclDragonUserCapture](docs/PclDragonUserCapture.md)
 - [PclDragonV2](docs/PclDragonV2.md)
 - [PclG2Apay](docs/PclG2Apay.md)
 - [PclG2ApayTopup](docs/PclG2ApayTopup.md)
 - [PclGeneralPayment](docs/PclGeneralPayment.md)
 - [PclHostedCard](docs/PclHostedCard.md)
 - [PclTrustly](docs/PclTrustly.md)
 - [PclWebmoneyUser](docs/PclWebmoneyUser.md)
 - [Plan](docs/Plan.md)
 - [Product](docs/Product.md)
 - [RedirectMerchantAction](docs/RedirectMerchantAction.md)
 - [RedirectMerchantActionData](docs/RedirectMerchantActionData.md)
 - [SortBy](docs/SortBy.md)
 - [Status](docs/Status.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionCustomer](docs/SubscriptionCustomer.md)
 - [SubscriptionSortBy](docs/SubscriptionSortBy.md)
 - [TerminalsCashback](docs/TerminalsCashback.md)
 - [ThreeDsData](docs/ThreeDsData.md)
 - [ThreeDsDataPa](docs/ThreeDsDataPa.md)
 - [TopupCustomOptions](docs/TopupCustomOptions.md)
 - [TopupCustomer](docs/TopupCustomer.md)
 - [TopupMethod](docs/TopupMethod.md)
 - [TopupMethodFee](docs/TopupMethodFee.md)
 - [TopupMethodTopup](docs/TopupMethodTopup.md)
 - [TransactionActions](docs/TransactionActions.md)
 - [TransactionCardInfo](docs/TransactionCardInfo.md)
 - [TransactionDetails](docs/TransactionDetails.md)
 - [TransactionFee](docs/TransactionFee.md)
 - [TransactionItems](docs/TransactionItems.md)
 - [TransactionRefund](docs/TransactionRefund.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionsBillingAddress](docs/TransactionsBillingAddress.md)
 - [TransactionsCashback](docs/TransactionsCashback.md)
 - [TransactionsCashbackClient](docs/TransactionsCashbackClient.md)
 - [TransactionsCustomer](docs/TransactionsCustomer.md)
 - [TransactionsItems](docs/TransactionsItems.md)
 - [TransactionsShippingAddress](docs/TransactionsShippingAddress.md)
 - [UpdateCard](docs/UpdateCard.md)
 - [WebmoneyDeleteOutput](docs/WebmoneyDeleteOutput.md)
 - [WebmoneyOutput](docs/WebmoneyOutput.md)

## Documentation For Authorization

## ApiKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author


