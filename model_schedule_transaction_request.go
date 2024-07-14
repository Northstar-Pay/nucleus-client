/*
Blnk

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ScheduleTransactionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleTransactionRequest{}

// ScheduleTransactionRequest struct for ScheduleTransactionRequest
type ScheduleTransactionRequest struct {
	Amount         int32  `json:"amount"`
	Tag            string `json:"tag"`
	Reference      string `json:"reference"`
	Currency       string `json:"currency"`
	Source         string `json:"source"`
	Destination    string `json:"destination"`
	AllowOverDraft bool   `json:"allow_over_draft"`
	ScheduledFor   string `json:"scheduled_for"`
}

type _ScheduleTransactionRequest ScheduleTransactionRequest

// NewScheduleTransactionRequest instantiates a new ScheduleTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleTransactionRequest(amount int32, tag string, reference string, currency string, source string, destination string, allowOverDraft bool, scheduledFor string) *ScheduleTransactionRequest {
	this := ScheduleTransactionRequest{}
	this.Amount = amount
	this.Tag = tag
	this.Reference = reference
	this.Currency = currency
	this.Source = source
	this.Destination = destination
	this.AllowOverDraft = allowOverDraft
	this.ScheduledFor = scheduledFor
	return &this
}

// NewScheduleTransactionRequestWithDefaults instantiates a new ScheduleTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleTransactionRequestWithDefaults() *ScheduleTransactionRequest {
	this := ScheduleTransactionRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ScheduleTransactionRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ScheduleTransactionRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetTag returns the Tag field value
func (o *ScheduleTransactionRequest) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ScheduleTransactionRequest) SetTag(v string) {
	o.Tag = v
}

// GetReference returns the Reference field value
func (o *ScheduleTransactionRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *ScheduleTransactionRequest) SetReference(v string) {
	o.Reference = v
}

// GetCurrency returns the Currency field value
func (o *ScheduleTransactionRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ScheduleTransactionRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetSource returns the Source field value
func (o *ScheduleTransactionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ScheduleTransactionRequest) SetSource(v string) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *ScheduleTransactionRequest) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *ScheduleTransactionRequest) SetDestination(v string) {
	o.Destination = v
}

// GetAllowOverDraft returns the AllowOverDraft field value
func (o *ScheduleTransactionRequest) GetAllowOverDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowOverDraft
}

// GetAllowOverDraftOk returns a tuple with the AllowOverDraft field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetAllowOverDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowOverDraft, true
}

// SetAllowOverDraft sets field value
func (o *ScheduleTransactionRequest) SetAllowOverDraft(v bool) {
	o.AllowOverDraft = v
}

// GetScheduledFor returns the ScheduledFor field value
func (o *ScheduleTransactionRequest) GetScheduledFor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value
// and a boolean to check if the value has been set.
func (o *ScheduleTransactionRequest) GetScheduledForOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledFor, true
}

// SetScheduledFor sets field value
func (o *ScheduleTransactionRequest) SetScheduledFor(v string) {
	o.ScheduledFor = v
}

func (o ScheduleTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleTransactionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["tag"] = o.Tag
	toSerialize["reference"] = o.Reference
	toSerialize["currency"] = o.Currency
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	toSerialize["allow_over_draft"] = o.AllowOverDraft
	toSerialize["scheduled_for"] = o.ScheduledFor
	return toSerialize, nil
}

func (o *ScheduleTransactionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"tag",
		"reference",
		"currency",
		"source",
		"destination",
		"allow_over_draft",
		"scheduled_for",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varScheduleTransactionRequest := _ScheduleTransactionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScheduleTransactionRequest)

	if err != nil {
		return err
	}

	*o = ScheduleTransactionRequest(varScheduleTransactionRequest)

	return err
}

type NullableScheduleTransactionRequest struct {
	value *ScheduleTransactionRequest
	isSet bool
}

func (v NullableScheduleTransactionRequest) Get() *ScheduleTransactionRequest {
	return v.value
}

func (v *NullableScheduleTransactionRequest) Set(val *ScheduleTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleTransactionRequest(val *ScheduleTransactionRequest) *NullableScheduleTransactionRequest {
	return &NullableScheduleTransactionRequest{value: val, isSet: true}
}

func (v NullableScheduleTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
