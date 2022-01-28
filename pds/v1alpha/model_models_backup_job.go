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

// ModelsBackupJob struct for ModelsBackupJob
type ModelsBackupJob struct {
	BackupId *string `json:"backup_id,omitempty"`
	CompletionTime *string `json:"completion_time,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	// Endpoint is populated by the Tekton pipeline on creation
	Endpoint *string `json:"endpoint,omitempty"`
	// FileSize is managed by the API
	FileSize *int32 `json:"file_size,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
	// State is managed by the API
	State *string `json:"state,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsBackupJob instantiates a new ModelsBackupJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsBackupJob() *ModelsBackupJob {
	this := ModelsBackupJob{}
	return &this
}

// NewModelsBackupJobWithDefaults instantiates a new ModelsBackupJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsBackupJobWithDefaults() *ModelsBackupJob {
	this := ModelsBackupJob{}
	return &this
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetBackupId() string {
	if o == nil || o.BackupId == nil {
		var ret string
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetBackupIdOk() (*string, bool) {
	if o == nil || o.BackupId == nil {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasBackupId() bool {
	if o != nil && o.BackupId != nil {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given string and assigns it to the BackupId field.
func (o *ModelsBackupJob) SetBackupId(v string) {
	o.BackupId = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCompletionTime() string {
	if o == nil || o.CompletionTime == nil {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCompletionTimeOk() (*string, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *ModelsBackupJob) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsBackupJob) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetDeletedAt() GormDeletedAt {
	if o == nil || o.DeletedAt == nil {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ModelsBackupJob) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ModelsBackupJob) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *ModelsBackupJob) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsBackupJob) SetId(v string) {
	o.Id = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ModelsBackupJob) SetStartTime(v string) {
	o.StartTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ModelsBackupJob) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsBackupJob) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsBackupJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupId != nil {
		toSerialize["backup_id"] = o.BackupId
	}
	if o.CompletionTime != nil {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.FileSize != nil {
		toSerialize["file_size"] = o.FileSize
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelsBackupJob struct {
	value *ModelsBackupJob
	isSet bool
}

func (v NullableModelsBackupJob) Get() *ModelsBackupJob {
	return v.value
}

func (v *NullableModelsBackupJob) Set(val *ModelsBackupJob) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsBackupJob) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsBackupJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsBackupJob(val *ModelsBackupJob) *NullableModelsBackupJob {
	return &NullableModelsBackupJob{value: val, isSet: true}
}

func (v NullableModelsBackupJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsBackupJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


