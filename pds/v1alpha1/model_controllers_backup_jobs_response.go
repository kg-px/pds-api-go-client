/*
PDS API

Portworx Data Services API Server

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ControllersBackupJobsResponse struct for ControllersBackupJobsResponse
type ControllersBackupJobsResponse struct {
	Data *[]ControllersBackupJobStatus `json:"data,omitempty"`
}

// NewControllersBackupJobsResponse instantiates a new ControllersBackupJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersBackupJobsResponse() *ControllersBackupJobsResponse {
	this := ControllersBackupJobsResponse{}
	return &this
}

// NewControllersBackupJobsResponseWithDefaults instantiates a new ControllersBackupJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersBackupJobsResponseWithDefaults() *ControllersBackupJobsResponse {
	this := ControllersBackupJobsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersBackupJobsResponse) GetData() []ControllersBackupJobStatus {
	if o == nil || o.Data == nil {
		var ret []ControllersBackupJobStatus
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersBackupJobsResponse) GetDataOk() (*[]ControllersBackupJobStatus, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersBackupJobsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ControllersBackupJobStatus and assigns it to the Data field.
func (o *ControllersBackupJobsResponse) SetData(v []ControllersBackupJobStatus) {
	o.Data = &v
}

func (o ControllersBackupJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableControllersBackupJobsResponse struct {
	value *ControllersBackupJobsResponse
	isSet bool
}

func (v NullableControllersBackupJobsResponse) Get() *ControllersBackupJobsResponse {
	return v.value
}

func (v *NullableControllersBackupJobsResponse) Set(val *ControllersBackupJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersBackupJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersBackupJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersBackupJobsResponse(val *ControllersBackupJobsResponse) *NullableControllersBackupJobsResponse {
	return &NullableControllersBackupJobsResponse{value: val, isSet: true}
}

func (v NullableControllersBackupJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersBackupJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


