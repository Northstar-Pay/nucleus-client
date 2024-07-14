# RefundTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Tag** | **string** |  | 
**Reference** | **string** |  | 
**Amount** | **int32** |  | 
**Currency** | **string** |  | 
**PaymentMethod** | **string** |  | 
**Description** | **string** |  | 
**Drcr** | **string** |  | 
**Status** | **string** |  | 
**LedgerId** | **string** |  | 
**BalanceId** | **string** |  | 
**CreditBalanceBefore** | **int32** |  | 
**DebitBalanceBefore** | **int32** |  | 
**CreditBalanceAfter** | **int32** |  | 
**DebitBalanceAfter** | **int32** |  | 
**BalanceBefore** | **int32** |  | 
**BalanceAfter** | **int32** |  | 
**CreatedAt** | **string** |  | 
**ScheduledFor** | **string** |  | 
**RiskToleranceThreshold** | **int32** |  | 
**RiskScore** | **float32** |  | 
**MetaData** | [**MetaData1**](MetaData1.md) |  | 
**GroupIds** | **NullableString** |  | 

## Methods

### NewRefundTransaction

`func NewRefundTransaction(id string, tag string, reference string, amount int32, currency string, paymentMethod string, description string, drcr string, status string, ledgerId string, balanceId string, creditBalanceBefore int32, debitBalanceBefore int32, creditBalanceAfter int32, debitBalanceAfter int32, balanceBefore int32, balanceAfter int32, createdAt string, scheduledFor string, riskToleranceThreshold int32, riskScore float32, metaData MetaData1, groupIds NullableString, ) *RefundTransaction`

NewRefundTransaction instantiates a new RefundTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundTransactionWithDefaults

`func NewRefundTransactionWithDefaults() *RefundTransaction`

NewRefundTransactionWithDefaults instantiates a new RefundTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RefundTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RefundTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RefundTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetTag

`func (o *RefundTransaction) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RefundTransaction) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RefundTransaction) SetTag(v string)`

SetTag sets Tag field to given value.


### GetReference

`func (o *RefundTransaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RefundTransaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RefundTransaction) SetReference(v string)`

SetReference sets Reference field to given value.


### GetAmount

`func (o *RefundTransaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RefundTransaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RefundTransaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *RefundTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RefundTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RefundTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *RefundTransaction) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *RefundTransaction) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *RefundTransaction) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDescription

`func (o *RefundTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefundTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefundTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDrcr

`func (o *RefundTransaction) GetDrcr() string`

GetDrcr returns the Drcr field if non-nil, zero value otherwise.

### GetDrcrOk

`func (o *RefundTransaction) GetDrcrOk() (*string, bool)`

GetDrcrOk returns a tuple with the Drcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrcr

`func (o *RefundTransaction) SetDrcr(v string)`

SetDrcr sets Drcr field to given value.


### GetStatus

`func (o *RefundTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefundTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefundTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLedgerId

`func (o *RefundTransaction) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *RefundTransaction) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *RefundTransaction) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetBalanceId

`func (o *RefundTransaction) GetBalanceId() string`

GetBalanceId returns the BalanceId field if non-nil, zero value otherwise.

### GetBalanceIdOk

`func (o *RefundTransaction) GetBalanceIdOk() (*string, bool)`

GetBalanceIdOk returns a tuple with the BalanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceId

`func (o *RefundTransaction) SetBalanceId(v string)`

SetBalanceId sets BalanceId field to given value.


### GetCreditBalanceBefore

`func (o *RefundTransaction) GetCreditBalanceBefore() int32`

GetCreditBalanceBefore returns the CreditBalanceBefore field if non-nil, zero value otherwise.

### GetCreditBalanceBeforeOk

`func (o *RefundTransaction) GetCreditBalanceBeforeOk() (*int32, bool)`

GetCreditBalanceBeforeOk returns a tuple with the CreditBalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceBefore

`func (o *RefundTransaction) SetCreditBalanceBefore(v int32)`

SetCreditBalanceBefore sets CreditBalanceBefore field to given value.


### GetDebitBalanceBefore

`func (o *RefundTransaction) GetDebitBalanceBefore() int32`

