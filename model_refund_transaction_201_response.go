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

// checks if the RefundTransaction201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundTransaction201Response{}

// RefundTransaction201Response struct for RefundTransaction201Response
type RefundTransaction201Response struct {
	Id                     string         `json:"id"`
	Tag                    string         `json:"tag"`
	Reference              string         `json:"reference"`
	Amount                 int32          `json:"amount"`
	Currency               string         `json:"currency"`
	PaymentMethod          string         `json:"payment_method"`
	Description            string         `json:"description"`
	Drcr                   string         `json:"drcr"`
	Status                 string         `json:"status"`
	LedgerId               string         `json:"ledger_id"`
	BalanceId              string         `json:"balance_id"`
	CreditBalanceBefore    int32          `json:"credit_balance_before"`
	DebitBalanceBefore     int32          `json:"debit_balance_before"`
	CreditBalanceAfter     int32          `json:"credit_balance_after"`
	DebitBalanceAfter      int32          `json:"debit_balance_after"`
	BalanceBefore          int32          `json:"balance_before"`
	BalanceAfter           int32          `json:"balance_after"`
	CreatedAt              string         `json:"created_at"`
	ScheduledFor           string         `json:"scheduled_for"`
	RiskToleranceThreshold int32          `json:"risk_tolerance_threshold"`
	RiskScore              float32        `json:"risk_score"`
	MetaData               MetaData1      `json:"meta_data"`
	GroupIds               NullableString `json:"group_ids"`
}

type _RefundTransaction201Response RefundTransaction201Response

// NewRefundTransaction201Response instantiates a new RefundTransaction201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundTransaction201Response(id string, tag string, reference string, amount int32, currency string, paymentMethod string, description string, drcr string, status string, ledgerId string, balanceId string, creditBalanceBefore int32, debitBalanceBefore int32, creditBalanceAfter int32, debitBalanceAfter int32, balanceBefore int32, balanceAfter int32, createdAt string, scheduledFor string, riskToleranceThreshold int32, riskScore float32, metaData MetaData1, groupIds NullableString) *RefundTransaction201Response {
	this := RefundTransaction201Response{}
	this.Id = id
	this.Tag = tag
	this.Reference = reference
	this.Amount = amount
	this.Currency = currency
	this.PaymentMethod = paymentMethod
	this.Description = description
	this.Drcr = drcr
	this.Status = status
	this.LedgerId = ledgerId
	this.BalanceId = balanceId
	this.CreditBalanceBefore = creditBalanceBefore
	this.DebitBalanceBefore = debitBalanceBefore
	this.CreditBalanceAfter = creditBalanceAfter
	this.DebitBalanceAfter = debitBalanceAfter
	this.BalanceBefore = balanceBefore
	this.BalanceAfter = balanceAfter
	this.CreatedAt = createdAt
	this.ScheduledFor = scheduledFor
	this.RiskToleranceThreshold = riskToleranceThreshold
	this.RiskScore = riskScore
	this.MetaData = metaData
	this.GroupIds = groupIds
	return &this
}

// NewRefundTransaction201ResponseWithDefaults instantiates a new RefundTransaction201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundTransaction201ResponseWithDefaults() *RefundTransaction201Response {
	this := RefundTransaction201Response{}
	return &this
}

// GetId returns the Id field value
func (o *RefundTransaction201Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RefundTransaction201Response) SetId(v string) {
	o.Id = v
}

// GetTag returns the Tag field value
func (o *RefundTransaction201Response) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *RefundTransaction201Response) SetTag(v string) {
	o.Tag = v
}

// GetReference returns the Reference field value
func (o *RefundTransaction201Response) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *RefundTransaction201Response) SetReference(v string) {
	o.Reference = v
}

// GetAmount returns the Amount field value
func (o *RefundTransaction201Response) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RefundTransaction201Response) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *RefundTransaction201Response) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RefundTransaction201Response) SetCurrency(v string) {
	o.Currency = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *RefundTransaction201Response) GetPaymentMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *RefundTransaction201Response) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetDescription returns the Description field value
func (o *RefundTransaction201Response) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RefundTransaction201Response) SetDescription(v string) {
	o.Description = v
}

// GetDrcr returns the Drcr field value
func (o *RefundTransaction201Response) GetDrcr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Drcr
}

