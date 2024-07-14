# CreateBalanceMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceId** | **string** |  | 
**Condition** | [**Condition**](Condition.md) |  | 
**CallBackUrl** | **string** |  | 

## Methods

### NewCreateBalanceMonitorRequest

`func NewCreateBalanceMonitorRequest(balanceId string, condition Condition, callBackUrl string, ) *CreateBalanceMonitorRequest`

NewCreateBalanceMonitorRequest instantiates a new CreateBalanceMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBalanceMonitorRequestWithDefaults

`func NewCreateBalanceMonitorRequestWithDefaults() *CreateBalanceMonitorRequest`

NewCreateBalanceMonitorRequestWithDefaults instantiates a new CreateBalanceMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceId

`func (o *CreateBalanceMonitorRequest) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *CreateBalanceMonitorRequest) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *CreateBalanceMonitorRequest) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetCondition

`func (o *CreateBalanceMonitorRequest) GetCondition() Condition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateBalanceMonitorRequest) GetConditionOk() (*Condition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateBalanceMonitorRequest) SetCondition(v Condition)`

SetCondition sets Condition field to given value.


### GetCallBackUrl

`func (o *CreateBalanceMonitorRequest) GetCallBackUrl() string`

GetCallBackUrl returns the CallBackUrl field if non-nil, zero value otherwise.

### GetCallBackUrlOk

`func (o *CreateBalanceMonitorRequest) GetCallBackUrlOk() (*string, bool)`

GetCallBackUrlOk returns a tuple with the CallBackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallBackUrl

`func (o *CreateBalanceMonitorRequest) SetCallBackUrl(v string)`

SetCallBackUrl sets CallBackUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


