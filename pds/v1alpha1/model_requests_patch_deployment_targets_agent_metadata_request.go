/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// RequestsPatchDeploymentTargetsAgentMetadataRequest struct for RequestsPatchDeploymentTargetsAgentMetadataRequest
type RequestsPatchDeploymentTargetsAgentMetadataRequest struct {
	KubeApiVersion *string `json:"kube_api_version,omitempty"`
	KubePlatform *string `json:"kube_platform,omitempty"`
	PxCsiEnabled *string `json:"px_csi_enabled,omitempty"`
	PxServiceNamespace *string `json:"px_service_namespace,omitempty"`
	PxVersion *string `json:"px_version,omitempty"`
}

// NewRequestsPatchDeploymentTargetsAgentMetadataRequest instantiates a new RequestsPatchDeploymentTargetsAgentMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsPatchDeploymentTargetsAgentMetadataRequest() *RequestsPatchDeploymentTargetsAgentMetadataRequest {
	this := RequestsPatchDeploymentTargetsAgentMetadataRequest{}
	return &this
}

// NewRequestsPatchDeploymentTargetsAgentMetadataRequestWithDefaults instantiates a new RequestsPatchDeploymentTargetsAgentMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsPatchDeploymentTargetsAgentMetadataRequestWithDefaults() *RequestsPatchDeploymentTargetsAgentMetadataRequest {
	this := RequestsPatchDeploymentTargetsAgentMetadataRequest{}
	return &this
}

// GetKubeApiVersion returns the KubeApiVersion field value if set, zero value otherwise.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetKubeApiVersion() string {
	if o == nil || o.KubeApiVersion == nil {
		var ret string
		return ret
	}
	return *o.KubeApiVersion
}

// GetKubeApiVersionOk returns a tuple with the KubeApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetKubeApiVersionOk() (*string, bool) {
	if o == nil || o.KubeApiVersion == nil {
		return nil, false
	}
	return o.KubeApiVersion, true
}

// HasKubeApiVersion returns a boolean if a field has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) HasKubeApiVersion() bool {
	if o != nil && o.KubeApiVersion != nil {
		return true
	}

	return false
}

// SetKubeApiVersion gets a reference to the given string and assigns it to the KubeApiVersion field.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) SetKubeApiVersion(v string) {
	o.KubeApiVersion = &v
}

// GetKubePlatform returns the KubePlatform field value if set, zero value otherwise.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetKubePlatform() string {
	if o == nil || o.KubePlatform == nil {
		var ret string
		return ret
	}
	return *o.KubePlatform
}

// GetKubePlatformOk returns a tuple with the KubePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetKubePlatformOk() (*string, bool) {
	if o == nil || o.KubePlatform == nil {
		return nil, false
	}
	return o.KubePlatform, true
}

// HasKubePlatform returns a boolean if a field has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) HasKubePlatform() bool {
	if o != nil && o.KubePlatform != nil {
		return true
	}

	return false
}

// SetKubePlatform gets a reference to the given string and assigns it to the KubePlatform field.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) SetKubePlatform(v string) {
	o.KubePlatform = &v
}

// GetPxCsiEnabled returns the PxCsiEnabled field value if set, zero value otherwise.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetPxCsiEnabled() string {
	if o == nil || o.PxCsiEnabled == nil {
		var ret string
		return ret
	}
	return *o.PxCsiEnabled
}

// GetPxCsiEnabledOk returns a tuple with the PxCsiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetPxCsiEnabledOk() (*string, bool) {
	if o == nil || o.PxCsiEnabled == nil {
		return nil, false
	}
	return o.PxCsiEnabled, true
}

// HasPxCsiEnabled returns a boolean if a field has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) HasPxCsiEnabled() bool {
	if o != nil && o.PxCsiEnabled != nil {
		return true
	}

	return false
}

// SetPxCsiEnabled gets a reference to the given string and assigns it to the PxCsiEnabled field.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) SetPxCsiEnabled(v string) {
	o.PxCsiEnabled = &v
}

// GetPxServiceNamespace returns the PxServiceNamespace field value if set, zero value otherwise.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetPxServiceNamespace() string {
	if o == nil || o.PxServiceNamespace == nil {
		var ret string
		return ret
	}
	return *o.PxServiceNamespace
}

// GetPxServiceNamespaceOk returns a tuple with the PxServiceNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetPxServiceNamespaceOk() (*string, bool) {
	if o == nil || o.PxServiceNamespace == nil {
		return nil, false
	}
	return o.PxServiceNamespace, true
}

// HasPxServiceNamespace returns a boolean if a field has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) HasPxServiceNamespace() bool {
	if o != nil && o.PxServiceNamespace != nil {
		return true
	}

	return false
}

// SetPxServiceNamespace gets a reference to the given string and assigns it to the PxServiceNamespace field.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) SetPxServiceNamespace(v string) {
	o.PxServiceNamespace = &v
}

// GetPxVersion returns the PxVersion field value if set, zero value otherwise.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetPxVersion() string {
	if o == nil || o.PxVersion == nil {
		var ret string
		return ret
	}
	return *o.PxVersion
}

// GetPxVersionOk returns a tuple with the PxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) GetPxVersionOk() (*string, bool) {
	if o == nil || o.PxVersion == nil {
		return nil, false
	}
	return o.PxVersion, true
}

// HasPxVersion returns a boolean if a field has been set.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) HasPxVersion() bool {
	if o != nil && o.PxVersion != nil {
		return true
	}

	return false
}

// SetPxVersion gets a reference to the given string and assigns it to the PxVersion field.
func (o *RequestsPatchDeploymentTargetsAgentMetadataRequest) SetPxVersion(v string) {
	o.PxVersion = &v
}

func (o RequestsPatchDeploymentTargetsAgentMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KubeApiVersion != nil {
		toSerialize["kube_api_version"] = o.KubeApiVersion
	}
	if o.KubePlatform != nil {
		toSerialize["kube_platform"] = o.KubePlatform
	}
	if o.PxCsiEnabled != nil {
		toSerialize["px_csi_enabled"] = o.PxCsiEnabled
	}
	if o.PxServiceNamespace != nil {
		toSerialize["px_service_namespace"] = o.PxServiceNamespace
	}
	if o.PxVersion != nil {
		toSerialize["px_version"] = o.PxVersion
	}
	return json.Marshal(toSerialize)
}

type NullableRequestsPatchDeploymentTargetsAgentMetadataRequest struct {
	value *RequestsPatchDeploymentTargetsAgentMetadataRequest
	isSet bool
}

func (v NullableRequestsPatchDeploymentTargetsAgentMetadataRequest) Get() *RequestsPatchDeploymentTargetsAgentMetadataRequest {
	return v.value
}

func (v *NullableRequestsPatchDeploymentTargetsAgentMetadataRequest) Set(val *RequestsPatchDeploymentTargetsAgentMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsPatchDeploymentTargetsAgentMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsPatchDeploymentTargetsAgentMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsPatchDeploymentTargetsAgentMetadataRequest(val *RequestsPatchDeploymentTargetsAgentMetadataRequest) *NullableRequestsPatchDeploymentTargetsAgentMetadataRequest {
	return &NullableRequestsPatchDeploymentTargetsAgentMetadataRequest{value: val, isSet: true}
}

func (v NullableRequestsPatchDeploymentTargetsAgentMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsPatchDeploymentTargetsAgentMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