// GetDrcrOk returns a tuple with the Drcr field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetDrcrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Drcr, true
}

// SetDrcr sets field value
func (o *RefundTransaction201Response) SetDrcr(v string) {
	o.Drcr = v
}

// GetStatus returns the Status field value
func (o *RefundTransaction201Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RefundTransaction201Response) SetStatus(v string) {
	o.Status = v
}

// GetLedgerId returns the LedgerId field value
func (o *RefundTransaction201Response) GetLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LedgerId, true
}

// SetLedgerId sets field value
func (o *RefundTransaction201Response) SetLedgerId(v string) {
	o.LedgerId = v
}

// GetBalanceId returns the BalanceId field value
func (o *RefundTransaction201Response) GetBalanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceId
}

// GetBalanceIdOk returns a tuple with the BalanceId field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetBalanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceId, true
}

// SetBalanceId sets field value
func (o *RefundTransaction201Response) SetBalanceId(v string) {
	o.BalanceId = v
}

// GetCreditBalanceBefore returns the CreditBalanceBefore field value
func (o *RefundTransaction201Response) GetCreditBalanceBefore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreditBalanceBefore
}

// GetCreditBalanceBeforeOk returns a tuple with the CreditBalanceBefore field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetCreditBalanceBeforeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditBalanceBefore, true
}

// SetCreditBalanceBefore sets field value
func (o *RefundTransaction201Response) SetCreditBalanceBefore(v int32) {
	o.CreditBalanceBefore = v
}

// GetDebitBalanceBefore returns the DebitBalanceBefore field value
func (o *RefundTransaction201Response) GetDebitBalanceBefore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DebitBalanceBefore
}

// GetDebitBalanceBeforeOk returns a tuple with the DebitBalanceBefore field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetDebitBalanceBeforeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebitBalanceBefore, true
}

// SetDebitBalanceBefore sets field value
func (o *RefundTransaction201Response) SetDebitBalanceBefore(v int32) {
	o.DebitBalanceBefore = v
}

// GetCreditBalanceAfter returns the CreditBalanceAfter field value
func (o *RefundTransaction201Response) GetCreditBalanceAfter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreditBalanceAfter
}

// GetCreditBalanceAfterOk returns a tuple with the CreditBalanceAfter field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetCreditBalanceAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditBalanceAfter, true
}

// SetCreditBalanceAfter sets field value
func (o *RefundTransaction201Response) SetCreditBalanceAfter(v int32) {
	o.CreditBalanceAfter = v
}

// GetDebitBalanceAfter returns the DebitBalanceAfter field value
func (o *RefundTransaction201Response) GetDebitBalanceAfter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DebitBalanceAfter
}

// GetDebitBalanceAfterOk returns a tuple with the DebitBalanceAfter field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetDebitBalanceAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebitBalanceAfter, true
}

// SetDebitBalanceAfter sets field value
func (o *RefundTransaction201Response) SetDebitBalanceAfter(v int32) {
	o.DebitBalanceAfter = v
}

// GetBalanceBefore returns the BalanceBefore field value
func (o *RefundTransaction201Response) GetBalanceBefore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BalanceBefore
}

// GetBalanceBeforeOk returns a tuple with the BalanceBefore field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetBalanceBeforeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceBefore, true
}

// SetBalanceBefore sets field value
func (o *RefundTransaction201Response) SetBalanceBefore(v int32) {
	o.BalanceBefore = v
}

// GetBalanceAfter returns the BalanceAfter field value
func (o *RefundTransaction201Response) GetBalanceAfter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BalanceAfter
}

// GetBalanceAfterOk returns a tuple with the BalanceAfter field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetBalanceAfterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceAfter, true
}

// SetBalanceAfter sets field value
func (o *RefundTransaction201Response) SetBalanceAfter(v int32) {
	o.BalanceAfter = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RefundTransaction201Response) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RefundTransaction201Response) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetScheduledFor returns the ScheduledFor field value
func (o *RefundTransaction201Response) GetScheduledFor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetScheduledForOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledFor, true
}

// SetScheduledFor sets field value
func (o *RefundTransaction201Response) SetScheduledFor(v string) {
	o.ScheduledFor = v
}

