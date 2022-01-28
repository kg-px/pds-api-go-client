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

// ControllersWhoAmIUser struct for ControllersWhoAmIUser
type ControllersWhoAmIUser struct {
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	DeletedAt *GormDeletedAt `json:"deleted_at,omitempty"`
	Email *string `json:"email,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewControllersWhoAmIUser instantiates a new ControllersWhoAmIUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersWhoAmIUser() *ControllersWhoAmIUser {
	this := ControllersWhoAmIUser{}
	return &this
}

// NewControllersWhoAmIUserWithDefaults instantiates a new ControllersWhoAmIUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersWhoAmIUserWithDefaults() *ControllersWhoAmIUser {
	this := ControllersWhoAmIUser{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ControllersWhoAmIUser) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIUser) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ControllersWhoAmIUser) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ControllersWhoAmIUser) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ControllersWhoAmIUser) GetDeletedAt() GormDeletedAt {
	if o == nil || o.DeletedAt == nil {
		var ret GormDeletedAt
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIUser) GetDeletedAtOk() (*GormDeletedAt, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ControllersWhoAmIUser) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given GormDeletedAt and assigns it to the DeletedAt field.
func (o *ControllersWhoAmIUser) SetDeletedAt(v GormDeletedAt) {
	o.DeletedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ControllersWhoAmIUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ControllersWhoAmIUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ControllersWhoAmIUser) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ControllersWhoAmIUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ControllersWhoAmIUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ControllersWhoAmIUser) SetId(v string) {
	o.Id = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ControllersWhoAmIUser) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersWhoAmIUser) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ControllersWhoAmIUser) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ControllersWhoAmIUser) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ControllersWhoAmIUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableControllersWhoAmIUser struct {
	value *ControllersWhoAmIUser
	isSet bool
}

func (v NullableControllersWhoAmIUser) Get() *ControllersWhoAmIUser {
	return v.value
}

func (v *NullableControllersWhoAmIUser) Set(val *ControllersWhoAmIUser) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersWhoAmIUser) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersWhoAmIUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersWhoAmIUser(val *ControllersWhoAmIUser) *NullableControllersWhoAmIUser {
	return &NullableControllersWhoAmIUser{value: val, isSet: true}
}

func (v NullableControllersWhoAmIUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersWhoAmIUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


