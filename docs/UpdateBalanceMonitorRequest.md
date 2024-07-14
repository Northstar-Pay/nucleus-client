# UpdateBalanceMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceId** | **string** |  | 
**Condition** | [**Condition**](Condition.md) |  | 

## Methods

### NewUpdateBalanceMonitorRequest

`func NewUpdateBalanceMonitorRequest(balanceId string, condition Condition, ) *UpdateBalanceMonitorRequest`

NewUpdateBalanceMonitorRequest instantiates a new UpdateBalanceMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBalanceMonitorRequestWithDefaults

`func NewUpdateBalanceMonitorRequestWithDefaults() *UpdateBalanceMonitorRequest`

NewUpdateBalanceMonitorRequestWithDefaults instantiates a new UpdateBalanceMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceId

`func (o *UpdateBalanceMonitorRequest) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *UpdateBalanceMonitorRequest) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *UpdateBalanceMonitorRequest) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetCondition

`func (o *UpdateBalanceMonitorRequest) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *UpdateBalanceMonitorRequest) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *UpdateBalanceMonitorRequest) SetCondition(v Condition)`

SetCondition sets Condition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


