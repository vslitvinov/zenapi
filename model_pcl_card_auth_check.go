/*
 * ZEN eCommerce Public API
 *
 * ZEN allows you to accept multiple global and local payment methods (including credit cards, instant bank transfers, and many more). The following document will guide you through the integration process and introduce you to the main features of the Public API..<br><br><p><strong><span data-contrast=\"auto\">1. Allowed HTTPs requests:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">PU</span></strong><strong><span data-contrast=\"auto\">T</span></strong><strong><span data-contrast=\"auto\">:</span></strong><span data-contrast=\"auto\">&nbsp;To create a resource&nbsp;</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">POST:</span></strong><span data-contrast=\"auto\">&nbsp;To update ra esource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">GET:</span></strong><span data-contrast=\"auto\">&nbsp;To get a resource or a list of resources</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">DELETE:</span></strong><span data-contrast=\"auto\">&nbsp;To delete a resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ul><p><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span><strong><span data-contrast=\"auto\">2. Description&nbsp;</span></strong><strong><span data-contrast=\"auto\">o</span></strong><strong><span data-contrast=\"auto\">f&nbsp;</span></strong><strong><span data-contrast=\"auto\">u</span></strong><strong><span data-contrast=\"auto\">sual&nbsp;</span></strong><strong><span data-contrast=\"auto\">s</span></strong><strong><span data-contrast=\"auto\">erver&nbsp;</span></strong><strong><span data-contrast=\"auto\">r</span></strong><strong><span data-contrast=\"auto\">esponses:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">200</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">--</span></strong><strong><span data-contrast=\"auto\">&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">OK</span></strong><span data-contrast=\"auto\"> - the request was successful (some API calls may return 201 instead)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">201&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Created</span></strong><span data-contrast=\"auto\"> - the request was successful and a resource was created</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">204  </span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">No Content</span></strong><span data-contrast=\"auto\"> - the request was successful, but there is no representation to return (</span><span data-contrast=\"auto\">i</span><span data-contrast=\"auto\">.e. the response is empty)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"4\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">400&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">Bad Request</span></strong><span data-contrast=\"auto\"> - the request could not be understood or was missing required parameters</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"5\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">401&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Unauthorized</span></strong><span data-contrast=\"auto\"> - authentication failed or user does not have permissions for requested operation</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">403&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\">&nbsp;</span><strong><span data-contrast=\"auto\">Forbidden </span></strong><span data-contrast=\"auto\"> - access denied</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">404&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\"> </span><strong><span data-contrast=\"auto\">Not Found</span></strong><span data-contrast=\"auto\"> - the resource was not found</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">405  </span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">Method Not Allowed</span></strong><span data-contrast=\"auto\"> - requested method is not supported for the resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li></ul><p><strong><span data-contrast=\"auto\">3.&nbsp;</span></strong><strong><span data-contrast=\"auto\">API Integration Guidelines</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><span data-contrast=\"auto\">To start using ZEN eCommerce Public API please follow these instructions:&nbsp;</span><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:6,&quot;335551620&quot;:6,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#before-you-begin%20\">https://www.zen.com/developer/api-integration#before-you-begin&nbsp;</a><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559737&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ol><p><strong><span data-contrast=\"auto\">4. IPN API Secrets</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><p><span data-contrast=\"auto\">Instant Payment Notification (IPN) informs you about transaction status updates, e.g., about changing the status from pending to rejected. </span><br /><span data-contrast=\"auto\">Information abut IPN is available in the following link:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#ipn\"><span data-contrast=\"auto\">https://www.zen.com/developer/api-integration#ipn</span></a></p> <p><strong><span data-contrast=\"auto\">4. Error codes</span></strong><p><span data-contrast=\"auto\">The full list of error codes can be found here:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#rejection-codes\">https://www.zen.com/developer/api-integration#rejection-codes</a></p>  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PclCardAuthCheck struct {
	// Type of transaction
	Type_ string `json:"type"`
	ScaExemptions string `json:"scaExemptions,omitempty"`
	BrowserDetails *PclCardOneClickBrowserDetails `json:"browserDetails"`
	ThreeDsData *PclCardOneClickThreeDsData `json:"threeDsData,omitempty"`
	AuthenticationServiceTarget string `json:"authenticationServiceTarget,omitempty"`
	Subscription *PclCardOneClickSubscription `json:"subscription,omitempty"`
	// Return URL after 3DS authentiaction
	ReturnVerifyUrl string `json:"returnVerifyUrl,omitempty"`
	// Defines whether card should be saved for future purchases/recurring payments
	SaveCard bool `json:"saveCard,omitempty"`
	// Text that will appear on customer Bank Statement. Can be used only for Credit Card payments.
	Descriptor string `json:"descriptor"`
	Card *PclCardOneTimeCard `json:"card"`
}
