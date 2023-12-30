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

// checks if the ApiDocsFailureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDocsFailureResponse{}

// ApiDocsFailureResponse struct for ApiDocsFailureResponse
type ApiDocsFailureResponse struct {
	Message *string `json:"message,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewApiDocsFailureResponse instantiates a new ApiDocsFailureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDocsFailureResponse() *ApiDocsFailureResponse {
	this := ApiDocsFailureResponse{}
	return &this
}

// NewApiDocsFailureResponseWithDefaults instantiates a new ApiDocsFailureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDocsFailureResponseWithDefaults() *ApiDocsFailureResponse {
	this := ApiDocsFailureResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiDocsFailureResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDocsFailureResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiDocsFailureResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiDocsFailureResponse) SetMessage(v string) {
	o.Message = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ApiDocsFailureResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDocsFailureResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ApiDocsFailureResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ApiDocsFailureResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o ApiDocsFailureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDocsFailureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableApiDocsFailureResponse struct {
	value *ApiDocsFailureResponse
	isSet bool
}

func (v NullableApiDocsFailureResponse) Get() *ApiDocsFailureResponse {
	return v.value
}

func (v *NullableApiDocsFailureResponse) Set(val *ApiDocsFailureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDocsFailureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDocsFailureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDocsFailureResponse(val *ApiDocsFailureResponse) *NullableApiDocsFailureResponse {
	return &NullableApiDocsFailureResponse{value: val, isSet: true}
}

func (v NullableApiDocsFailureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDocsFailureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


