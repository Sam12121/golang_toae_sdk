/*
Toae ThreatMapper

Toae Runtime API provides programmatic control over Toae microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: toaedev@toaesecurity.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the GraphNodeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphNodeInfo{}

// GraphNodeInfo struct for GraphNodeInfo
type GraphNodeInfo struct {
	CloudComplianceCount int32 `json:"cloud_compliance_count"`
	ComplianceCount int32 `json:"compliance_count"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	SecretsCount int32 `json:"secrets_count"`
	VulnerabilityCount int32 `json:"vulnerability_count"`
}

type _GraphNodeInfo GraphNodeInfo

// NewGraphNodeInfo instantiates a new GraphNodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphNodeInfo(cloudComplianceCount int32, complianceCount int32, name string, nodeId string, secretsCount int32, vulnerabilityCount int32) *GraphNodeInfo {
	this := GraphNodeInfo{}
	this.CloudComplianceCount = cloudComplianceCount
	this.ComplianceCount = complianceCount
	this.Name = name
	this.NodeId = nodeId
	this.SecretsCount = secretsCount
	this.VulnerabilityCount = vulnerabilityCount
	return &this
}

// NewGraphNodeInfoWithDefaults instantiates a new GraphNodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphNodeInfoWithDefaults() *GraphNodeInfo {
	this := GraphNodeInfo{}
	return &this
}

// GetCloudComplianceCount returns the CloudComplianceCount field value
func (o *GraphNodeInfo) GetCloudComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudComplianceCount
}

// GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetCloudComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudComplianceCount, true
}

// SetCloudComplianceCount sets field value
func (o *GraphNodeInfo) SetCloudComplianceCount(v int32) {
	o.CloudComplianceCount = v
}

// GetComplianceCount returns the ComplianceCount field value
func (o *GraphNodeInfo) GetComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComplianceCount
}

// GetComplianceCountOk returns a tuple with the ComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCount, true
}

// SetComplianceCount sets field value
func (o *GraphNodeInfo) SetComplianceCount(v int32) {
	o.ComplianceCount = v
}

// GetName returns the Name field value
func (o *GraphNodeInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GraphNodeInfo) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value
func (o *GraphNodeInfo) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *GraphNodeInfo) SetNodeId(v string) {
	o.NodeId = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *GraphNodeInfo) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *GraphNodeInfo) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetVulnerabilityCount returns the VulnerabilityCount field value
func (o *GraphNodeInfo) GetVulnerabilityCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilityCount
}

// GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field value
// and a boolean to check if the value has been set.
func (o *GraphNodeInfo) GetVulnerabilityCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityCount, true
}

// SetVulnerabilityCount sets field value
func (o *GraphNodeInfo) SetVulnerabilityCount(v int32) {
	o.VulnerabilityCount = v
}

func (o GraphNodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphNodeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_compliance_count"] = o.CloudComplianceCount
	toSerialize["compliance_count"] = o.ComplianceCount
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	toSerialize["secrets_count"] = o.SecretsCount
	toSerialize["vulnerability_count"] = o.VulnerabilityCount
	return toSerialize, nil
}

func (o *GraphNodeInfo) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_compliance_count",
		"compliance_count",
		"name",
		"node_id",
		"secrets_count",
		"vulnerability_count",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGraphNodeInfo := _GraphNodeInfo{}

	err = json.Unmarshal(bytes, &varGraphNodeInfo)

	if err != nil {
		return err
	}

	*o = GraphNodeInfo(varGraphNodeInfo)

	return err
}

type NullableGraphNodeInfo struct {
	value *GraphNodeInfo
	isSet bool
}

func (v NullableGraphNodeInfo) Get() *GraphNodeInfo {
	return v.value
}

func (v *NullableGraphNodeInfo) Set(val *GraphNodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphNodeInfo(val *GraphNodeInfo) *NullableGraphNodeInfo {
	return &NullableGraphNodeInfo{value: val, isSet: true}
}

func (v NullableGraphNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


