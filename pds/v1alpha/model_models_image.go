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

// ModelsImage struct for ModelsImage
type ModelsImage struct {
	Build *string `json:"build,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	DataServiceId *string `json:"data_service_id,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Registry *string `json:"registry,omitempty"`
	Tag *string `json:"tag,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
}

// NewModelsImage instantiates a new ModelsImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsImage() *ModelsImage {
	this := ModelsImage{}
	return &this
}

// NewModelsImageWithDefaults instantiates a new ModelsImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsImageWithDefaults() *ModelsImage {
	this := ModelsImage{}
	return &this
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *ModelsImage) GetBuild() string {
	if o == nil || o.Build == nil {
		var ret string
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetBuildOk() (*string, bool) {
	if o == nil || o.Build == nil {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *ModelsImage) HasBuild() bool {
	if o != nil && o.Build != nil {
		return true
	}

	return false
}

// SetBuild gets a reference to the given string and assigns it to the Build field.
func (o *ModelsImage) SetBuild(v string) {
	o.Build = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsImage) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsImage) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsImage) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDataServiceId returns the DataServiceId field value if set, zero value otherwise.
func (o *ModelsImage) GetDataServiceId() string {
	if o == nil || o.DataServiceId == nil {
		var ret string
		return ret
	}
	return *o.DataServiceId
}

// GetDataServiceIdOk returns a tuple with the DataServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetDataServiceIdOk() (*string, bool) {
	if o == nil || o.DataServiceId == nil {
		return nil, false
	}
	return o.DataServiceId, true
}

// HasDataServiceId returns a boolean if a field has been set.
func (o *ModelsImage) HasDataServiceId() bool {
	if o != nil && o.DataServiceId != nil {
		return true
	}

	return false
}

// SetDataServiceId gets a reference to the given string and assigns it to the DataServiceId field.
func (o *ModelsImage) SetDataServiceId(v string) {
	o.DataServiceId = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ModelsImage) GetDeletedAt() GormDeletedAt {
	if o == nil || o.DeletedAt == nil {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelsImage) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ModelsImage) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsImage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsImage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsImage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsImage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsImage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsImage) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ModelsImage) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ModelsImage) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ModelsImage) SetNamespace(v string) {
	o.Namespace = &v
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ModelsImage) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ModelsImage) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ModelsImage) SetRegistry(v string) {
	o.Registry = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ModelsImage) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ModelsImage) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ModelsImage) SetTag(v string) {
	o.Tag = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsImage) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsImage) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsImage) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *ModelsImage) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsImage) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *ModelsImage) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *ModelsImage) SetVersionId(v string) {
	o.VersionId = &v
}

func (o ModelsImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Build != nil {
		toSerialize["build"] = o.Build
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DataServiceId != nil {
		toSerialize["data_service_id"] = o.DataServiceId
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.VersionId != nil {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableModelsImage struct {
	value *ModelsImage
	isSet bool
}

func (v NullableModelsImage) Get() *ModelsImage {
	return v.value
}

func (v *NullableModelsImage) Set(val *ModelsImage) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsImage) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsImage(val *ModelsImage) *NullableModelsImage {
	return &NullableModelsImage{value: val, isSet: true}
}

func (v NullableModelsImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


