// Package gorequest inspired by Nodejs SuperAgent provides easy-way to write http client
package gorequest

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type Request *http.Request
type Response *http.Response

// HTTP methods we support
const (
	POST    = "POST"
	GET     = "GET"
	HEAD    = "HEAD"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
)

// Types we support.
const (
	TypeJSON       = "json"
	TypeXML        = "xml"
	TypeUrlencoded = "urlencoded"
	TypeForm       = "form"
	TypeFormData   = "form-data"
	TypeHTML       = "html"
	TypeText       = "text"
	TypeMultipart  = "multipart"
)

type superAgentRetryable struct {
	RetryableStatus []int
	RetryerTime     time.Duration
	RetryerCount    int
	Attempt         int
	Enable          bool
}

// A SuperAgent is a object storing all request data for client.
type SuperAgent struct {
	Url                  string
	Method               string
	Header               http.Header
	TargetType           string
	ForceType            string
	Data                 map[string]interface{}
	SliceData            []interface{}
	FormData             url.Values
	QueryData            url.Values
	FileData             []File
	BounceToRawString    bool
	RawString            string
	Client               *http.Client
	Transport            *http.Transport
	Cookies              []*http.Cookie
	Errors               []error
	BasicAuth            struct{ Username, Password string }
	Debug                bool
	CurlCommand          bool
	logger               Logger
	Retryable            superAgentRetryable
	DoNotClearSuperAgent bool
	isClone              bool
	context              context.Context
}

var DisableTransportSwap = false

// Used to create a new SuperAgent object.
func New() *SuperAgent { _ = "STUB: not implemented"; return nil }

// disable keep alives by default, see this issue https://github.com/parnurzeal/gorequest/issues/75