GetDebitBalanceBefore returns the DebitBalanceBefore field if non-nil, zero value otherwise.

### GetDebitBalanceBeforeOk

`func (o *RefundTransaction) GetDebitBalanceBeforeOk() (*int32, bool)`

GetDebitBalanceBeforeOk returns a tuple with the DebitBalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitBalanceBefore

`func (o *RefundTransaction) SetDebitBalanceBefore(v int32)`

SetDebitBalanceBefore sets DebitBalanceBefore field to given value.


### GetCreditBalanceAfter

`func (o *RefundTransaction) GetCreditBalanceAfter() int32`

GetCreditBalanceAfter returns the CreditBalanceAfter field if non-nil, zero value otherwise.

### GetCreditBalanceAfterOk

`func (o *RefundTransaction) GetCreditBalanceAfterOk() (*int32, bool)`

GetCreditBalanceAfterOk returns a tuple with the CreditBalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceAfter

`func (o *RefundTransaction) SetCreditBalanceAfter(v int32)`

SetCreditBalanceAfter sets CreditBalanceAfter field to given value.


### GetDebitBalanceAfter

`func (o *RefundTransaction) GetDebitBalanceAfter() int32`

GetDebitBalanceAfter returns the DebitBalanceAfter field if non-nil, zero value otherwise.

### GetDebitBalanceAfterOk

`func (o *RefundTransaction) GetDebitBalanceAfterOk() (*int32, bool)`

GetDebitBalanceAfterOk returns a tuple with the DebitBalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitBalanceAfter

`func (o *RefundTransaction) SetDebitBalanceAfter(v int32)`

SetDebitBalanceAfter sets DebitBalanceAfter field to given value.


### GetBalanceBefore

`func (o *RefundTransaction) GetBalanceBefore() int32`

GetBalanceBefore returns the BalanceBefore field if non-nil, zero value otherwise.

### GetBalanceBeforeOk

`func (o *RefundTransaction) GetBalanceBeforeOk() (*int32, bool)`

GetBalanceBeforeOk returns a tuple with the BalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBefore

`func (o *RefundTransaction) SetBalanceBefore(v int32)`

SetBalanceBefore sets BalanceBefore field to given value.


### GetBalanceAfter

`func (o *RefundTransaction) GetBalanceAfter() int32`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *RefundTransaction) GetBalanceAfterOk() (*int32, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *RefundTransaction) SetBalanceAfter(v int32)`

SetBalanceAfter sets BalanceAfter field to given value.


### GetCreatedAt

`func (o *RefundTransaction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RefundTransaction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RefundTransaction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetScheduledFor

`func (o *RefundTransaction) GetScheduledFor() string`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *RefundTransaction) GetScheduledForOk() (*string, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *RefundTransaction) SetScheduledFor(v string)`

SetScheduledFor sets ScheduledFor field to given value.


### GetRiskToleranceThreshold

`func (o *RefundTransaction) GetRiskToleranceThreshold() int32`

GetRiskToleranceThreshold returns the RiskToleranceThreshold field if non-nil, zero value otherwise.

### GetRiskToleranceThresholdOk

`func (o *RefundTransaction) GetRiskToleranceThresholdOk() (*int32, bool)`

GetRiskToleranceThresholdOk returns a tuple with the RiskToleranceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskToleranceThreshold

`func (o *RefundTransaction) SetRiskToleranceThreshold(v int32)`

SetRiskToleranceThreshold sets RiskToleranceThreshold field to given value.


### GetRiskScore

`func (o *RefundTransaction) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *RefundTransaction) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *RefundTransaction) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.


### GetMetaData

`func (o *RefundTransaction) GetMetaData() MetaData1`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *RefundTransaction) GetMetaDataOk() (*MetaData1, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *RefundTransaction) SetMetaData(v MetaData1)`

SetMetaData sets MetaData field to given value.


### GetGroupIds

`func (o *RefundTransaction) GetGroupIds() string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *RefundTransaction) GetGroupIdsOk() (*string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *RefundTransaction) SetGroupIds(v string)`

SetGroupIds sets GroupIds field to given value.


### SetGroupIdsNil

`func (o *RefundTransaction) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *RefundTransaction) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


