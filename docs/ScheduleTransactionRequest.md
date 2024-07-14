# ScheduleTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Tag** | **string** |  | 
**Reference** | **string** |  | 
**Currency** | **string** |  | 
**Source** | **string** |  | 
**Destination** | **string** |  | 
**AllowOverDraft** | **bool** |  | 
**ScheduledFor** | **string** |  | 

## Methods

### NewScheduleTransactionRequest

`func NewScheduleTransactionRequest(amount int32, tag string, reference string, currency string, source string, destination string, allowOverDraft bool, scheduledFor string, ) *ScheduleTransactionRequest`

NewScheduleTransactionRequest instantiates a new ScheduleTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTransactionRequestWithDefaults

`func NewScheduleTransactionRequestWithDefaults() *ScheduleTransactionRequest`

NewScheduleTransactionRequestWithDefaults instantiates a new ScheduleTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ScheduleTransactionRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ScheduleTransactionRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ScheduleTransactionRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetTag

`func (o *ScheduleTransactionRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ScheduleTransactionRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ScheduleTransactionRequest) SetTag(v string)`

SetTag sets Tag field to given value.


### GetReference

`func (o *ScheduleTransactionRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ScheduleTransactionRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ScheduleTransactionRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetCurrency

`func (o *ScheduleTransactionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ScheduleTransactionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ScheduleTransactionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSource

`func (o *ScheduleTransactionRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ScheduleTransactionRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ScheduleTransactionRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ScheduleTransactionRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ScheduleTransactionRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ScheduleTransactionRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetAllowOverDraft

`func (o *ScheduleTransactionRequest) GetAllowOverDraft() bool`

GetAllowOverDraft returns the AllowOverDraft field if non-nil, zero value otherwise.

### GetAllowOverDraftOk

`func (o *ScheduleTransactionRequest) GetAllowOverDraftOk() (*bool, bool)`

GetAllowOverDraftOk returns a tuple with the AllowOverDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverDraft

`func (o *ScheduleTransactionRequest) SetAllowOverDraft(v bool)`

SetAllowOverDraft sets AllowOverDraft field to given value.


### GetScheduledFor

`func (o *ScheduleTransactionRequest) GetScheduledFor() string`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *ScheduleTransactionRequest) GetScheduledForOk() (*string, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *ScheduleTransactionRequest) SetScheduledFor(v string)`

SetScheduledFor sets ScheduledFor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


