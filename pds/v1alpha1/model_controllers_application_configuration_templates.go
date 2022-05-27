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

// ControllersApplicationConfigurationTemplates struct for ControllersApplicationConfigurationTemplates
type ControllersApplicationConfigurationTemplates struct {
	Data []ModelsApplicationConfigurationTemplate `json:"data,omitempty"`
}

// NewControllersApplicationConfigurationTemplates instantiates a new ControllersApplicationConfigurationTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersApplicationConfigurationTemplates() *ControllersApplicationConfigurationTemplates {
	this := ControllersApplicationConfigurationTemplates{}
	return &this
}

// NewControllersApplicationConfigurationTemplatesWithDefaults instantiates a new ControllersApplicationConfigurationTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersApplicationConfigurationTemplatesWithDefaults() *ControllersApplicationConfigurationTemplates {
	this := ControllersApplicationConfigurationTemplates{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControllersApplicationConfigurationTemplates) GetData() []ModelsApplicationConfigurationTemplate {
	if o == nil || o.Data == nil {
		var ret []ModelsApplicationConfigurationTemplate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersApplicationConfigurationTemplates) GetDataOk() ([]ModelsApplicationConfigurationTemplate, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControllersApplicationConfigurationTemplates) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsApplicationConfigurationTemplate and assigns it to the Data field.
func (o *ControllersApplicationConfigurationTemplates) SetData(v []ModelsApplicationConfigurationTemplate) {
	o.Data = v
}

func (o ControllersApplicationConfigurationTemplates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableControllersApplicationConfigurationTemplates struct {
	value *ControllersApplicationConfigurationTemplates
	isSet bool
}

func (v NullableControllersApplicationConfigurationTemplates) Get() *ControllersApplicationConfigurationTemplates {
	return v.value
}

func (v *NullableControllersApplicationConfigurationTemplates) Set(val *ControllersApplicationConfigurationTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersApplicationConfigurationTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersApplicationConfigurationTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersApplicationConfigurationTemplates(val *ControllersApplicationConfigurationTemplates) *NullableControllersApplicationConfigurationTemplates {
	return &NullableControllersApplicationConfigurationTemplates{value: val, isSet: true}
}

func (v NullableControllersApplicationConfigurationTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersApplicationConfigurationTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