// GetRiskToleranceThreshold returns the RiskToleranceThreshold field value
func (o *RefundTransaction201Response) GetRiskToleranceThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RiskToleranceThreshold
}

// GetRiskToleranceThresholdOk returns a tuple with the RiskToleranceThreshold field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetRiskToleranceThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskToleranceThreshold, true
}

// SetRiskToleranceThreshold sets field value
func (o *RefundTransaction201Response) SetRiskToleranceThreshold(v int32) {
	o.RiskToleranceThreshold = v
}

// GetRiskScore returns the RiskScore field value
func (o *RefundTransaction201Response) GetRiskScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetRiskScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskScore, true
}

// SetRiskScore sets field value
func (o *RefundTransaction201Response) SetRiskScore(v float32) {
	o.RiskScore = v
}

// GetMetaData returns the MetaData field value
func (o *RefundTransaction201Response) GetMetaData() MetaData1 {
	if o == nil {
		var ret MetaData1
		return ret
	}

	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value
// and a boolean to check if the value has been set.
func (o *RefundTransaction201Response) GetMetaDataOk() (*MetaData1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaData, true
}

// SetMetaData sets field value
func (o *RefundTransaction201Response) SetMetaData(v MetaData1) {
	o.MetaData = v
}

// GetGroupIds returns the GroupIds field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RefundTransaction201Response) GetGroupIds() string {
	if o == nil || o.GroupIds.Get() == nil {
		var ret string
		return ret
	}

	return *o.GroupIds.Get()
}

// GetGroupIdsOk returns a tuple with the GroupIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundTransaction201Response) GetGroupIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupIds.Get(), o.GroupIds.IsSet()
}

// SetGroupIds sets field value
func (o *RefundTransaction201Response) SetGroupIds(v string) {
	o.GroupIds.Set(&v)
}

func (o RefundTransaction201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundTransaction201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["tag"] = o.Tag
	toSerialize["reference"] = o.Reference
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	toSerialize["payment_method"] = o.PaymentMethod
	toSerialize["description"] = o.Description
	toSerialize["drcr"] = o.Drcr
	toSerialize["status"] = o.Status
	toSerialize["ledger_id"] = o.LedgerId
	toSerialize["balance_id"] = o.BalanceId
	toSerialize["credit_balance_before"] = o.CreditBalanceBefore
	toSerialize["debit_balance_before"] = o.DebitBalanceBefore
	toSerialize["credit_balance_after"] = o.CreditBalanceAfter
	toSerialize["debit_balance_after"] = o.DebitBalanceAfter
	toSerialize["balance_before"] = o.BalanceBefore
	toSerialize["balance_after"] = o.BalanceAfter
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["scheduled_for"] = o.ScheduledFor
	toSerialize["risk_tolerance_threshold"] = o.RiskToleranceThreshold
	toSerialize["risk_score"] = o.RiskScore
	toSerialize["meta_data"] = o.MetaData
	toSerialize["group_ids"] = o.GroupIds.Get()
	return toSerialize, nil
}

func (o *RefundTransaction201Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"tag",
		"reference",
		"amount",
		"currency",
		"payment_method",
		"description",
		"drcr",
		"status",
		"ledger_id",
		"balance_id",
		"credit_balance_before",
		"debit_balance_before",
		"credit_balance_after",
		"debit_balance_after",
		"balance_before",
		"balance_after",
		"created_at",
		"scheduled_for",
		"risk_tolerance_threshold",
		"risk_score",
		"meta_data",
		"group_ids",
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

	varRefundTransaction201Response := _RefundTransaction201Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefundTransaction201Response)

	if err != nil {
		return err
	}

	*o = RefundTransaction201Response(varRefundTransaction201Response)

	return err
}

type NullableRefundTransaction201Response struct {
	value *RefundTransaction201Response
	isSet bool
}

func (v NullableRefundTransaction201Response) Get() *RefundTransaction201Response {
	return v.value
}

func (v *NullableRefundTransaction201Response) Set(val *RefundTransaction201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundTransaction201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundTransaction201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundTransaction201Response(val *RefundTransaction201Response) *NullableRefundTransaction201Response {
	return &NullableRefundTransaction201Response{value: val, isSet: true}
}

func (v NullableRefundTransaction201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundTransaction201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
