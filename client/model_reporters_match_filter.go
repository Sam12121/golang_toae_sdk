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

// checks if the ReportersMatchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersMatchFilter{}

// ReportersMatchFilter struct for ReportersMatchFilter
type ReportersMatchFilter struct {
	FilterIn map[string][]interface{} `json:"filter_in"`
}

type _ReportersMatchFilter ReportersMatchFilter

// NewReportersMatchFilter instantiates a new ReportersMatchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersMatchFilter(filterIn map[string][]interface{}) *ReportersMatchFilter {
	this := ReportersMatchFilter{}
	this.FilterIn = filterIn
	return &this
}

// NewReportersMatchFilterWithDefaults instantiates a new ReportersMatchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersMatchFilterWithDefaults() *ReportersMatchFilter {
	this := ReportersMatchFilter{}
	return &this
}

// GetFilterIn returns the FilterIn field value
// If the value is explicit nil, the zero value for map[string][]interface{} will be returned
func (o *ReportersMatchFilter) GetFilterIn() map[string][]interface{} {
	if o == nil {
		var ret map[string][]interface{}
		return ret
	}

	return o.FilterIn
}

// GetFilterInOk returns a tuple with the FilterIn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersMatchFilter) GetFilterInOk() (*map[string][]interface{}, bool) {
	if o == nil || IsNil(o.FilterIn) {
		return nil, false
	}
	return &o.FilterIn, true
}

// SetFilterIn sets field value
func (o *ReportersMatchFilter) SetFilterIn(v map[string][]interface{}) {
	o.FilterIn = v
}

func (o ReportersMatchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersMatchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilterIn != nil {
		toSerialize["filter_in"] = o.FilterIn
	}
	return toSerialize, nil
}

func (o *ReportersMatchFilter) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filter_in",
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

	varReportersMatchFilter := _ReportersMatchFilter{}

	err = json.Unmarshal(bytes, &varReportersMatchFilter)

	if err != nil {
		return err
	}

	*o = ReportersMatchFilter(varReportersMatchFilter)

	return err
}

type NullableReportersMatchFilter struct {
	value *ReportersMatchFilter
	isSet bool
}

func (v NullableReportersMatchFilter) Get() *ReportersMatchFilter {
	return v.value
}

func (v *NullableReportersMatchFilter) Set(val *ReportersMatchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersMatchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersMatchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersMatchFilter(val *ReportersMatchFilter) *NullableReportersMatchFilter {
	return &NullableReportersMatchFilter{value: val, isSet: true}
}

func (v NullableReportersMatchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersMatchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


