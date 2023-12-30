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

// checks if the ModelDeleteRegistryBulkReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelDeleteRegistryBulkReq{}

// ModelDeleteRegistryBulkReq struct for ModelDeleteRegistryBulkReq
type ModelDeleteRegistryBulkReq struct {
	RegistryIds []string `json:"registry_ids"`
}

type _ModelDeleteRegistryBulkReq ModelDeleteRegistryBulkReq

// NewModelDeleteRegistryBulkReq instantiates a new ModelDeleteRegistryBulkReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelDeleteRegistryBulkReq(registryIds []string) *ModelDeleteRegistryBulkReq {
	this := ModelDeleteRegistryBulkReq{}
	this.RegistryIds = registryIds
	return &this
}

// NewModelDeleteRegistryBulkReqWithDefaults instantiates a new ModelDeleteRegistryBulkReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelDeleteRegistryBulkReqWithDefaults() *ModelDeleteRegistryBulkReq {
	this := ModelDeleteRegistryBulkReq{}
	return &this
}

// GetRegistryIds returns the RegistryIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelDeleteRegistryBulkReq) GetRegistryIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RegistryIds
}

// GetRegistryIdsOk returns a tuple with the RegistryIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelDeleteRegistryBulkReq) GetRegistryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RegistryIds) {
		return nil, false
	}
	return o.RegistryIds, true
}

// SetRegistryIds sets field value
func (o *ModelDeleteRegistryBulkReq) SetRegistryIds(v []string) {
	o.RegistryIds = v
}

func (o ModelDeleteRegistryBulkReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelDeleteRegistryBulkReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RegistryIds != nil {
		toSerialize["registry_ids"] = o.RegistryIds
	}
	return toSerialize, nil
}

func (o *ModelDeleteRegistryBulkReq) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"registry_ids",
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

	varModelDeleteRegistryBulkReq := _ModelDeleteRegistryBulkReq{}

	err = json.Unmarshal(bytes, &varModelDeleteRegistryBulkReq)

	if err != nil {
		return err
	}

	*o = ModelDeleteRegistryBulkReq(varModelDeleteRegistryBulkReq)

	return err
}

type NullableModelDeleteRegistryBulkReq struct {
	value *ModelDeleteRegistryBulkReq
	isSet bool
}

func (v NullableModelDeleteRegistryBulkReq) Get() *ModelDeleteRegistryBulkReq {
	return v.value
}

func (v *NullableModelDeleteRegistryBulkReq) Set(val *ModelDeleteRegistryBulkReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDeleteRegistryBulkReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDeleteRegistryBulkReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDeleteRegistryBulkReq(val *ModelDeleteRegistryBulkReq) *NullableModelDeleteRegistryBulkReq {
	return &NullableModelDeleteRegistryBulkReq{value: val, isSet: true}
}

func (v NullableModelDeleteRegistryBulkReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDeleteRegistryBulkReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


