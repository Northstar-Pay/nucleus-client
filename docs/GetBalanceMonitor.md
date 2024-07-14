# GetBalanceMonitor

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

### NewGetBalanceMonitor

`func NewGetBalanceMonitor(monitorId string, balanceId string, condition Condition, description string, callBackUrl string, createdAt string, ) *GetBalanceMonitor`

NewGetBalanceMonitor instantiates a new GetBalanceMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceMonitorWithDefaults

`func NewGetBalanceMonitorWithDefaults() *GetBalanceMonitor`

NewGetBalanceMonitorWithDefaults instantiates a new GetBalanceMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *GetBalanceMonitor) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *GetBalanceMonitor) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *GetBalanceMonitor) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.


### GetBalanceId

`func (o *GetBalanceMonitor) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *GetBalanceMonitor) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *GetBalanceMonitor) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetCondition

`func (o *GetBalanceMonitor) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetBalanceMonitor) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetBalanceMonitor) SetCondition(v Condition)`

SetCondition sets Condition field to given value.


### GetDescription

`func (o *GetBalanceMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBalanceMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBalanceMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCallBackUrl

`func (o *GetBalanceMonitor) GetCallBackUrl() string`

GetCallBackUrl returns the CallBackUrl field if non-nil, zero value otherwise.

### GetCallBackUrlOk

`func (o *GetBalanceMonitor) GetCallBackUrlOk() (*string, bool)`

GetCallBackUrlOk returns a tuple with the CallBackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBackUrl

`func (o *GetBalanceMonitor) SetCallBackUrl(v string)`

SetCallBackUrl sets CallBackUrl field to given value.


### GetCreatedAt

`func (o *GetBalanceMonitor) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetBalanceMonitor) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetBalanceMonitor) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


