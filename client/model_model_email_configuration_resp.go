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

// checks if the ModelEmailConfigurationResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelEmailConfigurationResp{}

// ModelEmailConfigurationResp struct for ModelEmailConfigurationResp
type ModelEmailConfigurationResp struct {
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	EmailId *string `json:"email_id,omitempty"`
	EmailProvider *string `json:"email_provider,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Port *string `json:"port,omitempty"`
	SesRegion *string `json:"ses_region,omitempty"`
	Smtp *string `json:"smtp,omitempty"`
}

// NewModelEmailConfigurationResp instantiates a new ModelEmailConfigurationResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelEmailConfigurationResp() *ModelEmailConfigurationResp {
	this := ModelEmailConfigurationResp{}
	return &this
}

// NewModelEmailConfigurationRespWithDefaults instantiates a new ModelEmailConfigurationResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelEmailConfigurationRespWithDefaults() *ModelEmailConfigurationResp {
	this := ModelEmailConfigurationResp{}
	return &this
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *ModelEmailConfigurationResp) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationResp) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *ModelEmailConfigurationResp) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *ModelEmailConfigurationResp) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *ModelEmailConfigurationResp) GetEmailId() string {
	if o == nil || IsNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationResp) GetEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmailId) {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *ModelEmailConfigurationResp) HasEmailId() bool {
	if o != nil && !IsNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *ModelEmailConfigurationResp) SetEmailId(v string) {
	o.EmailId = &v
}

// GetEmailProvider returns the EmailProvider field value if set, zero value otherwise.
func (o *ModelEmailConfigurationResp) GetEmailProvider() string {
	if o == nil || IsNil(o.EmailProvider) {
		var ret string
		return ret
	}
	return *o.EmailProvider
}

// GetEmailProviderOk returns a tuple with the EmailProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationResp) GetEmailProviderOk() (*string, bool) {
	if o == nil || IsNil(o.EmailProvider) {
		return nil, false
	}
	return o.EmailProvider, true
}

// HasEmailProvider returns a boolean if a field has been set.
func (o *ModelEmailConfigurationResp) HasEmailProvider() bool {
	if o != nil && !IsNil(o.EmailProvider) {
		return true
	}

	return false
}

// SetEmailProvider gets a reference to the given string and assigns it to the EmailProvider field.
func (o *ModelEmailConfigurationResp) SetEmailProvider(v string) {
	o.EmailProvider = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelEmailConfigurationResp) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationResp) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelEmailConfigurationResp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelEmailConfigurationResp) SetId(v int32) {
	o.Id = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModelEmailConfigurationResp) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationResp) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModelEmailConfigurationResp) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *ModelEmailConfigurationResp) SetPort(v string) {
	o.Port = &v
}

// GetSesRegion returns the SesRegion field value if set, zero value otherwise.
func (o *ModelEmailConfigurationResp) GetSesRegion() string {
	if o == nil || IsNil(o.SesRegion) {
		var ret string
		return ret
	}
	return *o.SesRegion
}

// GetSesRegionOk returns a tuple with the SesRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationResp) GetSesRegionOk() (*string, bool) {
	if o == nil || IsNil(o.SesRegion) {
		return nil, false
	}
	return o.SesRegion, true
}

// HasSesRegion returns a boolean if a field has been set.
func (o *ModelEmailConfigurationResp) HasSesRegion() bool {
	if o != nil && !IsNil(o.SesRegion) {
		return true
	}

	return false
}

// SetSesRegion gets a reference to the given string and assigns it to the SesRegion field.
func (o *ModelEmailConfigurationResp) SetSesRegion(v string) {
	o.SesRegion = &v
}

// GetSmtp returns the Smtp field value if set, zero value otherwise.
func (o *ModelEmailConfigurationResp) GetSmtp() string {
	if o == nil || IsNil(o.Smtp) {
		var ret string
		return ret
	}
	return *o.Smtp
}

// GetSmtpOk returns a tuple with the Smtp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationResp) GetSmtpOk() (*string, bool) {
	if o == nil || IsNil(o.Smtp) {
		return nil, false
	}
	return o.Smtp, true
}

// HasSmtp returns a boolean if a field has been set.
func (o *ModelEmailConfigurationResp) HasSmtp() bool {
	if o != nil && !IsNil(o.Smtp) {
		return true
	}

	return false
}

// SetSmtp gets a reference to the given string and assigns it to the Smtp field.
func (o *ModelEmailConfigurationResp) SetSmtp(v string) {
	o.Smtp = &v
}

func (o ModelEmailConfigurationResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelEmailConfigurationResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if !IsNil(o.EmailId) {
		toSerialize["email_id"] = o.EmailId
	}
	if !IsNil(o.EmailProvider) {
		toSerialize["email_provider"] = o.EmailProvider
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SesRegion) {
		toSerialize["ses_region"] = o.SesRegion
	}
	if !IsNil(o.Smtp) {
		toSerialize["smtp"] = o.Smtp
	}
	return toSerialize, nil
}

type NullableModelEmailConfigurationResp struct {
	value *ModelEmailConfigurationResp
	isSet bool
}

func (v NullableModelEmailConfigurationResp) Get() *ModelEmailConfigurationResp {
	return v.value
}

func (v *NullableModelEmailConfigurationResp) Set(val *ModelEmailConfigurationResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelEmailConfigurationResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelEmailConfigurationResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelEmailConfigurationResp(val *ModelEmailConfigurationResp) *NullableModelEmailConfigurationResp {
	return &NullableModelEmailConfigurationResp{value: val, isSet: true}
}

func (v NullableModelEmailConfigurationResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelEmailConfigurationResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


