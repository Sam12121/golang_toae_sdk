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
)

// checks if the SearchResultGroupResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResultGroupResp{}

// SearchResultGroupResp struct for SearchResultGroupResp
type SearchResultGroupResp struct {
	Groups []SearchResultGroup `json:"groups,omitempty"`
}

// NewSearchResultGroupResp instantiates a new SearchResultGroupResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultGroupResp() *SearchResultGroupResp {
	this := SearchResultGroupResp{}
	return &this
}

// NewSearchResultGroupRespWithDefaults instantiates a new SearchResultGroupResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultGroupRespWithDefaults() *SearchResultGroupResp {
	this := SearchResultGroupResp{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchResultGroupResp) GetGroups() []SearchResultGroup {
	if o == nil {
		var ret []SearchResultGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchResultGroupResp) GetGroupsOk() ([]SearchResultGroup, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SearchResultGroupResp) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []SearchResultGroup and assigns it to the Groups field.
func (o *SearchResultGroupResp) SetGroups(v []SearchResultGroup) {
	o.Groups = v
}

func (o SearchResultGroupResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResultGroupResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableSearchResultGroupResp struct {
	value *SearchResultGroupResp
	isSet bool
}

func (v NullableSearchResultGroupResp) Get() *SearchResultGroupResp {
	return v.value
}

func (v *NullableSearchResultGroupResp) Set(val *SearchResultGroupResp) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultGroupResp) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultGroupResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultGroupResp(val *SearchResultGroupResp) *NullableSearchResultGroupResp {
	return &NullableSearchResultGroupResp{value: val, isSet: true}
}

func (v NullableSearchResultGroupResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultGroupResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


