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

// ModelsBackupTargetState struct for ModelsBackupTargetState
type ModelsBackupTargetState struct {
	BackupTargetId *string `json:"backup_target_id,omitempty"`
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	// Predefined API error code.
	ErrorCode *string `json:"error_code,omitempty"`
	// More detailed error message possibly containing the root cause of the error, not suitable to show in the UI.
	ErrorDetails *string `json:"error_details,omitempty"`
	// High level human-readable error message determined by the ErrorCode.
	ErrorMessage *string `json:"error_message,omitempty"`
	// ID of the credentials in PX cluster.
	PxCredentialsId *string `json:"px_credentials_id,omitempty"`
	// Name of the credentials in PX cluster. This will be used when creating a new backup.
	PxCredentialsName *string `json:"px_credentials_name,omitempty"`
	// State of the synchronization of credentials.
	State *string `json:"state,omitempty"`
}

// NewModelsBackupTargetState instantiates a new ModelsBackupTargetState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsBackupTargetState() *ModelsBackupTargetState {
	this := ModelsBackupTargetState{}
	return &this
}

// NewModelsBackupTargetStateWithDefaults instantiates a new ModelsBackupTargetState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsBackupTargetStateWithDefaults() *ModelsBackupTargetState {
	this := ModelsBackupTargetState{}
	return &this
}

// GetBackupTargetId returns the BackupTargetId field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetBackupTargetId() string {
	if o == nil || o.BackupTargetId == nil {
		var ret string
		return ret
	}
	return *o.BackupTargetId
}

// GetBackupTargetIdOk returns a tuple with the BackupTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetBackupTargetIdOk() (*string, bool) {
	if o == nil || o.BackupTargetId == nil {
		return nil, false
	}
	return o.BackupTargetId, true
}

// HasBackupTargetId returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasBackupTargetId() bool {
	if o != nil && o.BackupTargetId != nil {
		return true
	}

	return false
}

// SetBackupTargetId gets a reference to the given string and assigns it to the BackupTargetId field.
func (o *ModelsBackupTargetState) SetBackupTargetId(v string) {
	o.BackupTargetId = &v
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *ModelsBackupTargetState) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ModelsBackupTargetState) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetErrorDetails() string {
	if o == nil || o.ErrorDetails == nil {
		var ret string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetErrorDetailsOk() (*string, bool) {
	if o == nil || o.ErrorDetails == nil {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasErrorDetails() bool {
	if o != nil && o.ErrorDetails != nil {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given string and assigns it to the ErrorDetails field.
func (o *ModelsBackupTargetState) SetErrorDetails(v string) {
	o.ErrorDetails = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ModelsBackupTargetState) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetPxCredentialsId returns the PxCredentialsId field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetPxCredentialsId() string {
	if o == nil || o.PxCredentialsId == nil {
		var ret string
		return ret
	}
	return *o.PxCredentialsId
}

// GetPxCredentialsIdOk returns a tuple with the PxCredentialsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetPxCredentialsIdOk() (*string, bool) {
	if o == nil || o.PxCredentialsId == nil {
		return nil, false
	}
	return o.PxCredentialsId, true
}

// HasPxCredentialsId returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasPxCredentialsId() bool {
	if o != nil && o.PxCredentialsId != nil {
		return true
	}

	return false
}

// SetPxCredentialsId gets a reference to the given string and assigns it to the PxCredentialsId field.
func (o *ModelsBackupTargetState) SetPxCredentialsId(v string) {
	o.PxCredentialsId = &v
}

// GetPxCredentialsName returns the PxCredentialsName field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetPxCredentialsName() string {
	if o == nil || o.PxCredentialsName == nil {
		var ret string
		return ret
	}
	return *o.PxCredentialsName
}

// GetPxCredentialsNameOk returns a tuple with the PxCredentialsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetPxCredentialsNameOk() (*string, bool) {
	if o == nil || o.PxCredentialsName == nil {
		return nil, false
	}
	return o.PxCredentialsName, true
}

// HasPxCredentialsName returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasPxCredentialsName() bool {
	if o != nil && o.PxCredentialsName != nil {
		return true
	}

	return false
}

// SetPxCredentialsName gets a reference to the given string and assigns it to the PxCredentialsName field.
func (o *ModelsBackupTargetState) SetPxCredentialsName(v string) {
	o.PxCredentialsName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ModelsBackupTargetState) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupTargetState) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ModelsBackupTargetState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ModelsBackupTargetState) SetState(v string) {
	o.State = &v
}

func (o ModelsBackupTargetState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupTargetId != nil {
		toSerialize["backup_target_id"] = o.BackupTargetId
	}
	if o.DeploymentTargetId != nil {
		toSerialize["deployment_target_id"] = o.DeploymentTargetId
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorDetails != nil {
		toSerialize["error_details"] = o.ErrorDetails
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.PxCredentialsId != nil {
		toSerialize["px_credentials_id"] = o.PxCredentialsId
	}
	if o.PxCredentialsName != nil {
		toSerialize["px_credentials_name"] = o.PxCredentialsName
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableModelsBackupTargetState struct {
	value *ModelsBackupTargetState
	isSet bool
}

func (v NullableModelsBackupTargetState) Get() *ModelsBackupTargetState {
	return v.value
}

func (v *NullableModelsBackupTargetState) Set(val *ModelsBackupTargetState) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsBackupTargetState) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsBackupTargetState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsBackupTargetState(val *ModelsBackupTargetState) *NullableModelsBackupTargetState {
	return &NullableModelsBackupTargetState{value: val, isSet: true}
}

func (v NullableModelsBackupTargetState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsBackupTargetState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


