# ControllersUpdateDeploymentScheduledBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPolicyId** | Pointer to **string** |  | [optional] 
**BackupTargetId** | Pointer to **string** | BackupTargetID and BackupPolicyID must be both specified or not. Set both to null to have no scheduled backup. | [optional] 

## Methods

### NewControllersUpdateDeploymentScheduledBackup

`func NewControllersUpdateDeploymentScheduledBackup() *ControllersUpdateDeploymentScheduledBackup`

NewControllersUpdateDeploymentScheduledBackup instantiates a new ControllersUpdateDeploymentScheduledBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersUpdateDeploymentScheduledBackupWithDefaults

`func NewControllersUpdateDeploymentScheduledBackupWithDefaults() *ControllersUpdateDeploymentScheduledBackup`

NewControllersUpdateDeploymentScheduledBackupWithDefaults instantiates a new ControllersUpdateDeploymentScheduledBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicyId

`func (o *ControllersUpdateDeploymentScheduledBackup) GetBackupPolicyId() string`

GetBackupPolicyId returns the BackupPolicyId field if non-nil, zero value otherwise.

### GetBackupPolicyIdOk

`func (o *ControllersUpdateDeploymentScheduledBackup) GetBackupPolicyIdOk() (*string, bool)`

GetBackupPolicyIdOk returns a tuple with the BackupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicyId

`func (o *ControllersUpdateDeploymentScheduledBackup) SetBackupPolicyId(v string)`

SetBackupPolicyId sets BackupPolicyId field to given value.

### HasBackupPolicyId

`func (o *ControllersUpdateDeploymentScheduledBackup) HasBackupPolicyId() bool`

HasBackupPolicyId returns a boolean if a field has been set.

### GetBackupTargetId

`func (o *ControllersUpdateDeploymentScheduledBackup) GetBackupTargetId() string`

GetBackupTargetId returns the BackupTargetId field if non-nil, zero value otherwise.

### GetBackupTargetIdOk

`func (o *ControllersUpdateDeploymentScheduledBackup) GetBackupTargetIdOk() (*string, bool)`

GetBackupTargetIdOk returns a tuple with the BackupTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTargetId

`func (o *ControllersUpdateDeploymentScheduledBackup) SetBackupTargetId(v string)`

SetBackupTargetId sets BackupTargetId field to given value.

### HasBackupTargetId

`func (o *ControllersUpdateDeploymentScheduledBackup) HasBackupTargetId() bool`

HasBackupTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


