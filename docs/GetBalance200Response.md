# GetBalance200Response

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
**Ledger** | [**Ledger**](Ledger.md) |  | 
**CreatedAt** | **string** |  | 
**MetaData** | **NullableString** |  | 

## Methods

### NewGetBalance200Response

`func NewGetBalance200Response(balanceId string, balance int32, creditBalance int32, debitBalance int32, currency string, currencyMultiplier int32, ledgerId string, identityId string, ledger Ledger, createdAt string, metaData NullableString, ) *GetBalance200Response`

NewGetBalance200Response instantiates a new GetBalance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalance200ResponseWithDefaults

`func NewGetBalance200ResponseWithDefaults() *GetBalance200Response`

NewGetBalance200ResponseWithDefaults instantiates a new GetBalance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceId

`func (o *GetBalance200Response) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *GetBalance200Response) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *GetBalance200Response) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetBalance

`func (o *GetBalance200Response) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetBalance200Response) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetBalance200Response) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetCreditBalance

`func (o *GetBalance200Response) GetCreditBalance() int32`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *GetBalance200Response) GetCreditBalanceOk() (*int32, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *GetBalance200Response) SetCreditBalance(v int32)`

SetCreditBalance sets CreditBalance field to given value.


### GetDebitBalance

`func (o *GetBalance200Response) GetDebitBalance() int32`

GetDebitBalance returns the DebitBalance field if non-nil, zero value otherwise.

### GetDebitBalanceOk

`func (o *GetBalance200Response) GetDebitBalanceOk() (*int32, bool)`

GetDebitBalanceOk returns a tuple with the DebitBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitBalance

`func (o *GetBalance200Response) SetDebitBalance(v int32)`

SetDebitBalance sets DebitBalance field to given value.


### GetCurrency

`func (o *GetBalance200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBalance200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBalance200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrencyMultiplier

`func (o *GetBalance200Response) GetCurrencyMultiplier() int32`

GetCurrencyMultiplier returns the CurrencyMultiplier field if non-nil, zero value otherwise.

### GetCurrencyMultiplierOk

`func (o *GetBalance200Response) GetCurrencyMultiplierOk() (*int32, bool)`

GetCurrencyMultiplierOk returns a tuple with the CurrencyMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyMultiplier

`func (o *GetBalance200Response) SetCurrencyMultiplier(v int32)`

SetCurrencyMultiplier sets CurrencyMultiplier field to given value.


### GetLedgerId

`func (o *GetBalance200Response) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *GetBalance200Response) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *GetBalance200Response) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetIdentityId

`func (o *GetBalance200Response) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *GetBalance200Response) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *GetBalance200Response) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetLedger

`func (o *GetBalance200Response) GetLedger() Ledger`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *GetBalance200Response) GetLedgerOk() (*Ledger, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *GetBalance200Response) SetLedger(v Ledger)`

SetLedger sets Ledger field to given value.


### GetCreatedAt

`func (o *GetBalance200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetBalance200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetBalance200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetaData

`func (o *GetBalance200Response) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *GetBalance200Response) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *GetBalance200Response) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.


### SetMetaDataNil

`func (o *GetBalance200Response) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *GetBalance200Response) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