func cloneMapArray(old map[string][]string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func shallowCopyData(old map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func shallowCopyDataSlice(old []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func shallowCopyFileArray(old []File) []File { _ = "STUB: not implemented"; return nil }

func shallowCopyCookies(old []*http.Cookie) []*http.Cookie { _ = "STUB: not implemented"; return nil }

func shallowCopyErrors(old []error) []error { _ = "STUB: not implemented"; return nil }

// just need to change the array pointer?
func copyRetryable(old superAgentRetryable) superAgentRetryable {
	_ = "STUB: not implemented"
	return *new(superAgentRetryable)
}

// Returns a copy of this superagent. Useful if you want to reuse the client/settings
// concurrently.
// Note: This does a shallow copy of the parent. So you will need to be
// careful of Data provided
// Note: It also directly re-uses the client and transport. If you modify the Timeout,
// or RedirectPolicy on a clone, the clone will have a new http.client. It is recommended
// that the base request set your timeout and redirect polices, and no modification of
// the client or transport happen after cloning.
// Note: DoNotClearSuperAgent is forced to "true" after Clone
func (s *SuperAgent) Clone() *SuperAgent { _ = "STUB: not implemented"; return nil }

// thread safe.. anyway

// Enable the debug mode which logs request/response detail
func (s *SuperAgent) SetDebug(enable bool) *SuperAgent { _ = "STUB: not implemented"; return nil }

// Enable the curlcommand mode which display a CURL command line
func (s *SuperAgent) SetCurlCommand(enable bool) *SuperAgent { _ = "STUB: not implemented"; return nil }

// Enable the DoNotClear mode for not clearing super agent and reuse for the next request
func (s *SuperAgent) SetDoNotClearSuperAgent(enable bool) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SuperAgent) SetLogger(logger Logger) *SuperAgent { _ = "STUB: not implemented"; return nil }

// Clear SuperAgent data for another new request.
func (s *SuperAgent) ClearSuperAgent() { _ = "STUB: not implemented"; return }

// Just a wrapper to initialize SuperAgent instance by method string
func (s *SuperAgent) CustomMethod(method, targetUrl string) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SuperAgent) Get(targetUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

func (s *SuperAgent) Post(targetUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

func (s *SuperAgent) Head(targetUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

func (s *SuperAgent) Put(targetUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

func (s *SuperAgent) Delete(targetUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

func (s *SuperAgent) Patch(targetUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

func (s *SuperAgent) Options(targetUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

// Set is used for setting header fields,
// this will overwrite the existed values of Header through AppendHeader().
// Example. To set `Accept` as `application/json`
//
//	gorequest.New().
//	  Post("/gamelist").
//	  Set("Accept", "application/json").
//	  End()
func (s *SuperAgent) Set(param string, value string) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// AppendHeader is used for setting header fileds with multiple values,
// Example. To set `Accept` as `application/json, text/plain`
//
//	gorequest.New().
//	  Post("/gamelist").
//	  AppendHeader("Accept", "application/json").
//	  AppendHeader("Accept", "text/plain").
//	  End()
func (s *SuperAgent) AppendHeader(param string, value string) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// Retryable is used for setting a Retryer policy
// Example. To set Retryer policy with 5 seconds between each attempt.
//          3 max attempt.
//          And StatusBadRequest and StatusInternalServerError as RetryableStatus

// gorequest.New().
//
//	Post("/gamelist").
//	Retry(3, 5 * time.Second, http.StatusBadRequest, http.StatusInternalServerError).
//	End()
func (s *SuperAgent) Retry(retryerCount int, retryerTime time.Duration, statusCode ...int) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// SetBasicAuth sets the basic authentication header
// Example. To set the header for username "myuser" and password "mypass"
//
//	gorequest.New()
//	  Post("/gamelist").
//	  SetBasicAuth("myuser", "mypass").
//	  End()
func (s *SuperAgent) SetBasicAuth(username string, password string) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// AddCookie adds a cookie to the request. The behavior is the same as AddCookie on Request from net/http
func (s *SuperAgent) AddCookie(c *http.Cookie) *SuperAgent { _ = "STUB: not implemented"; return nil }

// AddCookies is a convenient method to add multiple cookies
func (s *SuperAgent) AddCookies(cookies []*http.Cookie) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

var Types = map[string]string{
	TypeJSON:       "application/json",
	TypeXML:        "application/xml",
	TypeForm:       "application/x-www-form-urlencoded",
	TypeFormData:   "application/x-www-form-urlencoded",
	TypeUrlencoded: "application/x-www-form-urlencoded",
	TypeHTML:       "text/html",
	TypeText:       "text/plain",
	TypeMultipart:  "multipart/form-data",
}

// Type is a convenience function to specify the data type to send.
// For example, to send data as `application/x-www-form-urlencoded` :
//
//	gorequest.New().
//	  Post("/recipe").
//	  Type("form").
//	  Send(`{ "name": "egg benedict", "category": "brunch" }`).
//	  End()
//
// This will POST the body "name=egg benedict&category=brunch" to url /recipe
//
// GoRequest supports
//
//	"text/html" uses "html"
//	"application/json" uses "json"
//	"application/xml" uses "xml"
//	"text/plain" uses "text"
//	"application/x-www-form-urlencoded" uses "urlencoded", "form" or "form-data"
func (s *SuperAgent) Type(typeStr string) *SuperAgent { _ = "STUB: not implemented"; return nil }

// Query function accepts either json string or strings which will form a query-string in url of GET method or body of POST method.
// For example, making "/search?query=bicycle&size=50x50&weight=20kg" using GET method:
//
//	gorequest.New().
//	  Get("/search").
//	  Query(`{ query: 'bicycle' }`).
//	  Query(`{ size: '50x50' }`).
//	  Query(`{ weight: '20kg' }`).
//	  End()
//
// Or you can put multiple json values:
//
//	gorequest.New().
//	  Get("/search").
//	  Query(`{ query: 'bicycle', size: '50x50', weight: '20kg' }`).
//	  End()
//
// Strings are also acceptable:
//
//	gorequest.New().
//	  Get("/search").
//	  Query("query=bicycle&size=50x50").
//	  Query("weight=20kg").
//	  End()
//
// Or even Mixed! :)
//
//	gorequest.New().
//	  Get("/search").
//	  Query("query=bicycle").
//	  Query(`{ size: '50x50', weight:'20kg' }`).
//	  End()
func (s *SuperAgent) Query(content interface{}) *SuperAgent { _ = "STUB: not implemented"; return nil }

func (s *SuperAgent) queryStruct(content interface{}) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SuperAgent) queryString(content string) *SuperAgent { _ = "STUB: not implemented"; return nil }

// TODO: need to check correct format of 'field=val&field=val&...'

func (s *SuperAgent) queryMap(content interface{}) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// As Go conventions accepts ; as a synonym for &. (https://github.com/golang/go/issues/2210)
// Thus, Query won't accept ; in a querystring if we provide something like fields=f1;f2;f3
// This Param is then created as an alternative method to solve this.
func (s *SuperAgent) Param(key string, value string) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// Set TLSClientConfig for underling Transport.
// One example is you can use it to disable security check (https):
//
//	gorequest.New().TLSClientConfig(&tls.Config{ InsecureSkipVerify: true}).
//	  Get("https://disable-security-check.com").
//	  End()
func (s *SuperAgent) TLSClientConfig(config *tls.Config) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SuperAgent) safeModifyTransport() { _ = "STUB: not implemented"; return }

// Note: Consider adding other fields from http.Transport if necessary for Go 1.24.3 compatibility

// Proxy function accepts a proxy url string to setup proxy url for any request.
// It provides a convenience way to setup proxy which have advantages over usual old ways.
// One example is you might try to set `http_proxy` environment. This means you are setting proxy up for all the requests.
// You will not be able to send different request with different proxy unless you change your `http_proxy` environment again.
// Another example is using Golang proxy setting. This is normal prefer way to do but too verbase compared to GoRequest's Proxy:
//
//	gorequest.New().Proxy("http://myproxy:9999").
//	  Post("http://www.google.com").
//	  End()
//
// To set no_proxy, just put empty string to Proxy func:
//
//	gorequest.New().Proxy("").
//	  Post("http://www.google.com").
//	  End()
func (s *SuperAgent) Proxy(proxyUrl string) *SuperAgent { _ = "STUB: not implemented"; return nil }

// RedirectPolicy accepts a function to define how to handle redirects. If the
// policy function returns an error, the next Request is not made and the previous
// request is returned.
//
// The policy function's arguments are the Request about to be made and the
// past requests in order of oldest first.
func (s *SuperAgent) RedirectPolicy(policy func(req Request, via []Request) error) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SuperAgent) safeModifyHttpClient() { _ = "STUB: not implemented"; return }

// Timeout sets the timeout for the HTTP client.
func (s *SuperAgent) Timeout(timeout time.Duration) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// Send function accepts either json string or query strings which is usually used to assign data to POST or PUT method.
// Without specifying any type, if you give Send with json data, you are doing requesting in json format:
//
//	gorequest.New().
//	  Post("/search").
//	  Send(`{ query: 'sushi' }`).
//	  End()
//
// While if you use at least one of querystring, GoRequest understands and automatically set the Content-Type to `application/x-www-form-urlencoded`
//
//	gorequest.New().
//	  Post("/search").
//	  Send("query=tonkatsu").
//	  End()
//
// So, if you want to strictly send json format, you need to use Type func to set it as `json` (Please see more details in Type function).
// You can also do multiple chain of Send:
//
//	gorequest.New().
//	  Post("/search").
//	  Send("query=bicycle&size=50x50").
//	  Send(`{ wheel: '4'}`).
//	  End()
//
// From v0.2.0, Send function provide another convenience way to work with Struct type. You can mix and match it with json and query string:
//
//	type BrowserVersionSupport struct {
//	  Chrome string
//	  Firefox string
//	}
//	ver := BrowserVersionSupport{ Chrome: "37.0.2041.6", Firefox: "30.0" }
//	gorequest.New().
//	  Post("/update_version").
//	  Send(ver).
//	  Send(`{"Safari":"5.1.10"}`).
//	  End()
//
// If you have set Type to text or Content-Type to text/plain, content will be sent as raw string in body instead of form
//
//	gorequest.New().
//	  Post("/greet").
//	  Type("text").
//	  Send("hello world").
//	  End()
func (s *SuperAgent) Send(content interface{}) *SuperAgent {
	_ = "STUB: not implemented"
	// TODO: add normal text mode or other mode to Send func
	return nil
}

// includes rune

// includes byte

// TODO: leave default for handling other types in the future, such as complex numbers, (nested) maps, etc

func makeSliceOfReflectValue(v reflect.Value) (slice []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// SendSlice (similar to SendString) returns SuperAgent's itself for any next chain and takes content []interface{} as a parameter.
// Its duty is to append slice of interface{} into s.SliceData ([]interface{}) which later changes into json array in the End() func.
func (s *SuperAgent) SendSlice(content []interface{}) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

func (s *SuperAgent) SendMap(content interface{}) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// SendStruct (similar to SendString) returns SuperAgent's itself for any next chain and takes content interface{} as a parameter.
// Its duty is to transfrom interface{} (implicitly always a struct) into s.Data (map[string]interface{}) which later changes into appropriate format such as json, form, text, etc. in the End() func.
func (s *SuperAgent) SendStruct(content interface{}) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// SendString returns SuperAgent's itself for any next chain and takes content string as a parameter.
// Its duty is to transform String into s.Data (map[string]interface{}) which later changes into appropriate format such as json, form, text, etc. in the End func.
// Send implicitly uses SendString and you should use Send instead of this.
func (s *SuperAgent) SendString(content string) *SuperAgent { _ = "STUB: not implemented"; return nil }

// add to SliceData

// bounce to rawstring if it is arrayjson, or others

// make it array if already have key

// check if previous data is one string or array

// make it just string if does not already have same key

// Dump all contents to RawString in case in the end user doesn't want json or form.

type File struct {
	Filename  string
	Fieldname string
	Data      []byte
}

// SendFile function works only with type "multipart". The function accepts one mandatory and up to two optional arguments. The mandatory (first) argument is the file.
// The function accepts a path to a file as string:
//
//	gorequest.New().
//	  Post("http://example.com").
//	  Type("multipart").
//	  SendFile("./example_file.ext").
//	  End()
//
// File can also be a []byte slice of a already file read by eg. ioutil.ReadFile:
//
//	b, _ := ioutil.ReadFile("./example_file.ext")
//	gorequest.New().
//	  Post("http://example.com").
//	  Type("multipart").
//	  SendFile(b).
//	  End()
//
// Furthermore file can also be a os.File:
//
//	f, _ := os.Open("./example_file.ext")
//	gorequest.New().
//	  Post("http://example.com").
//	  Type("multipart").
//	  SendFile(f).
//	  End()
//
// The first optional argument (second argument overall) is the filename, which will be automatically determined when file is a string (path) or a os.File.
// When file is a []byte slice, filename defaults to "filename". In all cases the automatically determined filename can be overwritten:
//
//	b, _ := ioutil.ReadFile("./example_file.ext")
//	gorequest.New().
//	  Post("http://example.com").
//	  Type("multipart").
//	  SendFile(b, "my_custom_filename").
//	  End()
//
// The second optional argument (third argument overall) is the fieldname in the multipart/form-data request. It defaults to fileNUMBER (eg. file1), where number is ascending and starts counting at 1.
// So if you send multiple files, the fieldnames will be file1, file2, ... unless it is overwritten. If fieldname is set to "file" it will be automatically set to fileNUMBER, where number is the greatest exsiting number+1 unless
// a third argument skipFileNumbering is provided and true.
//
//	b, _ := ioutil.ReadFile("./example_file.ext")
//	gorequest.New().
//	  Post("http://example.com").
//	  Type("multipart").
//	  SendFile(b, "", "my_custom_fieldname"). // filename left blank, will become "example_file.ext"
//	  End()
func (s *SuperAgent) SendFile(file interface{}, args ...interface{}) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

func changeMapToURLValues(data map[string]interface{}) url.Values {
	_ = "STUB: not implemented"
	return *new(url.Values)
}

// if a number, change to string
// json.Number used to protect against a wrong (for GoRequest) default conversion
// which always converts number to float64.
// This type is caused by using Decoder.UseNumber()

// TODO add all other int-Types (int8, int16, ...)

// following slices are mostly needed for tests

// these slices are used in practice like sending a struct

// TODO add ptr, arrays, ...

// End is the most important function that you need to call when ending the chain. The request won't proceed without calling it.
// End function returns Response which matchs the structure of Response type in Golang's http package (but without Body data). The body data itself returns as a string in a 2nd return value.
// Lastly but worth noticing, error array (NOTE: not just single error value) is returned as a 3rd value and nil otherwise.
//
// For example:
//
//	resp, body, errs := gorequest.New().Get("http://www.google.com").End()
//	if errs != nil {
//	  fmt.Println(errs)
//	}
//	fmt.Println(resp, body)
//
// Moreover, End function also supports callback which you can put as a parameter.
// This extends the flexibility and makes GoRequest fun and clean! You can use GoRequest in whatever style you love!
//
// For example:
//
//	func printBody(resp gorequest.Response, body string, errs []error){
//	  fmt.Println(resp.Status)
//	}
//	gorequest.New().Get("http://www..google.com").End(printBody)
func (s *SuperAgent) End(callback ...func(response Response, body string, errs []error)) (Response, string, []error) {
	_ = "STUB: not implemented"
	return *new(Response), "", nil
}

// EndBytes should be used when you want the body as bytes. The callbacks work the same way as with `End`, except that a byte array is used instead of a string.
func (s *SuperAgent) EndBytes(callback ...func(response Response, body []byte, errs []error)) (Response, []byte, []error) {
	_ = "STUB: not implemented"
	return *new(Response), nil, nil
}

func (s *SuperAgent) isRetryableRequest(resp Response) bool {
	_ = "STUB: not implemented"
	return false
}

func contains(respStatus int, statuses []int) bool { _ = "STUB: not implemented"; return false }

func (s *SuperAgent) Context(ctx context.Context) *SuperAgent {
	_ = "STUB: not implemented"
	return nil
}

// EndStruct should be used when you want the body as a struct. The callbacks work the same way as with `End`, except that a struct is used instead of a string.
func (s *SuperAgent) EndStruct(v interface{}, callback ...func(response Response, v interface{}, body []byte, errs []error)) (Response, []byte, []error) {
	_ = "STUB: not implemented"
	return *new(Response), nil, nil
}

func (s *SuperAgent) getResponseBytes() (Response, []byte, []error) {
	_ = "STUB: not implemented"
	return *new(Response), nil, nil
}

// check whether there is an error. if yes, return all errors

// check if there is forced type

// If forcetype is not set, check whether user set Content-Type header.
// If yes, also bounce to the correct supported TargetType automatically.

// if slice and map get mixed, let's bounce to rawstring

// Make Request

// Set Transport

// Log details of this request

// Display CURL command line

// Send request

// Log details of this response

// Reset resp.Body so it can be use again

func (s *SuperAgent) MakeRequest() (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is only set when the request body content is non-empty.

// !!! Important Note !!!
//
// Throughout this region, contentReader and contentType are only set when
// the contents will be non-empty.
// This is done avoid ever sending a non-nil request body with nil contents
// to http.NewRequest, because it contains logic which dependends on
// whether or not the body is "nil".
//
// See PR #136 for more information:
//
//     https://github.com/parnurzeal/gorequest/pull/136
//

// If-case to give support to json array. we check if
// 1) Map only: send it as json map from s.Data
// 2) Array or Mix of map & array or others: send it as rawstring from s.RawString

// copied from CreateFormField() in mime/multipart/writer.go

// add the files

// close before call to FormDataContentType ! otherwise its not valid multipart

// let's return an error instead of an nil pointer exception here

// Setting the Host header is a special case, see this issue: https://github.com/golang/go/issues/7682

// https://github.com/parnurzeal/gorequest/issues/164
// Don't infer the content type header if an overrride is already provided.

// Add all querystring from Query func

// Add basic auth

// Add cookies

// AsCurlCommand returns a string representing the runnable `curl' command
// version of the request.
func (s *SuperAgent) AsCurlCommand() (string, error) { _ = "STUB: not implemented"; return "", nil }
