# CreateAccount

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
**MetaData** | **NullableString** |  | 
**Ledger** | **NullableString** |  | 
**Balance** | **NullableString** |  | 
**Identity** | **NullableString** |  | 

## Methods

### NewCreateAccount

`func NewCreateAccount(accountId string, name string, number string, bankName string, currency string, createdAt string, balanceId string, identityId string, ledgerId string, metaData NullableString, ledger NullableString, balance NullableString, identity NullableString, ) *CreateAccount`

NewCreateAccount instantiates a new CreateAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountWithDefaults

`func NewCreateAccountWithDefaults() *CreateAccount`

NewCreateAccountWithDefaults instantiates a new CreateAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *CreateAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccount) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *CreateAccount) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateAccount) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateAccount) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetBankName

`func (o *CreateAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *CreateAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *CreateAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetCurrency

`func (o *CreateAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *CreateAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetBalanceId

`func (o *CreateAccount) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *CreateAccount) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *CreateAccount) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetIdentityId

`func (o *CreateAccount) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateAccount) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateAccount) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetLedgerId

`func (o *CreateAccount) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *CreateAccount) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *CreateAccount) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetMetaData

`func (o *CreateAccount) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CreateAccount) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CreateAccount) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.


### SetMetaDataNil

`func (o *CreateAccount) SetMetaDataNil(b bool)`

 SetMetaDataNil sets the value for MetaData to be an explicit nil

### UnsetMetaData
`func (o *CreateAccount) UnsetMetaData()`

UnsetMetaData ensures that no value is present for MetaData, not even an explicit nil
### GetLedger

`func (o *CreateAccount) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *CreateAccount) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *CreateAccount) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### SetLedgerNil

`func (o *CreateAccount) SetLedgerNil(b bool)`

 SetLedgerNil sets the value for Ledger to be an explicit nil

### UnsetLedger
`func (o *CreateAccount) UnsetLedger()`

UnsetLedger ensures that no value is present for Ledger, not even an explicit nil
### GetBalance

`func (o *CreateAccount) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CreateAccount) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CreateAccount) SetBalance(v string)`

SetBalance sets Balance field to given value.


### SetBalanceNil

`func (o *CreateAccount) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *CreateAccount) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetIdentity

`func (o *CreateAccount) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CreateAccount) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CreateAccount) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### SetIdentityNil

`func (o *CreateAccount) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *CreateAccount) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


