# CreateBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedgerId** | **string** |  | 
**Currency** | **string** |  | 

## Methods

### NewCreateBalanceRequest

`func NewCreateBalanceRequest(ledgerId string, currency string, ) *CreateBalanceRequest`

NewCreateBalanceRequest instantiates a new CreateBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBalanceRequestWithDefaults

`func NewCreateBalanceRequestWithDefaults() *CreateBalanceRequest`

NewCreateBalanceRequestWithDefaults instantiates a new CreateBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedgerId

`func (o *CreateBalanceRequest) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *CreateBalanceRequest) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *CreateBalanceRequest) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetCurrency

`func (o *CreateBalanceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateBalanceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateBalanceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


