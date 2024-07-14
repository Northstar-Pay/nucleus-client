# RecordTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Reference** | **string** |  | 
**Sources** | [**[]Source**](Source.md) |  | 
**Destinations** | [**[]Destination**](Destination.md) |  | 
**Currency** | **string** |  | 
**Source** | **string** |  | 
**Destination** | **string** |  | 
**AllowOverDraft** | **bool** |  | 
**Inflight** | **bool** |  | 

## Methods

### NewRecordTransactionRequest

`func NewRecordTransactionRequest(amount int32, reference string, sources []Source, destinations []Destination, currency string, source string, destination string, allowOverDraft bool, inflight bool, ) *RecordTransactionRequest`

NewRecordTransactionRequest instantiates a new RecordTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordTransactionRequestWithDefaults

`func NewRecordTransactionRequestWithDefaults() *RecordTransactionRequest`

NewRecordTransactionRequestWithDefaults instantiates a new RecordTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RecordTransactionRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RecordTransactionRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RecordTransactionRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetReference

`func (o *RecordTransactionRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RecordTransactionRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RecordTransactionRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSources

`func (o *RecordTransactionRequest) GetSources() []Source`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RecordTransactionRequest) GetSourcesOk() (*[]Source, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RecordTransactionRequest) SetSources(v []Source)`

SetSources sets Sources field to given value.


### GetDestinations

`func (o *RecordTransactionRequest) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *RecordTransactionRequest) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *RecordTransactionRequest) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.


### GetCurrency

`func (o *RecordTransactionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RecordTransactionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RecordTransactionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSource

`func (o *RecordTransactionRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordTransactionRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordTransactionRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetDestination

`func (o *RecordTransactionRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RecordTransactionRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RecordTransactionRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetAllowOverDraft

`func (o *RecordTransactionRequest) GetAllowOverDraft() bool`

GetAllowOverDraft returns the AllowOverDraft field if non-nil, zero value otherwise.

### GetAllowOverDraftOk

`func (o *RecordTransactionRequest) GetAllowOverDraftOk() (*bool, bool)`

GetAllowOverDraftOk returns a tuple with the AllowOverDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverDraft

`func (o *RecordTransactionRequest) SetAllowOverDraft(v bool)`

SetAllowOverDraft sets AllowOverDraft field to given value.


### GetInflight

`func (o *RecordTransactionRequest) GetInflight() bool`

GetInflight returns the Inflight field if non-nil, zero value otherwise.

### GetInflightOk

`func (o *RecordTransactionRequest) GetInflightOk() (*bool, bool)`

GetInflightOk returns a tuple with the Inflight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflight

`func (o *RecordTransactionRequest) SetInflight(v bool)`

SetInflight sets Inflight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


