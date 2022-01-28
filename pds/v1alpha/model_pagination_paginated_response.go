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

// PaginationPaginatedResponse struct for PaginationPaginatedResponse
type PaginationPaginatedResponse struct {
	Count *int32 `json:"count,omitempty"`
	Members *map[string]interface{} `json:"members,omitempty"`
	Page *int32 `json:"page,omitempty"`
	PerPage *int32 `json:"per_page,omitempty"`
	Total *int32 `json:"total,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
}

// NewPaginationPaginatedResponse instantiates a new PaginationPaginatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationPaginatedResponse() *PaginationPaginatedResponse {
	this := PaginationPaginatedResponse{}
	return &this
}

// NewPaginationPaginatedResponseWithDefaults instantiates a new PaginationPaginatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationPaginatedResponseWithDefaults() *PaginationPaginatedResponse {
	this := PaginationPaginatedResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginationPaginatedResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPaginatedResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginationPaginatedResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginationPaginatedResponse) SetCount(v int32) {
	o.Count = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *PaginationPaginatedResponse) GetMembers() map[string]interface{} {
	if o == nil || o.Members == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPaginatedResponse) GetMembersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *PaginationPaginatedResponse) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given map[string]interface{} and assigns it to the Members field.
func (o *PaginationPaginatedResponse) SetMembers(v map[string]interface{}) {
	o.Members = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginationPaginatedResponse) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPaginatedResponse) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginationPaginatedResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginationPaginatedResponse) SetPage(v int32) {
	o.Page = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *PaginationPaginatedResponse) GetPerPage() int32 {
	if o == nil || o.PerPage == nil {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPaginatedResponse) GetPerPageOk() (*int32, bool) {
	if o == nil || o.PerPage == nil {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *PaginationPaginatedResponse) HasPerPage() bool {
	if o != nil && o.PerPage != nil {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *PaginationPaginatedResponse) SetPerPage(v int32) {
	o.PerPage = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PaginationPaginatedResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPaginatedResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PaginationPaginatedResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PaginationPaginatedResponse) SetTotal(v int32) {
	o.Total = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *PaginationPaginatedResponse) GetTotalPages() int32 {
	if o == nil || o.TotalPages == nil {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationPaginatedResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || o.TotalPages == nil {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *PaginationPaginatedResponse) HasTotalPages() bool {
	if o != nil && o.TotalPages != nil {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *PaginationPaginatedResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o PaginationPaginatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PerPage != nil {
		toSerialize["per_page"] = o.PerPage
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.TotalPages != nil {
		toSerialize["total_pages"] = o.TotalPages
	}
	return json.Marshal(toSerialize)
}

type NullablePaginationPaginatedResponse struct {
	value *PaginationPaginatedResponse
	isSet bool
}

func (v NullablePaginationPaginatedResponse) Get() *PaginationPaginatedResponse {
	return v.value
}

func (v *NullablePaginationPaginatedResponse) Set(val *PaginationPaginatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationPaginatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationPaginatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationPaginatedResponse(val *PaginationPaginatedResponse) *NullablePaginationPaginatedResponse {
	return &NullablePaginationPaginatedResponse{value: val, isSet: true}
}

func (v NullablePaginationPaginatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationPaginatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


