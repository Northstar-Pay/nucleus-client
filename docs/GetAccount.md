# GetAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Name** | **string** |  | 
**Number** | **string** |  | 
**BankName** | **string** |  | 
**Currency** | **string** |  | 
**CreatedAt** | **string** |  | 
**BalanceId** | **string** |  | 
**IdentityId** | **string** |  | 
**LedgerId** | **string** |  | 
**MetaData** | [**MetaData2**](MetaData2.md) |  | 
**Ledger** | [**Ledger**](Ledger.md) |  | 
**Balance** | [**Balance**](Balance.md) |  | 
**Identity** | [**Identity**](Identity.md) |  | 

## Methods

### NewGetAccount

`func NewGetAccount(accountId string, name string, number string, bankName string, currency string, createdAt string, balanceId string, identityId string, ledgerId string, metaData MetaData2, ledger Ledger, balance Balance, identity Identity, ) *GetAccount`

NewGetAccount instantiates a new GetAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountWithDefaults

`func NewGetAccountWithDefaults() *GetAccount`

NewGetAccountWithDefaults instantiates a new GetAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *GetAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAccount) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *GetAccount) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GetAccount) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GetAccount) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetBankName

`func (o *GetAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *GetAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *GetAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetCurrency

`func (o *GetAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *GetAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetBalanceId

`func (o *GetAccount) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *GetAccount) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *GetAccount) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetIdentityId

`func (o *GetAccount) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *GetAccount) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *GetAccount) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetLedgerId

`func (o *GetAccount) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *GetAccount) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *GetAccount) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetMetaData

`func (o *GetAccount) GetMetaData() MetaData2`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *GetAccount) GetMetaDataOk() (*MetaData2, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *GetAccount) SetMetaData(v MetaData2)`

SetMetaData sets MetaData field to given value.


### GetLedger

`func (o *GetAccount) GetLedger() Ledger`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *GetAccount) GetLedgerOk() (*Ledger, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *GetAccount) SetLedger(v Ledger)`

SetLedger sets Ledger field to given value.


### GetBalance

`func (o *GetAccount) GetBalance() Balance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetAccount) GetBalanceOk() (*Balance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetAccount) SetBalance(v Balance)`

SetBalance sets Balance field to given value.


### GetIdentity

`func (o *GetAccount) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *GetAccount) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *GetAccount) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


