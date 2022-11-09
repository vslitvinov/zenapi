/*
 * ZEN eCommerce Public API
 *
 * ZEN allows you to accept multiple global and local payment methods (including credit cards, instant bank transfers, and many more). The following document will guide you through the integration process and introduce you to the main features of the Public API..<br><br><p><strong><span data-contrast=\"auto\">1. Allowed HTTPs requests:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">PU</span></strong><strong><span data-contrast=\"auto\">T</span></strong><strong><span data-contrast=\"auto\">:</span></strong><span data-contrast=\"auto\">&nbsp;To create a resource&nbsp;</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">POST:</span></strong><span data-contrast=\"auto\">&nbsp;To update ra esource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">GET:</span></strong><span data-contrast=\"auto\">&nbsp;To get a resource or a list of resources</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">DELETE:</span></strong><span data-contrast=\"auto\">&nbsp;To delete a resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ul><p><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span><strong><span data-contrast=\"auto\">2. Description&nbsp;</span></strong><strong><span data-contrast=\"auto\">o</span></strong><strong><span data-contrast=\"auto\">f&nbsp;</span></strong><strong><span data-contrast=\"auto\">u</span></strong><strong><span data-contrast=\"auto\">sual&nbsp;</span></strong><strong><span data-contrast=\"auto\">s</span></strong><strong><span data-contrast=\"auto\">erver&nbsp;</span></strong><strong><span data-contrast=\"auto\">r</span></strong><strong><span data-contrast=\"auto\">esponses:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">200</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">--</span></strong><strong><span data-contrast=\"auto\">&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">OK</span></strong><span data-contrast=\"auto\"> - the request was successful (some API calls may return 201 instead)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">201&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Created</span></strong><span data-contrast=\"auto\"> - the request was successful and a resource was created</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">204  </span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">No Content</span></strong><span data-contrast=\"auto\"> - the request was successful, but there is no representation to return (</span><span data-contrast=\"auto\">i</span><span data-contrast=\"auto\">.e. the response is empty)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"4\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">400&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">Bad Request</span></strong><span data-contrast=\"auto\"> - the request could not be understood or was missing required parameters</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"5\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">401&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Unauthorized</span></strong><span data-contrast=\"auto\"> - authentication failed or user does not have permissions for requested operation</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">403&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\">&nbsp;</span><strong><span data-contrast=\"auto\">Forbidden </span></strong><span data-contrast=\"auto\"> - access denied</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">404&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\"> </span><strong><span data-contrast=\"auto\">Not Found</span></strong><span data-contrast=\"auto\"> - the resource was not found</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">405  </span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">Method Not Allowed</span></strong><span data-contrast=\"auto\"> - requested method is not supported for the resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li></ul><p><strong><span data-contrast=\"auto\">3.&nbsp;</span></strong><strong><span data-contrast=\"auto\">API Integration Guidelines</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><span data-contrast=\"auto\">To start using ZEN eCommerce Public API please follow these instructions:&nbsp;</span><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:6,&quot;335551620&quot;:6,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#before-you-begin%20\">https://www.zen.com/developer/api-integration#before-you-begin&nbsp;</a><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559737&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ol><p><strong><span data-contrast=\"auto\">4. IPN API Secrets</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><p><span data-contrast=\"auto\">Instant Payment Notification (IPN) informs you about transaction status updates, e.g., about changing the status from pending to rejected. </span><br /><span data-contrast=\"auto\">Information abut IPN is available in the following link:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#ipn\"><span data-contrast=\"auto\">https://www.zen.com/developer/api-integration#ipn</span></a></p> <p><strong><span data-contrast=\"auto\">4. Error codes</span></strong><p><span data-contrast=\"auto\">The full list of error codes can be found here:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#rejection-codes\">https://www.zen.com/developer/api-integration#rejection-codes</a></p>  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Additional information about the Cardholder’s account provided by the Merchant. Optional, but strongly recommended to include.
type PclCardOneClickThreeDsDataCardholderAccountAcctInfo struct {
	// Length of time that the cardholder has had the account with the Merchant. Values accepted: 01 = No account (guest check-out) 02 = Created during this transaction 03 = Less than 30 days 04 = 30−60 days 05 = More than 60 days
	ChAccAgeInd string `json:"chAccAgeInd,omitempty"`
	// Date that the cardholder's account with the Merchant was last changed. Including Billing or Shipping address, new payment account, or new user(s) added. Format accepted: YYYYMMDD
	ChAccChange string `json:"chAccChange,omitempty"`
	// Length of time since the cardholder's account information with the Merchant was last changed. Including Billing or Shipping address, new payment account, or new user(s) added. Values accepted: 01 = Changed during this transaction 02 = Less than 30 days 03 = 30−60 days 04 = More than 60 days
	ChAccChangeInd string `json:"chAccChangeInd,omitempty"`
	// Date that the cardholder opened the account with the Merchant. Date format = YYYYMMDD
	ChAccDate string `json:"chAccDate,omitempty"`
	// Date that cardholder's account with the Merchant had a password change or account reset. Format accepted: YYYYMMDD
	ChAccPwChange string `json:"chAccPwChange,omitempty"`
	// Indicates the length of time since the cardholder's account with the Merchant had a password change or account reset. Values accepted: 01 = No change 02 = Changed during this transaction 03 = Less than 30 days 04 = 30−60 days 05 = More than 60 days
	ChAccPwChangeInd string `json:"chAccPwChangeInd,omitempty"`
	// Number of purchases with this cardholder account during the previous six months.
	NbPurchaseAccount float64 `json:"nbPurchaseAccount,omitempty"`
	// Date that the payment account was enrolled in the cardholder's account with the Merchant. Format accepted: YYYYMMDD
	PaymentAccAge string `json:"paymentAccAge,omitempty"`
	// Indicates the length of time that the payment account was enrolled in the cardholder's account with the Merchant. Values accepted: 01 = No account (guest check-out) 02 = During this transaction 03 = Less than 30 days 04 = 30−60 days 05 = More than 60 days
	PaymentAccInd string `json:"paymentAccInd,omitempty"`
	// Number of Add Card attempts in the last 24 hours.
	ProvisionAttemptsDay float64 `json:"provisionAttemptsDay,omitempty"`
	// Date when the shipping address used for this transaction was first used with the Merchant. Format accepted: YYYYMMDD
	ShipAddressUsage string `json:"shipAddressUsage,omitempty"`
	// Indicates when the shipping address used for this transaction was first used with the Merchant. Values accepted: 01 = This transaction 02 = Less than 30 days 03 = 30−60 days 04 = More than 60 days
	ShipAddressUsageInd string `json:"shipAddressUsageInd,omitempty"`
	// Indicates if the Cardholder Name on the account is identical to the shipping Name used for this transaction. Values accepted: 01 = Account Name identical to shipping Name 02 = Account Name different than shipping Name
	ShipNameIndicator string `json:"shipNameIndicator,omitempty"`
	// Indicates whether the Merchant has experienced suspicious activity (including previous fraud) on the cardholder account. Values accepted: 01 = No suspicious activity has been observed 02 = Suspicious activity has been observed
	SuspiciousAccActivity string `json:"suspiciousAccActivity,omitempty"`
	// Number of transactions (successful and abandoned) for this cardholder account with the Merchant across all payment accounts in the previous 24 hours. Length: maximum 3 characters
	TxnActivityDay float64 `json:"txnActivityDay,omitempty"`
	// Number of transactions (successful and abandoned) for this cardholder account with the Merchant across all payment accounts in the previous year.
	TxnActivityYear float64 `json:"txnActivityYear,omitempty"`
}