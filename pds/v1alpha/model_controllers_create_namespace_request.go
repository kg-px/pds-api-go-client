/*
PDS API

Portworx Data Services API Server

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ControllersCreateNamespaceRequest struct for ControllersCreateNamespaceRequest
type ControllersCreateNamespaceRequest struct {
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewControllersCreateNamespaceRequest instantiates a new ControllersCreateNamespaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersCreateNamespaceRequest() *ControllersCreateNamespaceRequest {
	this := ControllersCreateNamespaceRequest{}
	return &this
}

// NewControllersCreateNamespaceRequestWithDefaults instantiates a new ControllersCreateNamespaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersCreateNamespaceRequestWithDefaults() *ControllersCreateNamespaceRequest {
	this := ControllersCreateNamespaceRequest{}
	return &this
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *ControllersCreateNamespaceRequest) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateNamespaceRequest) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *ControllersCreateNamespaceRequest) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *ControllersCreateNamespaceRequest) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllersCreateNamespaceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateNamespaceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllersCreateNamespaceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllersCreateNamespaceRequest) SetName(v string) {
	o.Name = &v
}

func (o ControllersCreateNamespaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeploymentTargetId != nil {
		toSerialize["deployment_target_id"] = o.DeploymentTargetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableControllersCreateNamespaceRequest struct {
	value *ControllersCreateNamespaceRequest
	isSet bool
}

func (v NullableControllersCreateNamespaceRequest) Get() *ControllersCreateNamespaceRequest {
	return v.value
}

func (v *NullableControllersCreateNamespaceRequest) Set(val *ControllersCreateNamespaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersCreateNamespaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersCreateNamespaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersCreateNamespaceRequest(val *ControllersCreateNamespaceRequest) *NullableControllersCreateNamespaceRequest {
	return &NullableControllersCreateNamespaceRequest{value: val, isSet: true}
}

func (v NullableControllersCreateNamespaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersCreateNamespaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


