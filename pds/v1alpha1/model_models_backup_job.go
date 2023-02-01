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

// ModelsBackupJob struct for ModelsBackupJob
type ModelsBackupJob struct {
	Backup *ModelsBackup `json:"backup,omitempty"`
	// BackupID which backup created the snapshot(nullable).
	BackupId *string `json:"backup_id,omitempty"`
	// CloudCredentialName credentials to access snapshot.Model
	CloudCredentialName *string `json:"cloud_credential_name,omitempty"`
	// CloudSnapID snapshot of the backup volume.
	CloudSnapId *string `json:"cloud_snap_id,omitempty"`
	// CompletionStatus of the backup job.
	CompletionStatus *string `json:"completion_status,omitempty"`
	CompletionTime *string `json:"completion_time,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	Deployment *ModelsDeployment `json:"deployment,omitempty"`
	// DeploymentID which deployment was backed up (nullable).
	DeploymentId *string `json:"deployment_id,omitempty"`
	ErrorCode *string `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	// FileSize of the CloudSnap image.
	FileSize *int32 `json:"file_size,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	// Name shown in UI.
	Name *string `json:"name,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
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

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetBackup() ModelsBackup {
	if o == nil || o.Backup == nil {
		var ret ModelsBackup
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetBackupOk() (*ModelsBackup, bool) {
	if o == nil || o.Backup == nil {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasBackup() bool {
	if o != nil && o.Backup != nil {
		return true
	}

	return false
}

// SetBackup gets a reference to the given ModelsBackup and assigns it to the Backup field.
func (o *ModelsBackupJob) SetBackup(v ModelsBackup) {
	o.Backup = &v
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

// GetCloudCredentialName returns the CloudCredentialName field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCloudCredentialName() string {
	if o == nil || o.CloudCredentialName == nil {
		var ret string
		return ret
	}
	return *o.CloudCredentialName
}

// GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCloudCredentialNameOk() (*string, bool) {
	if o == nil || o.CloudCredentialName == nil {
		return nil, false
	}
	return o.CloudCredentialName, true
}

// HasCloudCredentialName returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCloudCredentialName() bool {
	if o != nil && o.CloudCredentialName != nil {
		return true
	}

	return false
}

// SetCloudCredentialName gets a reference to the given string and assigns it to the CloudCredentialName field.
func (o *ModelsBackupJob) SetCloudCredentialName(v string) {
	o.CloudCredentialName = &v
}

// GetCloudSnapId returns the CloudSnapId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCloudSnapId() string {
	if o == nil || o.CloudSnapId == nil {
		var ret string
		return ret
	}
	return *o.CloudSnapId
}

// GetCloudSnapIdOk returns a tuple with the CloudSnapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCloudSnapIdOk() (*string, bool) {
	if o == nil || o.CloudSnapId == nil {
		return nil, false
	}
	return o.CloudSnapId, true
}

// HasCloudSnapId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCloudSnapId() bool {
	if o != nil && o.CloudSnapId != nil {
		return true
	}

	return false
}

// SetCloudSnapId gets a reference to the given string and assigns it to the CloudSnapId field.
func (o *ModelsBackupJob) SetCloudSnapId(v string) {
	o.CloudSnapId = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetCompletionStatus() string {
	if o == nil || o.CompletionStatus == nil {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetCompletionStatusOk() (*string, bool) {
	if o == nil || o.CompletionStatus == nil {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasCompletionStatus() bool {
	if o != nil && o.CompletionStatus != nil {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *ModelsBackupJob) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
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

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetDeployment() ModelsDeployment {
	if o == nil || o.Deployment == nil {
		var ret ModelsDeployment
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetDeploymentOk() (*ModelsDeployment, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasDeployment() bool {
	if o != nil && o.Deployment != nil {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given ModelsDeployment and assigns it to the Deployment field.
func (o *ModelsBackupJob) SetDeployment(v ModelsDeployment) {
	o.Deployment = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *ModelsBackupJob) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ModelsBackupJob) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ModelsBackupJob) SetErrorMessage(v string) {
	o.ErrorMessage = &v
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsBackupJob) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsBackupJob) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsBackupJob) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsBackupJob) SetName(v string) {
	o.Name = &v
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
	if o.Backup != nil {
		toSerialize["backup"] = o.Backup
	}
	if o.BackupId != nil {
		toSerialize["backup_id"] = o.BackupId
	}
	if o.CloudCredentialName != nil {
		toSerialize["cloud_credential_name"] = o.CloudCredentialName
	}
	if o.CloudSnapId != nil {
		toSerialize["cloud_snap_id"] = o.CloudSnapId
	}
	if o.CompletionStatus != nil {
		toSerialize["completion_status"] = o.CompletionStatus
	}
	if o.CompletionTime != nil {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Deployment != nil {
		toSerialize["deployment"] = o.Deployment
	}
	if o.DeploymentId != nil {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.FileSize != nil {
		toSerialize["file_size"] = o.FileSize
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
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


