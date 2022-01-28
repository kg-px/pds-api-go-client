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

// ControllersPaginatedServiceAccounts struct for ControllersPaginatedServiceAccounts
type ControllersPaginatedServiceAccounts struct {
	Data *[]ModelsServiceAccount `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewControllersPaginatedServiceAccounts instantiates a new ControllersPaginatedServiceAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPaginatedServiceAccounts() *ControllersPaginatedServiceAccounts {
	this := ControllersPaginatedServiceAccounts{}
	return &this
}

// NewControllersPaginatedServiceAccountsWithDefaults instantiates a new ControllersPaginatedServiceAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPaginatedServiceAccountsWithDefaults() *ControllersPaginatedServiceAccounts {
	this := ControllersPaginatedServiceAccounts{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersPaginatedServiceAccounts) GetData() []ModelsServiceAccount {
	if o == nil || o.Data == nil {
		var ret []ModelsServiceAccount
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedServiceAccounts) GetDataOk() (*[]ModelsServiceAccount, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersPaginatedServiceAccounts) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsServiceAccount and assigns it to the Data field.
func (o *ControllersPaginatedServiceAccounts) SetData(v []ModelsServiceAccount) {
	o.Data = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ControllersPaginatedServiceAccounts) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPaginatedServiceAccounts) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ControllersPaginatedServiceAccounts) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ControllersPaginatedServiceAccounts) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ControllersPaginatedServiceAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableControllersPaginatedServiceAccounts struct {
	value *ControllersPaginatedServiceAccounts
	isSet bool
}

func (v NullableControllersPaginatedServiceAccounts) Get() *ControllersPaginatedServiceAccounts {
	return v.value
}

func (v *NullableControllersPaginatedServiceAccounts) Set(val *ControllersPaginatedServiceAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPaginatedServiceAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPaginatedServiceAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPaginatedServiceAccounts(val *ControllersPaginatedServiceAccounts) *NullableControllersPaginatedServiceAccounts {
	return &NullableControllersPaginatedServiceAccounts{value: val, isSet: true}
}

func (v NullableControllersPaginatedServiceAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPaginatedServiceAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


