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

// checks if the DiagnosisGenerateConsoleDiagnosticLogsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosisGenerateConsoleDiagnosticLogsRequest{}

// DiagnosisGenerateConsoleDiagnosticLogsRequest struct for DiagnosisGenerateConsoleDiagnosticLogsRequest
type DiagnosisGenerateConsoleDiagnosticLogsRequest struct {
	Tail int32 `json:"tail"`
}

type _DiagnosisGenerateConsoleDiagnosticLogsRequest DiagnosisGenerateConsoleDiagnosticLogsRequest

// NewDiagnosisGenerateConsoleDiagnosticLogsRequest instantiates a new DiagnosisGenerateConsoleDiagnosticLogsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosisGenerateConsoleDiagnosticLogsRequest(tail int32) *DiagnosisGenerateConsoleDiagnosticLogsRequest {
	this := DiagnosisGenerateConsoleDiagnosticLogsRequest{}
	this.Tail = tail
	return &this
}

// NewDiagnosisGenerateConsoleDiagnosticLogsRequestWithDefaults instantiates a new DiagnosisGenerateConsoleDiagnosticLogsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosisGenerateConsoleDiagnosticLogsRequestWithDefaults() *DiagnosisGenerateConsoleDiagnosticLogsRequest {
	this := DiagnosisGenerateConsoleDiagnosticLogsRequest{}
	return &this
}

// GetTail returns the Tail field value
func (o *DiagnosisGenerateConsoleDiagnosticLogsRequest) GetTail() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tail
}

// GetTailOk returns a tuple with the Tail field value
// and a boolean to check if the value has been set.
func (o *DiagnosisGenerateConsoleDiagnosticLogsRequest) GetTailOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tail, true
}

// SetTail sets field value
func (o *DiagnosisGenerateConsoleDiagnosticLogsRequest) SetTail(v int32) {
	o.Tail = v
}

func (o DiagnosisGenerateConsoleDiagnosticLogsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosisGenerateConsoleDiagnosticLogsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tail"] = o.Tail
	return toSerialize, nil
}

func (o *DiagnosisGenerateConsoleDiagnosticLogsRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tail",
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

	varDiagnosisGenerateConsoleDiagnosticLogsRequest := _DiagnosisGenerateConsoleDiagnosticLogsRequest{}

	err = json.Unmarshal(bytes, &varDiagnosisGenerateConsoleDiagnosticLogsRequest)

	if err != nil {
		return err
	}

	*o = DiagnosisGenerateConsoleDiagnosticLogsRequest(varDiagnosisGenerateConsoleDiagnosticLogsRequest)

	return err
}

type NullableDiagnosisGenerateConsoleDiagnosticLogsRequest struct {
	value *DiagnosisGenerateConsoleDiagnosticLogsRequest
	isSet bool
}

func (v NullableDiagnosisGenerateConsoleDiagnosticLogsRequest) Get() *DiagnosisGenerateConsoleDiagnosticLogsRequest {
	return v.value
}

func (v *NullableDiagnosisGenerateConsoleDiagnosticLogsRequest) Set(val *DiagnosisGenerateConsoleDiagnosticLogsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosisGenerateConsoleDiagnosticLogsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosisGenerateConsoleDiagnosticLogsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosisGenerateConsoleDiagnosticLogsRequest(val *DiagnosisGenerateConsoleDiagnosticLogsRequest) *NullableDiagnosisGenerateConsoleDiagnosticLogsRequest {
	return &NullableDiagnosisGenerateConsoleDiagnosticLogsRequest{value: val, isSet: true}
}

func (v NullableDiagnosisGenerateConsoleDiagnosticLogsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosisGenerateConsoleDiagnosticLogsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


