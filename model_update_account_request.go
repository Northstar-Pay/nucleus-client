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

// checks if the UpdateAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccountRequest{}

// UpdateAccountRequest struct for UpdateAccountRequest
type UpdateAccountRequest struct {
	BankName   string `json:"bank_name"`
	Number     string `json:"number"`
	BalanceId  string `json:"balance_id"`
	IdentityId string `json:"identity_id"`
}

type _UpdateAccountRequest UpdateAccountRequest

// NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountRequest(bankName string, number string, balanceId string, identityId string) *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	this.BankName = bankName
	this.Number = number
	this.BalanceId = balanceId
	this.IdentityId = identityId
	return &this
}

// NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	return &this
}

// GetBankName returns the BankName field value
func (o *UpdateAccountRequest) GetBankName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankName, true
}

// SetBankName sets field value
func (o *UpdateAccountRequest) SetBankName(v string) {
	o.BankName = v
}

// GetNumber returns the Number field value
func (o *UpdateAccountRequest) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *UpdateAccountRequest) SetNumber(v string) {
	o.Number = v
}

// GetBalanceId returns the BalanceId field value
func (o *UpdateAccountRequest) GetBalanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceId
}

// GetBalanceIdOk returns a tuple with the BalanceId field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetBalanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceId, true
}

// SetBalanceId sets field value
func (o *UpdateAccountRequest) SetBalanceId(v string) {
	o.BalanceId = v
}

// GetIdentityId returns the IdentityId field value
func (o *UpdateAccountRequest) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *UpdateAccountRequest) SetIdentityId(v string) {
	o.IdentityId = v
}

func (o UpdateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bank_name"] = o.BankName
	toSerialize["number"] = o.Number
	toSerialize["balance_id"] = o.BalanceId
	toSerialize["identity_id"] = o.IdentityId
	return toSerialize, nil
}

func (o *UpdateAccountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bank_name",
		"number",
		"balance_id",
		"identity_id",
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

	varUpdateAccountRequest := _UpdateAccountRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateAccountRequest)

	if err != nil {
		return err
	}

	*o = UpdateAccountRequest(varUpdateAccountRequest)

	return err
}

type NullableUpdateAccountRequest struct {
	value *UpdateAccountRequest
	isSet bool
}

func (v NullableUpdateAccountRequest) Get() *UpdateAccountRequest {
	return v.value
}

func (v *NullableUpdateAccountRequest) Set(val *UpdateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountRequest(val *UpdateAccountRequest) *NullableUpdateAccountRequest {
	return &NullableUpdateAccountRequest{value: val, isSet: true}
}

func (v NullableUpdateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
