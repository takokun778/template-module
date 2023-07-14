// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// V1HealthAPI request
	V1HealthAPI(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// V1HealthCache request
	V1HealthCache(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// V1HealthDB request
	V1HealthDB(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// V1Hello request with any body
	V1HelloWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	V1Hello(ctx context.Context, body V1HelloJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) V1HealthAPI(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1HealthAPIRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1HealthCache(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1HealthCacheRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1HealthDB(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1HealthDBRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1HelloWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1HelloRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) V1Hello(ctx context.Context, body V1HelloJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewV1HelloRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewV1HealthAPIRequest generates requests for V1HealthAPI
func NewV1HealthAPIRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/health/api")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewV1HealthCacheRequest generates requests for V1HealthCache
func NewV1HealthCacheRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/health/cache")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewV1HealthDBRequest generates requests for V1HealthDB
func NewV1HealthDBRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/health/db")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewV1HelloRequest calls the generic V1Hello builder with application/json body
func NewV1HelloRequest(server string, body V1HelloJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewV1HelloRequestWithBody(server, "application/json", bodyReader)
}

// NewV1HelloRequestWithBody generates requests for V1Hello with any type of body
func NewV1HelloRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/hello")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// V1HealthAPI request
	V1HealthAPIWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthAPIResponse, error)

	// V1HealthCache request
	V1HealthCacheWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthCacheResponse, error)

	// V1HealthDB request
	V1HealthDBWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthDBResponse, error)

	// V1Hello request with any body
	V1HelloWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*V1HelloResponse, error)

	V1HelloWithResponse(ctx context.Context, body V1HelloJSONRequestBody, reqEditors ...RequestEditorFn) (*V1HelloResponse, error)
}

type V1HealthAPIResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r V1HealthAPIResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r V1HealthAPIResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type V1HealthCacheResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r V1HealthCacheResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r V1HealthCacheResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type V1HealthDBResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r V1HealthDBResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r V1HealthDBResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type V1HelloResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *V1HelloResponseSchema
}

// Status returns HTTPResponse.Status
func (r V1HelloResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r V1HelloResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// V1HealthAPIWithResponse request returning *V1HealthAPIResponse
func (c *ClientWithResponses) V1HealthAPIWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthAPIResponse, error) {
	rsp, err := c.V1HealthAPI(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1HealthAPIResponse(rsp)
}

// V1HealthCacheWithResponse request returning *V1HealthCacheResponse
func (c *ClientWithResponses) V1HealthCacheWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthCacheResponse, error) {
	rsp, err := c.V1HealthCache(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1HealthCacheResponse(rsp)
}

// V1HealthDBWithResponse request returning *V1HealthDBResponse
func (c *ClientWithResponses) V1HealthDBWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*V1HealthDBResponse, error) {
	rsp, err := c.V1HealthDB(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1HealthDBResponse(rsp)
}

// V1HelloWithBodyWithResponse request with arbitrary body returning *V1HelloResponse
func (c *ClientWithResponses) V1HelloWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*V1HelloResponse, error) {
	rsp, err := c.V1HelloWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1HelloResponse(rsp)
}

func (c *ClientWithResponses) V1HelloWithResponse(ctx context.Context, body V1HelloJSONRequestBody, reqEditors ...RequestEditorFn) (*V1HelloResponse, error) {
	rsp, err := c.V1Hello(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseV1HelloResponse(rsp)
}

// ParseV1HealthAPIResponse parses an HTTP response from a V1HealthAPIWithResponse call
func ParseV1HealthAPIResponse(rsp *http.Response) (*V1HealthAPIResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &V1HealthAPIResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseV1HealthCacheResponse parses an HTTP response from a V1HealthCacheWithResponse call
func ParseV1HealthCacheResponse(rsp *http.Response) (*V1HealthCacheResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &V1HealthCacheResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseV1HealthDBResponse parses an HTTP response from a V1HealthDBWithResponse call
func ParseV1HealthDBResponse(rsp *http.Response) (*V1HealthDBResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &V1HealthDBResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseV1HelloResponse parses an HTTP response from a V1HelloWithResponse call
func ParseV1HelloResponse(rsp *http.Response) (*V1HelloResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &V1HelloResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest V1HelloResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}