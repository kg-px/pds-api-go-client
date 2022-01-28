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

// ModelsConfiguration struct for ModelsConfiguration
type ModelsConfiguration struct {
	DefaultValue *map[string]interface{} `json:"default_value,omitempty"`
	Name *string `json:"name,omitempty"`
	Required *bool `json:"required,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewModelsConfiguration instantiates a new ModelsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsConfiguration() *ModelsConfiguration {
	this := ModelsConfiguration{}
	return &this
}

// NewModelsConfigurationWithDefaults instantiates a new ModelsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsConfigurationWithDefaults() *ModelsConfiguration {
	this := ModelsConfiguration{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *ModelsConfiguration) GetDefaultValue() map[string]interface{} {
	if o == nil || o.DefaultValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsConfiguration) GetDefaultValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *ModelsConfiguration) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given map[string]interface{} and assigns it to the DefaultValue field.
func (o *ModelsConfiguration) SetDefaultValue(v map[string]interface{}) {
	o.DefaultValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsConfiguration) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsConfiguration) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsConfiguration) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsConfiguration) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ModelsConfiguration) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsConfiguration) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ModelsConfiguration) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ModelsConfiguration) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelsConfiguration) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsConfiguration) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelsConfiguration) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelsConfiguration) SetType(v string) {
	o.Type = &v
}

func (o ModelsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultValue != nil {
		toSerialize["default_value"] = o.DefaultValue
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableModelsConfiguration struct {
	value *ModelsConfiguration
	isSet bool
}

func (v NullableModelsConfiguration) Get() *ModelsConfiguration {
	return v.value
}

func (v *NullableModelsConfiguration) Set(val *ModelsConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsConfiguration(val *ModelsConfiguration) *NullableModelsConfiguration {
	return &NullableModelsConfiguration{value: val, isSet: true}
}

func (v NullableModelsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


