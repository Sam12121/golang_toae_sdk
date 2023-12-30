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

// checks if the GraphIndividualThreatGraphRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphIndividualThreatGraphRequest{}

// GraphIndividualThreatGraphRequest struct for GraphIndividualThreatGraphRequest
type GraphIndividualThreatGraphRequest struct {
	GraphType string `json:"graph_type"`
	IssueType string `json:"issue_type"`
	NodeIds []string `json:"node_ids,omitempty"`
}

type _GraphIndividualThreatGraphRequest GraphIndividualThreatGraphRequest

// NewGraphIndividualThreatGraphRequest instantiates a new GraphIndividualThreatGraphRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphIndividualThreatGraphRequest(graphType string, issueType string) *GraphIndividualThreatGraphRequest {
	this := GraphIndividualThreatGraphRequest{}
	this.GraphType = graphType
	this.IssueType = issueType
	return &this
}

// NewGraphIndividualThreatGraphRequestWithDefaults instantiates a new GraphIndividualThreatGraphRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphIndividualThreatGraphRequestWithDefaults() *GraphIndividualThreatGraphRequest {
	this := GraphIndividualThreatGraphRequest{}
	return &this
}

// GetGraphType returns the GraphType field value
func (o *GraphIndividualThreatGraphRequest) GetGraphType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GraphType
}

// GetGraphTypeOk returns a tuple with the GraphType field value
// and a boolean to check if the value has been set.
func (o *GraphIndividualThreatGraphRequest) GetGraphTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GraphType, true
}

// SetGraphType sets field value
func (o *GraphIndividualThreatGraphRequest) SetGraphType(v string) {
	o.GraphType = v
}

// GetIssueType returns the IssueType field value
func (o *GraphIndividualThreatGraphRequest) GetIssueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value
// and a boolean to check if the value has been set.
func (o *GraphIndividualThreatGraphRequest) GetIssueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueType, true
}

// SetIssueType sets field value
func (o *GraphIndividualThreatGraphRequest) SetIssueType(v string) {
	o.IssueType = v
}

// GetNodeIds returns the NodeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphIndividualThreatGraphRequest) GetNodeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphIndividualThreatGraphRequest) GetNodeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// HasNodeIds returns a boolean if a field has been set.
func (o *GraphIndividualThreatGraphRequest) HasNodeIds() bool {
	if o != nil && IsNil(o.NodeIds) {
		return true
	}

	return false
}

// SetNodeIds gets a reference to the given []string and assigns it to the NodeIds field.
func (o *GraphIndividualThreatGraphRequest) SetNodeIds(v []string) {
	o.NodeIds = v
}

func (o GraphIndividualThreatGraphRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphIndividualThreatGraphRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["graph_type"] = o.GraphType
	toSerialize["issue_type"] = o.IssueType
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	return toSerialize, nil
}

func (o *GraphIndividualThreatGraphRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"graph_type",
		"issue_type",
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

	varGraphIndividualThreatGraphRequest := _GraphIndividualThreatGraphRequest{}

	err = json.Unmarshal(bytes, &varGraphIndividualThreatGraphRequest)

	if err != nil {
		return err
	}

	*o = GraphIndividualThreatGraphRequest(varGraphIndividualThreatGraphRequest)

	return err
}

type NullableGraphIndividualThreatGraphRequest struct {
	value *GraphIndividualThreatGraphRequest
	isSet bool
}

func (v NullableGraphIndividualThreatGraphRequest) Get() *GraphIndividualThreatGraphRequest {
	return v.value
}

func (v *NullableGraphIndividualThreatGraphRequest) Set(val *GraphIndividualThreatGraphRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphIndividualThreatGraphRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphIndividualThreatGraphRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphIndividualThreatGraphRequest(val *GraphIndividualThreatGraphRequest) *NullableGraphIndividualThreatGraphRequest {
	return &NullableGraphIndividualThreatGraphRequest{value: val, isSet: true}
}

func (v NullableGraphIndividualThreatGraphRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphIndividualThreatGraphRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

