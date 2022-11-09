/*
 * ZEN eCommerce Public API
 *
 * ZEN allows you to accept multiple global and local payment methods (including credit cards, instant bank transfers, and many more). The following document will guide you through the integration process and introduce you to the main features of the Public API..<br><br><p><strong><span data-contrast=\"auto\">1. Allowed HTTPs requests:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">PU</span></strong><strong><span data-contrast=\"auto\">T</span></strong><strong><span data-contrast=\"auto\">:</span></strong><span data-contrast=\"auto\">&nbsp;To create a resource&nbsp;</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">POST:</span></strong><span data-contrast=\"auto\">&nbsp;To update ra esource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">GET:</span></strong><span data-contrast=\"auto\">&nbsp;To get a resource or a list of resources</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"10\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">DELETE:</span></strong><span data-contrast=\"auto\">&nbsp;To delete a resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ul><p><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span><strong><span data-contrast=\"auto\">2. Description&nbsp;</span></strong><strong><span data-contrast=\"auto\">o</span></strong><strong><span data-contrast=\"auto\">f&nbsp;</span></strong><strong><span data-contrast=\"auto\">u</span></strong><strong><span data-contrast=\"auto\">sual&nbsp;</span></strong><strong><span data-contrast=\"auto\">s</span></strong><strong><span data-contrast=\"auto\">erver&nbsp;</span></strong><strong><span data-contrast=\"auto\">r</span></strong><strong><span data-contrast=\"auto\">esponses:</span></strong><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559685&quot;:720,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><ul><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">200</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">--</span></strong><strong><span data-contrast=\"auto\">&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">OK</span></strong><span data-contrast=\"auto\"> - the request was successful (some API calls may return 201 instead)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">201&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Created</span></strong><span data-contrast=\"auto\"> - the request was successful and a resource was created</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">204  </span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">No Content</span></strong><span data-contrast=\"auto\"> - the request was successful, but there is no representation to return (</span><span data-contrast=\"auto\">i</span><span data-contrast=\"auto\">.e. the response is empty)</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"4\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">400&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;&nbsp;</span></strong><strong><span data-contrast=\"auto\">Bad Request</span></strong><span data-contrast=\"auto\"> - the request could not be understood or was missing required parameters</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"5\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">401&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">  Unauthorized</span></strong><span data-contrast=\"auto\"> - authentication failed or user does not have permissions for requested operation</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"1\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">403&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\">&nbsp;</span><strong><span data-contrast=\"auto\">Forbidden </span></strong><span data-contrast=\"auto\"> - access denied</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"2\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">404&nbsp;</span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><span data-contrast=\"auto\"> </span><strong><span data-contrast=\"auto\">Not Found</span></strong><span data-contrast=\"auto\"> - the resource was not found</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li><li data-leveltext=\"\" data-font=\"Symbol\" data-listid=\"11\" aria-setsize=\"-1\" data-aria-posinset=\"3\" data-aria-level=\"1\"><strong><span data-contrast=\"auto\">405  </span></strong><strong><span data-contrast=\"auto\">--&gt;</span></strong><strong><span data-contrast=\"auto\">&nbsp;</span></strong><strong><span data-contrast=\"auto\">Method Not Allowed</span></strong><span data-contrast=\"auto\"> - requested method is not supported for the resource</span><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:240}\">&nbsp;</span></li></ul><p><strong><span data-contrast=\"auto\">3.&nbsp;</span></strong><strong><span data-contrast=\"auto\">API Integration Guidelines</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><span data-contrast=\"auto\">To start using ZEN eCommerce Public API please follow these instructions:&nbsp;</span><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:6,&quot;335551620&quot;:6,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#before-you-begin%20\">https://www.zen.com/developer/api-integration#before-you-begin&nbsp;</a><span data-ccp-props=\"{&quot;134233279&quot;:true,&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559737&quot;:0,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></li></ol><p><strong><span data-contrast=\"auto\">4. IPN API Secrets</span></strong><span data-ccp-props=\"{&quot;201341983&quot;:0,&quot;335551550&quot;:1,&quot;335551620&quot;:1,&quot;335559685&quot;:360,&quot;335559739&quot;:160,&quot;335559740&quot;:259}\">&nbsp;</span></p><p><p><span data-contrast=\"auto\">Instant Payment Notification (IPN) informs you about transaction status updates, e.g., about changing the status from pending to rejected. </span><br /><span data-contrast=\"auto\">Information abut IPN is available in the following link:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#ipn\"><span data-contrast=\"auto\">https://www.zen.com/developer/api-integration#ipn</span></a></p> <p><strong><span data-contrast=\"auto\">4. Error codes</span></strong><p><span data-contrast=\"auto\">The full list of error codes can be found here:&nbsp;</span><br /><a href=\"https://www.zen.com/developer/api-integration#rejection-codes\">https://www.zen.com/developer/api-integration#rejection-codes</a></p>  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the ZEN eCommerce Public API API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	PaymentProfilesApi *PaymentProfilesApiService

	TerminalsApi *TerminalsApiService

	TransactionsApi *TransactionsApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.PaymentProfilesApi = (*PaymentProfilesApiService)(&c.common)
	c.TerminalsApi = (*TerminalsApiService)(&c.common)
	c.TransactionsApi = (*TransactionsApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
		if strings.Contains(contentType, "application/xml") {
			if err = xml.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		} else if strings.Contains(contentType, "application/json") {
			if err = json.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
