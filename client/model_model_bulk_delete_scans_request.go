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

// checks if the ModelBulkDeleteScansRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelBulkDeleteScansRequest{}

// ModelBulkDeleteScansRequest struct for ModelBulkDeleteScansRequest
type ModelBulkDeleteScansRequest struct {
	Filters ReportersFieldsFilters `json:"filters"`
	ScanType string `json:"scan_type"`
}

type _ModelBulkDeleteScansRequest ModelBulkDeleteScansRequest

// NewModelBulkDeleteScansRequest instantiates a new ModelBulkDeleteScansRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelBulkDeleteScansRequest(filters ReportersFieldsFilters, scanType string) *ModelBulkDeleteScansRequest {
	this := ModelBulkDeleteScansRequest{}
	this.Filters = filters
	this.ScanType = scanType
	return &this
}

// NewModelBulkDeleteScansRequestWithDefaults instantiates a new ModelBulkDeleteScansRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelBulkDeleteScansRequestWithDefaults() *ModelBulkDeleteScansRequest {
	this := ModelBulkDeleteScansRequest{}
	return &this
}

// GetFilters returns the Filters field value
func (o *ModelBulkDeleteScansRequest) GetFilters() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ModelBulkDeleteScansRequest) GetFiltersOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ModelBulkDeleteScansRequest) SetFilters(v ReportersFieldsFilters) {
	o.Filters = v
}

// GetScanType returns the ScanType field value
func (o *ModelBulkDeleteScansRequest) GetScanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value
// and a boolean to check if the value has been set.
func (o *ModelBulkDeleteScansRequest) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanType, true
}

// SetScanType sets field value
func (o *ModelBulkDeleteScansRequest) SetScanType(v string) {
	o.ScanType = v
}

func (o ModelBulkDeleteScansRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelBulkDeleteScansRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filters"] = o.Filters
	toSerialize["scan_type"] = o.ScanType
	return toSerialize, nil
}

func (o *ModelBulkDeleteScansRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filters",
		"scan_type",
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

	varModelBulkDeleteScansRequest := _ModelBulkDeleteScansRequest{}

	err = json.Unmarshal(bytes, &varModelBulkDeleteScansRequest)

	if err != nil {
		return err
	}

	*o = ModelBulkDeleteScansRequest(varModelBulkDeleteScansRequest)

	return err
}

type NullableModelBulkDeleteScansRequest struct {
	value *ModelBulkDeleteScansRequest
	isSet bool
}

func (v NullableModelBulkDeleteScansRequest) Get() *ModelBulkDeleteScansRequest {
	return v.value
}

func (v *NullableModelBulkDeleteScansRequest) Set(val *ModelBulkDeleteScansRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBulkDeleteScansRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBulkDeleteScansRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBulkDeleteScansRequest(val *ModelBulkDeleteScansRequest) *NullableModelBulkDeleteScansRequest {
	return &NullableModelBulkDeleteScansRequest{value: val, isSet: true}
}

func (v NullableModelBulkDeleteScansRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBulkDeleteScansRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


