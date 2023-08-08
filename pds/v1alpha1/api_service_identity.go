/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// ServiceIdentityApiService ServiceIdentityApi service
type ServiceIdentityApiService service

type ApiApiAccountsIdServiceIdentityGetRequest struct {
	ctx context.Context
	ApiService *ServiceIdentityApiService
	id string
	sortBy *string
	limit *string
	continuation *string
	id2 *string
	name *string
	createdBy *string
	clientId *string
	enabled *bool
}

// A given Service Identity attribute to sort results by (one of: id, name, created_at, created_by)
func (r ApiApiAccountsIdServiceIdentityGetRequest) SortBy(sortBy string) ApiApiAccountsIdServiceIdentityGetRequest {
	r.sortBy = &sortBy
	return r
}
// Maximum number of rows to return (could be less)
func (r ApiApiAccountsIdServiceIdentityGetRequest) Limit(limit string) ApiApiAccountsIdServiceIdentityGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiAccountsIdServiceIdentityGetRequest) Continuation(continuation string) ApiApiAccountsIdServiceIdentityGetRequest {
	r.continuation = &continuation
	return r
}
// Filter results by service identity id
func (r ApiApiAccountsIdServiceIdentityGetRequest) Id2(id2 string) ApiApiAccountsIdServiceIdentityGetRequest {
	r.id2 = &id2
	return r
}
// Filter results by name
func (r ApiApiAccountsIdServiceIdentityGetRequest) Name(name string) ApiApiAccountsIdServiceIdentityGetRequest {
	r.name = &name
	return r
}
// Filter results by created_by
func (r ApiApiAccountsIdServiceIdentityGetRequest) CreatedBy(createdBy string) ApiApiAccountsIdServiceIdentityGetRequest {
	r.createdBy = &createdBy
	return r
}
// Filter results by client_id
func (r ApiApiAccountsIdServiceIdentityGetRequest) ClientId(clientId string) ApiApiAccountsIdServiceIdentityGetRequest {
	r.clientId = &clientId
	return r
}
// Filter results by enabled
func (r ApiApiAccountsIdServiceIdentityGetRequest) Enabled(enabled bool) ApiApiAccountsIdServiceIdentityGetRequest {
	r.enabled = &enabled
	return r
}

func (r ApiApiAccountsIdServiceIdentityGetRequest) Execute() (*ModelsPaginatedResultModelsServiceIdentity, *http.Response, error) {
	return r.ApiService.ApiAccountsIdServiceIdentityGetExecute(r)
}

