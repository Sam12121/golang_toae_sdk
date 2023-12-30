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

// checks if the DiagnosisGenerateCloudScannerDiagnosticLogsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosisGenerateCloudScannerDiagnosticLogsRequest{}

// DiagnosisGenerateCloudScannerDiagnosticLogsRequest struct for DiagnosisGenerateCloudScannerDiagnosticLogsRequest
type DiagnosisGenerateCloudScannerDiagnosticLogsRequest struct {
	NodeIds []DiagnosisNodeIdentifier `json:"node_ids"`
	Tail int32 `json:"tail"`
}

type _DiagnosisGenerateCloudScannerDiagnosticLogsRequest DiagnosisGenerateCloudScannerDiagnosticLogsRequest

// NewDiagnosisGenerateCloudScannerDiagnosticLogsRequest instantiates a new DiagnosisGenerateCloudScannerDiagnosticLogsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosisGenerateCloudScannerDiagnosticLogsRequest(nodeIds []DiagnosisNodeIdentifier, tail int32) *DiagnosisGenerateCloudScannerDiagnosticLogsRequest {
	this := DiagnosisGenerateCloudScannerDiagnosticLogsRequest{}
	this.NodeIds = nodeIds
	this.Tail = tail
	return &this
}

// NewDiagnosisGenerateCloudScannerDiagnosticLogsRequestWithDefaults instantiates a new DiagnosisGenerateCloudScannerDiagnosticLogsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosisGenerateCloudScannerDiagnosticLogsRequestWithDefaults() *DiagnosisGenerateCloudScannerDiagnosticLogsRequest {
	this := DiagnosisGenerateCloudScannerDiagnosticLogsRequest{}
	return &this
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []DiagnosisNodeIdentifier will be returned
func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetNodeIds() []DiagnosisNodeIdentifier {
	if o == nil {
		var ret []DiagnosisNodeIdentifier
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetNodeIdsOk() ([]DiagnosisNodeIdentifier, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) SetNodeIds(v []DiagnosisNodeIdentifier) {
	o.NodeIds = v
}

// GetTail returns the Tail field value
func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetTail() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tail
}

// GetTailOk returns a tuple with the Tail field value
// and a boolean to check if the value has been set.
func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetTailOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tail, true
}

// SetTail sets field value
func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) SetTail(v int32) {
	o.Tail = v
}

func (o DiagnosisGenerateCloudScannerDiagnosticLogsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosisGenerateCloudScannerDiagnosticLogsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	toSerialize["tail"] = o.Tail
	return toSerialize, nil
}

func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_ids",
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

	varDiagnosisGenerateCloudScannerDiagnosticLogsRequest := _DiagnosisGenerateCloudScannerDiagnosticLogsRequest{}

	err = json.Unmarshal(bytes, &varDiagnosisGenerateCloudScannerDiagnosticLogsRequest)

	if err != nil {
		return err
	}

	*o = DiagnosisGenerateCloudScannerDiagnosticLogsRequest(varDiagnosisGenerateCloudScannerDiagnosticLogsRequest)

	return err
}

type NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest struct {
	value *DiagnosisGenerateCloudScannerDiagnosticLogsRequest
	isSet bool
}

func (v NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest) Get() *DiagnosisGenerateCloudScannerDiagnosticLogsRequest {
	return v.value
}

func (v *NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest) Set(val *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest(val *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) *NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest {
	return &NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest{value: val, isSet: true}
}

func (v NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosisGenerateCloudScannerDiagnosticLogsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


