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

// checks if the CreateBalanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBalanceRequest{}

// CreateBalanceRequest struct for CreateBalanceRequest
type CreateBalanceRequest struct {
	LedgerId string `json:"ledger_id"`
	Currency string `json:"currency"`
}

type _CreateBalanceRequest CreateBalanceRequest

// NewCreateBalanceRequest instantiates a new CreateBalanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBalanceRequest(ledgerId string, currency string) *CreateBalanceRequest {
	this := CreateBalanceRequest{}
	this.LedgerId = ledgerId
	this.Currency = currency
	return &this
}

// NewCreateBalanceRequestWithDefaults instantiates a new CreateBalanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBalanceRequestWithDefaults() *CreateBalanceRequest {
	this := CreateBalanceRequest{}
	return &this
}

// GetLedgerId returns the LedgerId field value
func (o *CreateBalanceRequest) GetLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value
// and a boolean to check if the value has been set.
func (o *CreateBalanceRequest) GetLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerId, true
}

// SetLedgerId sets field value
func (o *CreateBalanceRequest) SetLedgerId(v string) {
	o.LedgerId = v
}

// GetCurrency returns the Currency field value
func (o *CreateBalanceRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateBalanceRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateBalanceRequest) SetCurrency(v string) {
	o.Currency = v
}

func (o CreateBalanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBalanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ledger_id"] = o.LedgerId
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *CreateBalanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ledger_id",
		"currency",
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

	varCreateBalanceRequest := _CreateBalanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBalanceRequest)

	if err != nil {
		return err
	}

	*o = CreateBalanceRequest(varCreateBalanceRequest)

	return err
}

type NullableCreateBalanceRequest struct {
	value *CreateBalanceRequest
	isSet bool
}

func (v NullableCreateBalanceRequest) Get() *CreateBalanceRequest {
	return v.value
}

func (v *NullableCreateBalanceRequest) Set(val *CreateBalanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBalanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBalanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBalanceRequest(val *CreateBalanceRequest) *NullableCreateBalanceRequest {
	return &NullableCreateBalanceRequest{value: val, isSet: true}
}

func (v NullableCreateBalanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBalanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
