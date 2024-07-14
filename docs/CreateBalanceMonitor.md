# CreateBalanceMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorId** | **string** |  | 
**BalanceId** | **string** |  | 
**Condition** | [**Condition**](Condition.md) |  | 
**Description** | **string** |  | 
**CallBackUrl** | **string** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewCreateBalanceMonitor

`func NewCreateBalanceMonitor(monitorId string, balanceId string, condition Condition, description string, callBackUrl string, createdAt string, ) *CreateBalanceMonitor`

NewCreateBalanceMonitor instantiates a new CreateBalanceMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBalanceMonitorWithDefaults

`func NewCreateBalanceMonitorWithDefaults() *CreateBalanceMonitor`

NewCreateBalanceMonitorWithDefaults instantiates a new CreateBalanceMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *CreateBalanceMonitor) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *CreateBalanceMonitor) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *CreateBalanceMonitor) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.


### GetBalanceId

`func (o *CreateBalanceMonitor) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *CreateBalanceMonitor) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *CreateBalanceMonitor) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetCondition

`func (o *CreateBalanceMonitor) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateBalanceMonitor) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateBalanceMonitor) SetCondition(v Condition)`

SetCondition sets Condition field to given value.


### GetDescription

`func (o *CreateBalanceMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBalanceMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBalanceMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCallBackUrl

`func (o *CreateBalanceMonitor) GetCallBackUrl() string`

GetCallBackUrl returns the CallBackUrl field if non-nil, zero value otherwise.

### GetCallBackUrlOk

`func (o *CreateBalanceMonitor) GetCallBackUrlOk() (*string, bool)`

GetCallBackUrlOk returns a tuple with the CallBackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBackUrl

`func (o *CreateBalanceMonitor) SetCallBackUrl(v string)`

SetCallBackUrl sets CallBackUrl field to given value.


### GetCreatedAt

`func (o *CreateBalanceMonitor) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateBalanceMonitor) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateBalanceMonitor) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


