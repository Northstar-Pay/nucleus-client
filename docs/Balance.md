# Balance

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

### NewBalance

`func NewBalance(balanceId string, balance int32, creditBalance int32, debitBalance int32, currency string, currencyMultiplier int32, ledgerId string, identityId string, createdAt string, metaData NullableString, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceId

`func (o *Balance) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *Balance) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *Balance) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetBalance

`func (o *Balance) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Balance) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Balance) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetCreditBalance

`func (o *Balance) GetCreditBalance() int32`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *Balance) GetCreditBalanceOk() (*int32, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *Balance) SetCreditBalance(v int32)`

SetCreditBalance sets CreditBalance field to given value.


### GetDebitBalance

`func (o *Balance) GetDebitBalance() int32`

GetDebitBalance returns the DebitBalance field if non-nil, zero value otherwise.

### GetDebitBalanceOk

`func (o *Balance) GetDebitBalanceOk() (*int32, bool)`

GetDebitBalanceOk returns a tuple with the DebitBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitBalance

`func (o *Balance) SetDebitBalance(v int32)`

SetDebitBalance sets DebitBalance field to given value.


### GetCurrency

`func (o *Balance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Balance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Balance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrencyMultiplier

`func (o *Balance) GetCurrencyMultiplier() int32`

GetCurrencyMultiplier returns the CurrencyMultiplier field if non-nil, zero value otherwise.

### GetCurrencyMultiplierOk

`func (o *Balance) GetCurrencyMultiplierOk() (*int32, bool)`

GetCurrencyMultiplierOk returns a tuple with the CurrencyMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyMultiplier

`func (o *Balance) SetCurrencyMultiplier(v int32)`

SetCurrencyMultiplier sets CurrencyMultiplier field to given value.


### GetLedgerId

`func (o *Balance) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *Balance) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *Balance) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetIdentityId

`func (o *Balance) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Balance) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Balance) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetCreatedAt

`func (o *Balance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Balance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Balance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetaData

`func (o *Balance) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *Balance) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *Balance) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.


### SetMetaDataNil

`func (o *Balance) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *Balance) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


