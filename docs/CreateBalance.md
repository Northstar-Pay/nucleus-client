# CreateBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceId** | **string** |  | 
**Balance** | **int32** |  | 
**CreditBalance** | **int32** |  | 
**DebitBalance** | **int32** |  | 
**Currency** | **string** |  | 
**CurrencyMultiplier** | **int32** |  | 
**LedgerId** | **string** |  | 
**IdentityId** | **string** |  | 
**CreatedAt** | **string** |  | 
**MetaData** | **NullableString** |  | 

## Methods

### NewCreateBalance

`func NewCreateBalance(balanceId string, balance int32, creditBalance int32, debitBalance int32, currency string, currencyMultiplier int32, ledgerId string, identityId string, createdAt string, metaData NullableString, ) *CreateBalance`

NewCreateBalance instantiates a new CreateBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBalanceWithDefaults

`func NewCreateBalanceWithDefaults() *CreateBalance`

NewCreateBalanceWithDefaults instantiates a new CreateBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceId

`func (o *CreateBalance) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *CreateBalance) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *CreateBalance) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetBalance

`func (o *CreateBalance) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CreateBalance) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CreateBalance) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetCreditBalance

`func (o *CreateBalance) GetCreditBalance() int32`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *CreateBalance) GetCreditBalanceOk() (*int32, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *CreateBalance) SetCreditBalance(v int32)`

SetCreditBalance sets CreditBalance field to given value.


### GetDebitBalance

`func (o *CreateBalance) GetDebitBalance() int32`

GetDebitBalance returns the DebitBalance field if non-nil, zero value otherwise.

### GetDebitBalanceOk

`func (o *CreateBalance) GetDebitBalanceOk() (*int32, bool)`

GetDebitBalanceOk returns a tuple with the DebitBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitBalance

`func (o *CreateBalance) SetDebitBalance(v int32)`

SetDebitBalance sets DebitBalance field to given value.


### GetCurrency

`func (o *CreateBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrencyMultiplier

`func (o *CreateBalance) GetCurrencyMultiplier() int32`

GetCurrencyMultiplier returns the CurrencyMultiplier field if non-nil, zero value otherwise.

### GetCurrencyMultiplierOk

`func (o *CreateBalance) GetCurrencyMultiplierOk() (*int32, bool)`

GetCurrencyMultiplierOk returns a tuple with the CurrencyMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyMultiplier

`func (o *CreateBalance) SetCurrencyMultiplier(v int32)`

SetCurrencyMultiplier sets CurrencyMultiplier field to given value.


### GetLedgerId

`func (o *CreateBalance) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *CreateBalance) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *CreateBalance) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetIdentityId

`func (o *CreateBalance) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateBalance) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateBalance) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetCreatedAt

`func (o *CreateBalance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateBalance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateBalance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetaData

`func (o *CreateBalance) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CreateBalance) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CreateBalance) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.


### SetMetaDataNil

`func (o *CreateBalance) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *CreateBalance) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


