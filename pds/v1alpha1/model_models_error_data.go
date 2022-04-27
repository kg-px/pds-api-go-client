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

// ModelsErrorData struct for ModelsErrorData
type ModelsErrorData struct {
	// Predefined API error code.
	Code *string `json:"code,omitempty"`
	// More detailed error message possibly containing the root cause.
	Details *string `json:"details,omitempty"`
	// High level human-readable error message determined by the ErrorCode.
	Message *string `json:"message,omitempty"`
}

// NewModelsErrorData instantiates a new ModelsErrorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsErrorData() *ModelsErrorData {
	this := ModelsErrorData{}
	return &this
}

// NewModelsErrorDataWithDefaults instantiates a new ModelsErrorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsErrorDataWithDefaults() *ModelsErrorData {
	this := ModelsErrorData{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ModelsErrorData) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsErrorData) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ModelsErrorData) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ModelsErrorData) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ModelsErrorData) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsErrorData) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ModelsErrorData) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ModelsErrorData) SetDetails(v string) {
	o.Details = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ModelsErrorData) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsErrorData) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ModelsErrorData) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ModelsErrorData) SetMessage(v string) {
	o.Message = &v
}

func (o ModelsErrorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableModelsErrorData struct {
	value *ModelsErrorData
	isSet bool
}

func (v NullableModelsErrorData) Get() *ModelsErrorData {
	return v.value
}

func (v *NullableModelsErrorData) Set(val *ModelsErrorData) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsErrorData) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsErrorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsErrorData(val *ModelsErrorData) *NullableModelsErrorData {
	return &NullableModelsErrorData{value: val, isSet: true}
}

func (v NullableModelsErrorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsErrorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