/*
ApiAccountsIdServiceIdentityGet List Service Identity

Lists all ServiceIdentity for an account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Account ID (must be valid UUID)
 @return ApiApiAccountsIdServiceIdentityGetRequest
*/
func (a *ServiceIdentityApiService) ApiAccountsIdServiceIdentityGet(ctx context.Context, id string) ApiApiAccountsIdServiceIdentityGetRequest {
	return ApiApiAccountsIdServiceIdentityGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsPaginatedResultModelsServiceIdentity
func (a *ServiceIdentityApiService) ApiAccountsIdServiceIdentityGetExecute(r ApiApiAccountsIdServiceIdentityGetRequest) (*ModelsPaginatedResultModelsServiceIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsPaginatedResultModelsServiceIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceIdentityApiService.ApiAccountsIdServiceIdentityGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts/{id}/service-identity"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.id2 != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id2, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.createdBy != nil {
		localVarQueryParams.Add("created_by", parameterToString(*r.createdBy, ""))
	}
	if r.clientId != nil {
		localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	}
	if r.enabled != nil {
		localVarQueryParams.Add("enabled", parameterToString(*r.enabled, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiAccountsIdServiceIdentityPostRequest struct {
	ctx context.Context
	ApiService *ServiceIdentityApiService
	id string
	body *RequestsServiceIdentityRequest
}

// Request body containing the service identity details.
func (r ApiApiAccountsIdServiceIdentityPostRequest) Body(body RequestsServiceIdentityRequest) ApiApiAccountsIdServiceIdentityPostRequest {
	r.body = &body
	return r
}

func (r ApiApiAccountsIdServiceIdentityPostRequest) Execute() (*ModelsServiceIdentityWithToken, *http.Response, error) {
	return r.ApiService.ApiAccountsIdServiceIdentityPostExecute(r)
}

/*
ApiAccountsIdServiceIdentityPost Create a Service Identity

Create a Service Identity for programmatic access to pds

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Account ID (must be valid UUID)
 @return ApiApiAccountsIdServiceIdentityPostRequest
*/
func (a *ServiceIdentityApiService) ApiAccountsIdServiceIdentityPost(ctx context.Context, id string) ApiApiAccountsIdServiceIdentityPostRequest {
	return ApiApiAccountsIdServiceIdentityPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsServiceIdentityWithToken
func (a *ServiceIdentityApiService) ApiAccountsIdServiceIdentityPostExecute(r ApiApiAccountsIdServiceIdentityPostRequest) (*ModelsServiceIdentityWithToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsServiceIdentityWithToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceIdentityApiService.ApiAccountsIdServiceIdentityPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/accounts/{id}/service-identity"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiServiceIdentityIdDeleteRequest struct {
	ctx context.Context
	ApiService *ServiceIdentityApiService
	id string
}


func (r ApiApiServiceIdentityIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiServiceIdentityIdDeleteExecute(r)
}

/*
ApiServiceIdentityIdDelete Delete service identity

Delete service identity by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Service Identity ID (must be valid UUID)
 @return ApiApiServiceIdentityIdDeleteRequest
*/
func (a *ServiceIdentityApiService) ApiServiceIdentityIdDelete(ctx context.Context, id string) ApiApiServiceIdentityIdDeleteRequest {
	return ApiApiServiceIdentityIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ServiceIdentityApiService) ApiServiceIdentityIdDeleteExecute(r ApiApiServiceIdentityIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceIdentityApiService.ApiServiceIdentityIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service-identity/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiServiceIdentityIdGetRequest struct {
	ctx context.Context
	ApiService *ServiceIdentityApiService
	id string
}


func (r ApiApiServiceIdentityIdGetRequest) Execute() (*ModelsServiceIdentity, *http.Response, error) {
	return r.ApiService.ApiServiceIdentityIdGetExecute(r)
}

/*
ApiServiceIdentityIdGet Get service identity by ID

Get service identity by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Service Identity ID (must be valid UUID)
 @return ApiApiServiceIdentityIdGetRequest
*/
func (a *ServiceIdentityApiService) ApiServiceIdentityIdGet(ctx context.Context, id string) ApiApiServiceIdentityIdGetRequest {
	return ApiApiServiceIdentityIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsServiceIdentity
func (a *ServiceIdentityApiService) ApiServiceIdentityIdGetExecute(r ApiApiServiceIdentityIdGetRequest) (*ModelsServiceIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsServiceIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceIdentityApiService.ApiServiceIdentityIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service-identity/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiServiceIdentityregenerateIdGetRequest struct {
	ctx context.Context
	ApiService *ServiceIdentityApiService
	id string
}


func (r ApiApiServiceIdentityregenerateIdGetRequest) Execute() (*ModelsServiceIdentityWithToken, *http.Response, error) {
	return r.ApiService.ApiServiceIdentityregenerateIdGetExecute(r)
}

/*
ApiServiceIdentityregenerateIdGet Regenerate service identity credentials

Revokes existing credentials and regenerate new ClientID and ClientToken

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Service Identity ID (must be valid UUID)
 @return ApiApiServiceIdentityregenerateIdGetRequest
*/
func (a *ServiceIdentityApiService) ApiServiceIdentityregenerateIdGet(ctx context.Context, id string) ApiApiServiceIdentityregenerateIdGetRequest {
	return ApiApiServiceIdentityregenerateIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsServiceIdentityWithToken
func (a *ServiceIdentityApiService) ApiServiceIdentityregenerateIdGetExecute(r ApiApiServiceIdentityregenerateIdGetRequest) (*ModelsServiceIdentityWithToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsServiceIdentityWithToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceIdentityApiService.ApiServiceIdentityregenerateIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/service-identity:regenerate/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServiceIdentityGenerateTokenPostRequest struct {
	ctx context.Context
	ApiService *ServiceIdentityApiService
	body *ControllersGenerateTokenRequest
}

// Request body containing the client id and client token.
func (r ApiServiceIdentityGenerateTokenPostRequest) Body(body ControllersGenerateTokenRequest) ApiServiceIdentityGenerateTokenPostRequest {
	r.body = &body
	return r
}

func (r ApiServiceIdentityGenerateTokenPostRequest) Execute() (*ControllersGenerateTokenResponse, *http.Response, error) {
	return r.ApiService.ServiceIdentityGenerateTokenPostExecute(r)
}

/*
ServiceIdentityGenerateTokenPost Generate JWT token for service identity

Generates JWT token for a service identity using client id and client token

**Authenticating for later requests using the ID token:**
`Authorization: Bearer {id_token}`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServiceIdentityGenerateTokenPostRequest
*/
func (a *ServiceIdentityApiService) ServiceIdentityGenerateTokenPost(ctx context.Context) ApiServiceIdentityGenerateTokenPostRequest {
	return ApiServiceIdentityGenerateTokenPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ControllersGenerateTokenResponse
func (a *ServiceIdentityApiService) ServiceIdentityGenerateTokenPostExecute(r ApiServiceIdentityGenerateTokenPostRequest) (*ControllersGenerateTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersGenerateTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceIdentityApiService.ServiceIdentityGenerateTokenPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service-identity/generate-token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
