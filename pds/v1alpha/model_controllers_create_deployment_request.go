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

// ControllersCreateDeploymentRequest struct for ControllersCreateDeploymentRequest
type ControllersCreateDeploymentRequest struct {
	ClusterName *string `json:"cluster_name,omitempty"`
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	DataServiceId *string `json:"data_service_id,omitempty"`
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	DnsZone *string `json:"dns_zone,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	LoadBalancerSourceRanges *[]string `json:"load_balancer_source_ranges,omitempty"`
	NamespaceId *string `json:"namespace_id,omitempty"`
	NodeCount *int32 `json:"node_count,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	Resources *map[string]interface{} `json:"resources,omitempty"`
	Service *string `json:"service,omitempty"`
	ServiceType *string `json:"service_type,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
}

// NewControllersCreateDeploymentRequest instantiates a new ControllersCreateDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersCreateDeploymentRequest() *ControllersCreateDeploymentRequest {
	this := ControllersCreateDeploymentRequest{}
	return &this
}

// NewControllersCreateDeploymentRequestWithDefaults instantiates a new ControllersCreateDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersCreateDeploymentRequestWithDefaults() *ControllersCreateDeploymentRequest {
	this := ControllersCreateDeploymentRequest{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ControllersCreateDeploymentRequest) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetConfiguration() map[string]interface{} {
	if o == nil || o.Configuration == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetConfigurationOk() (*map[string]interface{}, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *ControllersCreateDeploymentRequest) SetConfiguration(v map[string]interface{}) {
	o.Configuration = &v
}

// GetDataServiceId returns the DataServiceId field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetDataServiceId() string {
	if o == nil || o.DataServiceId == nil {
		var ret string
		return ret
	}
	return *o.DataServiceId
}

// GetDataServiceIdOk returns a tuple with the DataServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetDataServiceIdOk() (*string, bool) {
	if o == nil || o.DataServiceId == nil {
		return nil, false
	}
	return o.DataServiceId, true
}

// HasDataServiceId returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasDataServiceId() bool {
	if o != nil && o.DataServiceId != nil {
		return true
	}

	return false
}

// SetDataServiceId gets a reference to the given string and assigns it to the DataServiceId field.
func (o *ControllersCreateDeploymentRequest) SetDataServiceId(v string) {
	o.DataServiceId = &v
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *ControllersCreateDeploymentRequest) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetDnsZone returns the DnsZone field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetDnsZone() string {
	if o == nil || o.DnsZone == nil {
		var ret string
		return ret
	}
	return *o.DnsZone
}

// GetDnsZoneOk returns a tuple with the DnsZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetDnsZoneOk() (*string, bool) {
	if o == nil || o.DnsZone == nil {
		return nil, false
	}
	return o.DnsZone, true
}

// HasDnsZone returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasDnsZone() bool {
	if o != nil && o.DnsZone != nil {
		return true
	}

	return false
}

// SetDnsZone gets a reference to the given string and assigns it to the DnsZone field.
func (o *ControllersCreateDeploymentRequest) SetDnsZone(v string) {
	o.DnsZone = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ControllersCreateDeploymentRequest) SetImageId(v string) {
	o.ImageId = &v
}

// GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetLoadBalancerSourceRanges() []string {
	if o == nil || o.LoadBalancerSourceRanges == nil {
		var ret []string
		return ret
	}
	return *o.LoadBalancerSourceRanges
}

// GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetLoadBalancerSourceRangesOk() (*[]string, bool) {
	if o == nil || o.LoadBalancerSourceRanges == nil {
		return nil, false
	}
	return o.LoadBalancerSourceRanges, true
}

// HasLoadBalancerSourceRanges returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasLoadBalancerSourceRanges() bool {
	if o != nil && o.LoadBalancerSourceRanges != nil {
		return true
	}

	return false
}

// SetLoadBalancerSourceRanges gets a reference to the given []string and assigns it to the LoadBalancerSourceRanges field.
func (o *ControllersCreateDeploymentRequest) SetLoadBalancerSourceRanges(v []string) {
	o.LoadBalancerSourceRanges = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetNamespaceId() string {
	if o == nil || o.NamespaceId == nil {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetNamespaceIdOk() (*string, bool) {
	if o == nil || o.NamespaceId == nil {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasNamespaceId() bool {
	if o != nil && o.NamespaceId != nil {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *ControllersCreateDeploymentRequest) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetNodeCount() int32 {
	if o == nil || o.NodeCount == nil {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetNodeCountOk() (*int32, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *ControllersCreateDeploymentRequest) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ControllersCreateDeploymentRequest) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetResources() map[string]interface{} {
	if o == nil || o.Resources == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetResourcesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given map[string]interface{} and assigns it to the Resources field.
func (o *ControllersCreateDeploymentRequest) SetResources(v map[string]interface{}) {
	o.Resources = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ControllersCreateDeploymentRequest) SetService(v string) {
	o.Service = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *ControllersCreateDeploymentRequest) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *ControllersCreateDeploymentRequest) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersCreateDeploymentRequest) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *ControllersCreateDeploymentRequest) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *ControllersCreateDeploymentRequest) SetVersionId(v string) {
	o.VersionId = &v
}

func (o ControllersCreateDeploymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterName != nil {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DataServiceId != nil {
		toSerialize["data_service_id"] = o.DataServiceId
	}
	if o.DeploymentTargetId != nil {
		toSerialize["deployment_target_id"] = o.DeploymentTargetId
	}
	if o.DnsZone != nil {
		toSerialize["dns_zone"] = o.DnsZone
	}
	if o.ImageId != nil {
		toSerialize["image_id"] = o.ImageId
	}
	if o.LoadBalancerSourceRanges != nil {
		toSerialize["load_balancer_source_ranges"] = o.LoadBalancerSourceRanges
	}
	if o.NamespaceId != nil {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if o.NodeCount != nil {
		toSerialize["node_count"] = o.NodeCount
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ServiceType != nil {
		toSerialize["service_type"] = o.ServiceType
	}
	if o.VersionId != nil {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableControllersCreateDeploymentRequest struct {
	value *ControllersCreateDeploymentRequest
	isSet bool
}

func (v NullableControllersCreateDeploymentRequest) Get() *ControllersCreateDeploymentRequest {
	return v.value
}

func (v *NullableControllersCreateDeploymentRequest) Set(val *ControllersCreateDeploymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersCreateDeploymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersCreateDeploymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersCreateDeploymentRequest(val *ControllersCreateDeploymentRequest) *NullableControllersCreateDeploymentRequest {
	return &NullableControllersCreateDeploymentRequest{value: val, isSet: true}
}

func (v NullableControllersCreateDeploymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersCreateDeploymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


