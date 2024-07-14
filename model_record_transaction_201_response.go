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

// checks if the RecordTransaction201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordTransaction201Response{}

// RecordTransaction201Response struct for RecordTransaction201Response
type RecordTransaction201Response struct {
	Amount         int32          `json:"amount"`
	Id             string         `json:"id"`
	AllowOverdraft bool           `json:"allow_overdraft"`
	Infligt        bool           `json:"infligt"`
	Source         string         `json:"source"`
	Destination    string         `json:"destination"`
	Reference      string         `json:"reference"`
	Currency       string         `json:"currency"`
	Description    string         `json:"description"`
	Status         string         `json:"status"`
	Hash           string         `json:"hash"`
	GroupIds       NullableString `json:"group_ids"`
	CreatedAt      string         `json:"created_at"`
	ScheduledFor   string         `json:"scheduled_for"`
}

type _RecordTransaction201Response RecordTransaction201Response

// NewRecordTransaction201Response instantiates a new RecordTransaction201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordTransaction201Response(amount int32, id string, allowOverdraft bool, infligt bool, source string, destination string, reference string, currency string, description string, status string, hash string, groupIds NullableString, createdAt string, scheduledFor string) *RecordTransaction201Response {
	this := RecordTransaction201Response{}
	this.Amount = amount
	this.Id = id
	this.AllowOverdraft = allowOverdraft
	this.Infligt = infligt
	this.Source = source
	this.Destination = destination
	this.Reference = reference
	this.Currency = currency
	this.Description = description
	this.Status = status
	this.Hash = hash
	this.GroupIds = groupIds
	this.CreatedAt = createdAt
	this.ScheduledFor = scheduledFor
	return &this
}

// NewRecordTransaction201ResponseWithDefaults instantiates a new RecordTransaction201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordTransaction201ResponseWithDefaults() *RecordTransaction201Response {
	this := RecordTransaction201Response{}
	return &this
}

// GetAmount returns the Amount field value
func (o *RecordTransaction201Response) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RecordTransaction201Response) SetAmount(v int32) {
	o.Amount = v
}

// GetId returns the Id field value
func (o *RecordTransaction201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RecordTransaction201Response) SetId(v string) {
	o.Id = v
}

// GetAllowOverdraft returns the AllowOverdraft field value
func (o *RecordTransaction201Response) GetAllowOverdraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowOverdraft
}

// GetAllowOverdraftOk returns a tuple with the AllowOverdraft field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetAllowOverdraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowOverdraft, true
}

// SetAllowOverdraft sets field value
func (o *RecordTransaction201Response) SetAllowOverdraft(v bool) {
	o.AllowOverdraft = v
}

// GetInfligt returns the Infligt field value
func (o *RecordTransaction201Response) GetInfligt() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Infligt
}

// GetInfligtOk returns a tuple with the Infligt field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetInfligtOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Infligt, true
}

// SetInfligt sets field value
func (o *RecordTransaction201Response) SetInfligt(v bool) {
	o.Infligt = v
}

// GetSource returns the Source field value
func (o *RecordTransaction201Response) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *RecordTransaction201Response) SetSource(v string) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *RecordTransaction201Response) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *RecordTransaction201Response) SetDestination(v string) {
	o.Destination = v
}

// GetReference returns the Reference field value
func (o *RecordTransaction201Response) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *RecordTransaction201Response) SetReference(v string) {
	o.Reference = v
}

// GetCurrency returns the Currency field value
func (o *RecordTransaction201Response) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RecordTransaction201Response) SetCurrency(v string) {
	o.Currency = v
}

// GetDescription returns the Description field value
func (o *RecordTransaction201Response) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RecordTransaction201Response) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value
func (o *RecordTransaction201Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RecordTransaction201Response) SetStatus(v string) {
	o.Status = v
}

// GetHash returns the Hash field value
func (o *RecordTransaction201Response) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *RecordTransaction201Response) SetHash(v string) {
	o.Hash = v
}

// GetGroupIds returns the GroupIds field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RecordTransaction201Response) GetGroupIds() string {
	if o == nil || o.GroupIds.Get() == nil {
		var ret string
		return ret
	}

	return *o.GroupIds.Get()
}

// GetGroupIdsOk returns a tuple with the GroupIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecordTransaction201Response) GetGroupIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupIds.Get(), o.GroupIds.IsSet()
}

// SetGroupIds sets field value
func (o *RecordTransaction201Response) SetGroupIds(v string) {
	o.GroupIds.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *RecordTransaction201Response) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RecordTransaction201Response) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetScheduledFor returns the ScheduledFor field value
func (o *RecordTransaction201Response) GetScheduledFor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value
// and a boolean to check if the value has been set.
func (o *RecordTransaction201Response) GetScheduledForOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledFor, true
}

// SetScheduledFor sets field value
func (o *RecordTransaction201Response) SetScheduledFor(v string) {
	o.ScheduledFor = v
}

func (o RecordTransaction201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordTransaction201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["id"] = o.Id
	toSerialize["allow_overdraft"] = o.AllowOverdraft
	toSerialize["infligt"] = o.Infligt
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	toSerialize["reference"] = o.Reference
	toSerialize["currency"] = o.Currency
	toSerialize["description"] = o.Description
	toSerialize["status"] = o.Status
	toSerialize["hash"] = o.Hash
	toSerialize["group_ids"] = o.GroupIds.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["scheduled_for"] = o.ScheduledFor
	return toSerialize, nil
}

func (o *RecordTransaction201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"id",
		"allow_overdraft",
		"infligt",
		"source",
		"destination",
		"reference",
		"currency",
		"description",
		"status",
		"hash",
		"group_ids",
		"created_at",
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

	varRecordTransaction201Response := _RecordTransaction201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecordTransaction201Response)

	if err != nil {
		return err
	}

	*o = RecordTransaction201Response(varRecordTransaction201Response)

	return err
}

type NullableRecordTransaction201Response struct {
	value *RecordTransaction201Response
	isSet bool
}

func (v NullableRecordTransaction201Response) Get() *RecordTransaction201Response {
	return v.value
}

func (v *NullableRecordTransaction201Response) Set(val *RecordTransaction201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordTransaction201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordTransaction201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordTransaction201Response(val *RecordTransaction201Response) *NullableRecordTransaction201Response {
	return &NullableRecordTransaction201Response{value: val, isSet: true}
}

func (v NullableRecordTransaction201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordTransaction201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
