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

// checks if the ModelSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSecret{}

// ModelSecret struct for ModelSecret
type ModelSecret struct {
	FullFilename string `json:"full_filename"`
	Level string `json:"level"`
	Masked bool `json:"masked"`
	MatchedContent string `json:"matched_content"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	Part string `json:"part"`
	RelativeEndingIndex int32 `json:"relative_ending_index"`
	RelativeStartingIndex int32 `json:"relative_starting_index"`
	Resources []ModelBasicNode `json:"resources,omitempty"`
	RuleId int32 `json:"rule_id"`
	Score float32 `json:"score"`
	SignatureToMatch string `json:"signature_to_match"`
	StartingIndex int32 `json:"starting_index"`
	UpdatedAt int32 `json:"updated_at"`
}

type _ModelSecret ModelSecret

// NewModelSecret instantiates a new ModelSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSecret(fullFilename string, level string, masked bool, matchedContent string, name string, nodeId string, part string, relativeEndingIndex int32, relativeStartingIndex int32, ruleId int32, score float32, signatureToMatch string, startingIndex int32, updatedAt int32) *ModelSecret {
	this := ModelSecret{}
	this.FullFilename = fullFilename
	this.Level = level
	this.Masked = masked
	this.MatchedContent = matchedContent
	this.Name = name
	this.NodeId = nodeId
	this.Part = part
	this.RelativeEndingIndex = relativeEndingIndex
	this.RelativeStartingIndex = relativeStartingIndex
	this.RuleId = ruleId
	this.Score = score
	this.SignatureToMatch = signatureToMatch
	this.StartingIndex = startingIndex
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelSecretWithDefaults instantiates a new ModelSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSecretWithDefaults() *ModelSecret {
	this := ModelSecret{}
	return &this
}

// GetFullFilename returns the FullFilename field value
func (o *ModelSecret) GetFullFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullFilename
}

// GetFullFilenameOk returns a tuple with the FullFilename field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetFullFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullFilename, true
}

// SetFullFilename sets field value
func (o *ModelSecret) SetFullFilename(v string) {
	o.FullFilename = v
}

// GetLevel returns the Level field value
func (o *ModelSecret) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ModelSecret) SetLevel(v string) {
	o.Level = v
}

// GetMasked returns the Masked field value
func (o *ModelSecret) GetMasked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetMaskedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelSecret) SetMasked(v bool) {
	o.Masked = v
}

// GetMatchedContent returns the MatchedContent field value
func (o *ModelSecret) GetMatchedContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchedContent
}

// GetMatchedContentOk returns a tuple with the MatchedContent field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetMatchedContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchedContent, true
}

// SetMatchedContent sets field value
func (o *ModelSecret) SetMatchedContent(v string) {
	o.MatchedContent = v
}

// GetName returns the Name field value
func (o *ModelSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelSecret) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value
func (o *ModelSecret) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelSecret) SetNodeId(v string) {
	o.NodeId = v
}

// GetPart returns the Part field value
func (o *ModelSecret) GetPart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Part
}

// GetPartOk returns a tuple with the Part field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetPartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Part, true
}

// SetPart sets field value
func (o *ModelSecret) SetPart(v string) {
	o.Part = v
}

// GetRelativeEndingIndex returns the RelativeEndingIndex field value
func (o *ModelSecret) GetRelativeEndingIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelativeEndingIndex
}

// GetRelativeEndingIndexOk returns a tuple with the RelativeEndingIndex field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetRelativeEndingIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativeEndingIndex, true
}

// SetRelativeEndingIndex sets field value
func (o *ModelSecret) SetRelativeEndingIndex(v int32) {
	o.RelativeEndingIndex = v
}

// GetRelativeStartingIndex returns the RelativeStartingIndex field value
func (o *ModelSecret) GetRelativeStartingIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelativeStartingIndex
}

// GetRelativeStartingIndexOk returns a tuple with the RelativeStartingIndex field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetRelativeStartingIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativeStartingIndex, true
}

// SetRelativeStartingIndex sets field value
func (o *ModelSecret) SetRelativeStartingIndex(v int32) {
	o.RelativeStartingIndex = v
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSecret) GetResources() []ModelBasicNode {
	if o == nil {
		var ret []ModelBasicNode
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecret) GetResourcesOk() ([]ModelBasicNode, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ModelSecret) HasResources() bool {
	if o != nil && IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ModelBasicNode and assigns it to the Resources field.
func (o *ModelSecret) SetResources(v []ModelBasicNode) {
	o.Resources = v
}

// GetRuleId returns the RuleId field value
func (o *ModelSecret) GetRuleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *ModelSecret) SetRuleId(v int32) {
	o.RuleId = v
}

// GetScore returns the Score field value
func (o *ModelSecret) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *ModelSecret) SetScore(v float32) {
	o.Score = v
}

// GetSignatureToMatch returns the SignatureToMatch field value
func (o *ModelSecret) GetSignatureToMatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureToMatch
}

// GetSignatureToMatchOk returns a tuple with the SignatureToMatch field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetSignatureToMatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureToMatch, true
}

// SetSignatureToMatch sets field value
func (o *ModelSecret) SetSignatureToMatch(v string) {
	o.SignatureToMatch = v
}

// GetStartingIndex returns the StartingIndex field value
func (o *ModelSecret) GetStartingIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartingIndex
}

// GetStartingIndexOk returns a tuple with the StartingIndex field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetStartingIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartingIndex, true
}

// SetStartingIndex sets field value
func (o *ModelSecret) SetStartingIndex(v int32) {
	o.StartingIndex = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelSecret) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelSecret) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelSecret) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o ModelSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["full_filename"] = o.FullFilename
	toSerialize["level"] = o.Level
	toSerialize["masked"] = o.Masked
	toSerialize["matched_content"] = o.MatchedContent
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	toSerialize["part"] = o.Part
	toSerialize["relative_ending_index"] = o.RelativeEndingIndex
	toSerialize["relative_starting_index"] = o.RelativeStartingIndex
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	toSerialize["rule_id"] = o.RuleId
	toSerialize["score"] = o.Score
	toSerialize["signature_to_match"] = o.SignatureToMatch
	toSerialize["starting_index"] = o.StartingIndex
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ModelSecret) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"full_filename",
		"level",
		"masked",
		"matched_content",
		"name",
		"node_id",
		"part",
		"relative_ending_index",
		"relative_starting_index",
		"rule_id",
		"score",
		"signature_to_match",
		"starting_index",
		"updated_at",
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

	varModelSecret := _ModelSecret{}

	err = json.Unmarshal(bytes, &varModelSecret)

	if err != nil {
		return err
	}

	*o = ModelSecret(varModelSecret)

	return err
}

type NullableModelSecret struct {
	value *ModelSecret
	isSet bool
}

func (v NullableModelSecret) Get() *ModelSecret {
	return v.value
}

func (v *NullableModelSecret) Set(val *ModelSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSecret(val *ModelSecret) *NullableModelSecret {
	return &NullableModelSecret{value: val, isSet: true}
}

func (v NullableModelSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


