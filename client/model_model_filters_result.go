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

// checks if the ModelFiltersResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelFiltersResult{}

// ModelFiltersResult struct for ModelFiltersResult
type ModelFiltersResult struct {
	Filters map[string][]string `json:"filters"`
}

type _ModelFiltersResult ModelFiltersResult

// NewModelFiltersResult instantiates a new ModelFiltersResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelFiltersResult(filters map[string][]string) *ModelFiltersResult {
	this := ModelFiltersResult{}
	this.Filters = filters
	return &this
}

// NewModelFiltersResultWithDefaults instantiates a new ModelFiltersResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelFiltersResultWithDefaults() *ModelFiltersResult {
	this := ModelFiltersResult{}
	return &this
}

// GetFilters returns the Filters field value
// If the value is explicit nil, the zero value for map[string][]string will be returned
func (o *ModelFiltersResult) GetFilters() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelFiltersResult) GetFiltersOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ModelFiltersResult) SetFilters(v map[string][]string) {
	o.Filters = v
}

func (o ModelFiltersResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelFiltersResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

func (o *ModelFiltersResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filters",
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

	varModelFiltersResult := _ModelFiltersResult{}

	err = json.Unmarshal(bytes, &varModelFiltersResult)

	if err != nil {
		return err
	}

	*o = ModelFiltersResult(varModelFiltersResult)

	return err
}

type NullableModelFiltersResult struct {
	value *ModelFiltersResult
	isSet bool
}

func (v NullableModelFiltersResult) Get() *ModelFiltersResult {
	return v.value
}

func (v *NullableModelFiltersResult) Set(val *ModelFiltersResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelFiltersResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelFiltersResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelFiltersResult(val *ModelFiltersResult) *NullableModelFiltersResult {
	return &NullableModelFiltersResult{value: val, isSet: true}
}

func (v NullableModelFiltersResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelFiltersResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


