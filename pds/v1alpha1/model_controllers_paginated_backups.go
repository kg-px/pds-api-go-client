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

// ControllersPaginatedBackups struct for ControllersPaginatedBackups
type ControllersPaginatedBackups struct {
	Data []ModelsBackup `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewControllersPaginatedBackups instantiates a new ControllersPaginatedBackups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPaginatedBackups() *ControllersPaginatedBackups {
	this := ControllersPaginatedBackups{}
	return &this
}

// NewControllersPaginatedBackupsWithDefaults instantiates a new ControllersPaginatedBackups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPaginatedBackupsWithDefaults() *ControllersPaginatedBackups {
	this := ControllersPaginatedBackups{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersPaginatedBackups) GetData() []ModelsBackup {
	if o == nil || o.Data == nil {
		var ret []ModelsBackup
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedBackups) GetDataOk() ([]ModelsBackup, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersPaginatedBackups) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsBackup and assigns it to the Data field.
func (o *ControllersPaginatedBackups) SetData(v []ModelsBackup) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ControllersPaginatedBackups) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedBackups) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ControllersPaginatedBackups) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ControllersPaginatedBackups) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ControllersPaginatedBackups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableControllersPaginatedBackups struct {
	value *ControllersPaginatedBackups
	isSet bool
}

func (v NullableControllersPaginatedBackups) Get() *ControllersPaginatedBackups {
	return v.value
}

func (v *NullableControllersPaginatedBackups) Set(val *ControllersPaginatedBackups) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPaginatedBackups) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPaginatedBackups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPaginatedBackups(val *ControllersPaginatedBackups) *NullableControllersPaginatedBackups {
	return &NullableControllersPaginatedBackups{value: val, isSet: true}
}

func (v NullableControllersPaginatedBackups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPaginatedBackups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


