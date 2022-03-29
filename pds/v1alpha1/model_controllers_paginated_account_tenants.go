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

// ControllersPaginatedAccountTenants struct for ControllersPaginatedAccountTenants
type ControllersPaginatedAccountTenants struct {
	Data *[]ModelsTenant `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewControllersPaginatedAccountTenants instantiates a new ControllersPaginatedAccountTenants object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPaginatedAccountTenants() *ControllersPaginatedAccountTenants {
	this := ControllersPaginatedAccountTenants{}
	return &this
}

// NewControllersPaginatedAccountTenantsWithDefaults instantiates a new ControllersPaginatedAccountTenants object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPaginatedAccountTenantsWithDefaults() *ControllersPaginatedAccountTenants {
	this := ControllersPaginatedAccountTenants{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersPaginatedAccountTenants) GetData() []ModelsTenant {
	if o == nil || o.Data == nil {
		var ret []ModelsTenant
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedAccountTenants) GetDataOk() (*[]ModelsTenant, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersPaginatedAccountTenants) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsTenant and assigns it to the Data field.
func (o *ControllersPaginatedAccountTenants) SetData(v []ModelsTenant) {
	o.Data = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ControllersPaginatedAccountTenants) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedAccountTenants) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ControllersPaginatedAccountTenants) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ControllersPaginatedAccountTenants) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ControllersPaginatedAccountTenants) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableControllersPaginatedAccountTenants struct {
	value *ControllersPaginatedAccountTenants
	isSet bool
}

func (v NullableControllersPaginatedAccountTenants) Get() *ControllersPaginatedAccountTenants {
	return v.value
}

func (v *NullableControllersPaginatedAccountTenants) Set(val *ControllersPaginatedAccountTenants) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPaginatedAccountTenants) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPaginatedAccountTenants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPaginatedAccountTenants(val *ControllersPaginatedAccountTenants) *NullableControllersPaginatedAccountTenants {
	return &NullableControllersPaginatedAccountTenants{value: val, isSet: true}
}

func (v NullableControllersPaginatedAccountTenants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPaginatedAccountTenants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


