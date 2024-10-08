// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Application defines model for Application.
type Application struct {
	CreatedAt          *string `json:"createdAt,omitempty"`
	Cron               *string `json:"cron,omitempty"`
	HasCI              *bool   `json:"hasCI,omitempty"`
	Id                 *string `json:"id,omitempty"`
	LastSuccessfulScan *Scan   `json:"lastSuccessfulScan,omitempty"`
	Name               *string `json:"name,omitempty"`
	Scans              *[]Scan `json:"scans,omitempty"`
	Url                *string `json:"url,omitempty"`
}

// IntrospectionResponse defines model for IntrospectionResponse.
type IntrospectionResponse struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	Id            *string `json:"id,omitempty"`
}

// Scan defines model for Scan.
type Scan struct {
	Alerts          *[]map[string]interface{} `json:"alerts,omitempty"`
	Application     *Application              `json:"application,omitempty"`
	CommitHash      *string                   `json:"commitHash,omitempty"`
	CompletionRatio *float32                  `json:"completionRatio,omitempty"`
	Configuration   *map[string]interface{}   `json:"configuration,omitempty"`
	CreatedAt       *string                   `json:"createdAt,omitempty"`
	Id              *string                   `json:"id,omitempty"`
	Score           *float32                  `json:"score,omitempty"`
	Status          *string                   `json:"status,omitempty"`
}

// PostApplicationApplicationIdStartScanJSONBody defines parameters for PostApplicationApplicationIdStartScan.
type PostApplicationApplicationIdStartScanJSONBody struct {
	// CommitHash See the commit identification section.
	CommitHash *string `json:"commitHash,omitempty"`

	// ConfigurationOverride See the configuration override section.
	ConfigurationOverride *string `json:"configurationOverride,omitempty"`

	// Introspection The stringified JSON introspection. See the introspection update section.
	Introspection *string `json:"introspection,omitempty"`
}

// PostApplicationsApplicationIdUploadIntrospectionJSONBody defines parameters for PostApplicationsApplicationIdUploadIntrospection.
type PostApplicationsApplicationIdUploadIntrospectionJSONBody struct {
	// IntrospectionResponse The stringified JSON introspection. See the introspection update section.
	IntrospectionResponse string `json:"introspectionResponse"`
}

// PostApplicationApplicationIdStartScanJSONRequestBody defines body for PostApplicationApplicationIdStartScan for application/json ContentType.
type PostApplicationApplicationIdStartScanJSONRequestBody PostApplicationApplicationIdStartScanJSONBody

// PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody defines body for PostApplicationsApplicationIdUploadIntrospection for application/json ContentType.
type PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody PostApplicationsApplicationIdUploadIntrospectionJSONBody

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
	// PostApplicationApplicationIdStartScanWithBody request with any body
	PostApplicationApplicationIdStartScanWithBody(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApplicationApplicationIdStartScan(ctx context.Context, applicationId string, body PostApplicationApplicationIdStartScanJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApplicationsApplicationIdUploadIntrospectionWithBody request with any body
	PostApplicationsApplicationIdUploadIntrospectionWithBody(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApplicationsApplicationIdUploadIntrospection(ctx context.Context, applicationId string, body PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetOrganizationOrganizationIdApplications request
	GetOrganizationOrganizationIdApplications(ctx context.Context, organizationId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetScansScanId request
	GetScansScanId(ctx context.Context, scanId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostApplicationApplicationIdStartScanWithBody(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApplicationApplicationIdStartScanRequestWithBody(c.Server, applicationId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApplicationApplicationIdStartScan(ctx context.Context, applicationId string, body PostApplicationApplicationIdStartScanJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApplicationApplicationIdStartScanRequest(c.Server, applicationId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApplicationsApplicationIdUploadIntrospectionWithBody(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApplicationsApplicationIdUploadIntrospectionRequestWithBody(c.Server, applicationId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApplicationsApplicationIdUploadIntrospection(ctx context.Context, applicationId string, body PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApplicationsApplicationIdUploadIntrospectionRequest(c.Server, applicationId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetOrganizationOrganizationIdApplications(ctx context.Context, organizationId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetOrganizationOrganizationIdApplicationsRequest(c.Server, organizationId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetScansScanId(ctx context.Context, scanId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetScansScanIdRequest(c.Server, scanId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostApplicationApplicationIdStartScanRequest calls the generic PostApplicationApplicationIdStartScan builder with application/json body
func NewPostApplicationApplicationIdStartScanRequest(server string, applicationId string, body PostApplicationApplicationIdStartScanJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApplicationApplicationIdStartScanRequestWithBody(server, applicationId, "application/json", bodyReader)
}

// NewPostApplicationApplicationIdStartScanRequestWithBody generates requests for PostApplicationApplicationIdStartScan with any type of body
func NewPostApplicationApplicationIdStartScanRequestWithBody(server string, applicationId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "applicationId", runtime.ParamLocationPath, applicationId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/application/%s/start-scan", pathParam0)
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

// NewPostApplicationsApplicationIdUploadIntrospectionRequest calls the generic PostApplicationsApplicationIdUploadIntrospection builder with application/json body
func NewPostApplicationsApplicationIdUploadIntrospectionRequest(server string, applicationId string, body PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApplicationsApplicationIdUploadIntrospectionRequestWithBody(server, applicationId, "application/json", bodyReader)
}

// NewPostApplicationsApplicationIdUploadIntrospectionRequestWithBody generates requests for PostApplicationsApplicationIdUploadIntrospection with any type of body
func NewPostApplicationsApplicationIdUploadIntrospectionRequestWithBody(server string, applicationId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "applicationId", runtime.ParamLocationPath, applicationId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/applications/%s/upload-introspection", pathParam0)
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

// NewGetOrganizationOrganizationIdApplicationsRequest generates requests for GetOrganizationOrganizationIdApplications
func NewGetOrganizationOrganizationIdApplicationsRequest(server string, organizationId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "organizationId", runtime.ParamLocationPath, organizationId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/organization/%s/applications", pathParam0)
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

// NewGetScansScanIdRequest generates requests for GetScansScanId
func NewGetScansScanIdRequest(server string, scanId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "scanId", runtime.ParamLocationPath, scanId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/scans/%s", pathParam0)
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
	// PostApplicationApplicationIdStartScanWithBodyWithResponse request with any body
	PostApplicationApplicationIdStartScanWithBodyWithResponse(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApplicationApplicationIdStartScanResponse, error)

	PostApplicationApplicationIdStartScanWithResponse(ctx context.Context, applicationId string, body PostApplicationApplicationIdStartScanJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApplicationApplicationIdStartScanResponse, error)

	// PostApplicationsApplicationIdUploadIntrospectionWithBodyWithResponse request with any body
	PostApplicationsApplicationIdUploadIntrospectionWithBodyWithResponse(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApplicationsApplicationIdUploadIntrospectionResponse, error)

	PostApplicationsApplicationIdUploadIntrospectionWithResponse(ctx context.Context, applicationId string, body PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApplicationsApplicationIdUploadIntrospectionResponse, error)

	// GetOrganizationOrganizationIdApplicationsWithResponse request
	GetOrganizationOrganizationIdApplicationsWithResponse(ctx context.Context, organizationId string, reqEditors ...RequestEditorFn) (*GetOrganizationOrganizationIdApplicationsResponse, error)

	// GetScansScanIdWithResponse request
	GetScansScanIdWithResponse(ctx context.Context, scanId string, reqEditors ...RequestEditorFn) (*GetScansScanIdResponse, error)
}

type PostApplicationApplicationIdStartScanResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
}

// Status returns HTTPResponse.Status
func (r PostApplicationApplicationIdStartScanResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApplicationApplicationIdStartScanResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApplicationsApplicationIdUploadIntrospectionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *IntrospectionResponse
}

// Status returns HTTPResponse.Status
func (r PostApplicationsApplicationIdUploadIntrospectionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApplicationsApplicationIdUploadIntrospectionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOrganizationOrganizationIdApplicationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Application
}

// Status returns HTTPResponse.Status
func (r GetOrganizationOrganizationIdApplicationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOrganizationOrganizationIdApplicationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetScansScanIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Scan
}

// Status returns HTTPResponse.Status
func (r GetScansScanIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetScansScanIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostApplicationApplicationIdStartScanWithBodyWithResponse request with arbitrary body returning *PostApplicationApplicationIdStartScanResponse
func (c *ClientWithResponses) PostApplicationApplicationIdStartScanWithBodyWithResponse(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApplicationApplicationIdStartScanResponse, error) {
	rsp, err := c.PostApplicationApplicationIdStartScanWithBody(ctx, applicationId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApplicationApplicationIdStartScanResponse(rsp)
}

func (c *ClientWithResponses) PostApplicationApplicationIdStartScanWithResponse(ctx context.Context, applicationId string, body PostApplicationApplicationIdStartScanJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApplicationApplicationIdStartScanResponse, error) {
	rsp, err := c.PostApplicationApplicationIdStartScan(ctx, applicationId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApplicationApplicationIdStartScanResponse(rsp)
}

// PostApplicationsApplicationIdUploadIntrospectionWithBodyWithResponse request with arbitrary body returning *PostApplicationsApplicationIdUploadIntrospectionResponse
func (c *ClientWithResponses) PostApplicationsApplicationIdUploadIntrospectionWithBodyWithResponse(ctx context.Context, applicationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApplicationsApplicationIdUploadIntrospectionResponse, error) {
	rsp, err := c.PostApplicationsApplicationIdUploadIntrospectionWithBody(ctx, applicationId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApplicationsApplicationIdUploadIntrospectionResponse(rsp)
}

func (c *ClientWithResponses) PostApplicationsApplicationIdUploadIntrospectionWithResponse(ctx context.Context, applicationId string, body PostApplicationsApplicationIdUploadIntrospectionJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApplicationsApplicationIdUploadIntrospectionResponse, error) {
	rsp, err := c.PostApplicationsApplicationIdUploadIntrospection(ctx, applicationId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApplicationsApplicationIdUploadIntrospectionResponse(rsp)
}

// GetOrganizationOrganizationIdApplicationsWithResponse request returning *GetOrganizationOrganizationIdApplicationsResponse
func (c *ClientWithResponses) GetOrganizationOrganizationIdApplicationsWithResponse(ctx context.Context, organizationId string, reqEditors ...RequestEditorFn) (*GetOrganizationOrganizationIdApplicationsResponse, error) {
	rsp, err := c.GetOrganizationOrganizationIdApplications(ctx, organizationId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetOrganizationOrganizationIdApplicationsResponse(rsp)
}

// GetScansScanIdWithResponse request returning *GetScansScanIdResponse
func (c *ClientWithResponses) GetScansScanIdWithResponse(ctx context.Context, scanId string, reqEditors ...RequestEditorFn) (*GetScansScanIdResponse, error) {
	rsp, err := c.GetScansScanId(ctx, scanId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetScansScanIdResponse(rsp)
}

// ParsePostApplicationApplicationIdStartScanResponse parses an HTTP response from a PostApplicationApplicationIdStartScanWithResponse call
func ParsePostApplicationApplicationIdStartScanResponse(rsp *http.Response) (*PostApplicationApplicationIdStartScanResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApplicationApplicationIdStartScanResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostApplicationsApplicationIdUploadIntrospectionResponse parses an HTTP response from a PostApplicationsApplicationIdUploadIntrospectionWithResponse call
func ParsePostApplicationsApplicationIdUploadIntrospectionResponse(rsp *http.Response) (*PostApplicationsApplicationIdUploadIntrospectionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApplicationsApplicationIdUploadIntrospectionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest IntrospectionResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetOrganizationOrganizationIdApplicationsResponse parses an HTTP response from a GetOrganizationOrganizationIdApplicationsWithResponse call
func ParseGetOrganizationOrganizationIdApplicationsResponse(rsp *http.Response) (*GetOrganizationOrganizationIdApplicationsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetOrganizationOrganizationIdApplicationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Application
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetScansScanIdResponse parses an HTTP response from a GetScansScanIdWithResponse call
func ParseGetScansScanIdResponse(rsp *http.Response) (*GetScansScanIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetScansScanIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Scan
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
