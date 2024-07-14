# RecordTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Id** | **string** |  | 
**AllowOverdraft** | **bool** |  | 
**Infligt** | **bool** |  | 
**Source** | **string** |  | 
**Destination** | **string** |  | 
**Reference** | **string** |  | 
**Currency** | **string** |  | 
**Description** | **string** |  | 
**Status** | **string** |  | 
**Hash** | **string** |  | 
**GroupIds** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**ScheduledFor** | **string** |  | 

## Methods

### NewRecordTransaction

`func NewRecordTransaction(amount int32, id string, allowOverdraft bool, infligt bool, source string, destination string, reference string, currency string, description string, status string, hash string, groupIds NullableString, createdAt string, scheduledFor string, ) *RecordTransaction`

NewRecordTransaction instantiates a new RecordTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordTransactionWithDefaults

`func NewRecordTransactionWithDefaults() *RecordTransaction`

NewRecordTransactionWithDefaults instantiates a new RecordTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RecordTransaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RecordTransaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RecordTransaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *RecordTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetAllowOverdraft

`func (o *RecordTransaction) GetAllowOverdraft() bool`

GetAllowOverdraft returns the AllowOverdraft field if non-nil, zero value otherwise.

### GetAllowOverdraftOk

`func (o *RecordTransaction) GetAllowOverdraftOk() (*bool, bool)`

GetAllowOverdraftOk returns a tuple with the AllowOverdraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverdraft

`func (o *RecordTransaction) SetAllowOverdraft(v bool)`

SetAllowOverdraft sets AllowOverdraft field to given value.


### GetInfligt

`func (o *RecordTransaction) GetInfligt() bool`

GetInfligt returns the Infligt field if non-nil, zero value otherwise.

### GetInfligtOk

`func (o *RecordTransaction) GetInfligtOk() (*bool, bool)`

GetInfligtOk returns a tuple with the Infligt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfligt

`func (o *RecordTransaction) SetInfligt(v bool)`

SetInfligt sets Infligt field to given value.


### GetSource

`func (o *RecordTransaction) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordTransaction) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordTransaction) SetSource(v string)`

SetSource sets Source field to given value.


### GetDestination

`func (o *RecordTransaction) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RecordTransaction) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RecordTransaction) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetReference

`func (o *RecordTransaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RecordTransaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RecordTransaction) SetReference(v string)`

SetReference sets Reference field to given value.


### GetCurrency

`func (o *RecordTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RecordTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RecordTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *RecordTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecordTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecordTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *RecordTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecordTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecordTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHash

`func (o *RecordTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RecordTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RecordTransaction) SetHash(v string)`

SetHash sets Hash field to given value.


### GetGroupIds

`func (o *RecordTransaction) GetGroupIds() string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *RecordTransaction) GetGroupIdsOk() (*string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *RecordTransaction) SetGroupIds(v string)`

SetGroupIds sets GroupIds field to given value.


### SetGroupIdsNil

`func (o *RecordTransaction) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *RecordTransaction) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil
### GetCreatedAt

`func (o *RecordTransaction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RecordTransaction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RecordTransaction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetScheduledFor

`func (o *RecordTransaction) GetScheduledFor() string`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *RecordTransaction) GetScheduledForOk() (*string, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *RecordTransaction) SetScheduledFor(v string)`

SetScheduledFor sets ScheduledFor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


