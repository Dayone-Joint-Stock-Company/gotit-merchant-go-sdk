# ResponseCheckMultipleSchemaDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Voucher code | [optional] 
**Value** | Pointer to **NullableInt64** | Value of voucher | [optional] 
**ProductId** | Pointer to **NullableInt64** | Product ID | [optional] 
**State** | Pointer to **NullableInt64** | State of voucher | [optional] 
**VoucherType** | Pointer to **string** | Voucher type, standard or conditional | [optional] 
**ExpiryDate** | Pointer to **string** | Expiry date of voucher (YYYY-MM-DD) | [optional] 
**CancelDate** | Pointer to **string** | Date cancel voucher (YYYY-MM-DD) | [optional] 
**Conditions** | Pointer to [**ResponseCheckMultipleSchemaDataInnerConditions**](ResponseCheckMultipleSchemaDataInnerConditions.md) |  | [optional] 
**Redemptions** | Pointer to [**ResponseCheckMultipleSchemaDataInnerRedemptions**](ResponseCheckMultipleSchemaDataInnerRedemptions.md) |  | [optional] 

## Methods

### NewResponseCheckMultipleSchemaDataInner

`func NewResponseCheckMultipleSchemaDataInner() *ResponseCheckMultipleSchemaDataInner`

NewResponseCheckMultipleSchemaDataInner instantiates a new ResponseCheckMultipleSchemaDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCheckMultipleSchemaDataInnerWithDefaults

`func NewResponseCheckMultipleSchemaDataInnerWithDefaults() *ResponseCheckMultipleSchemaDataInner`

NewResponseCheckMultipleSchemaDataInnerWithDefaults instantiates a new ResponseCheckMultipleSchemaDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResponseCheckMultipleSchemaDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResponseCheckMultipleSchemaDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResponseCheckMultipleSchemaDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetValue

`func (o *ResponseCheckMultipleSchemaDataInner) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseCheckMultipleSchemaDataInner) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponseCheckMultipleSchemaDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ResponseCheckMultipleSchemaDataInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ResponseCheckMultipleSchemaDataInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetProductId

`func (o *ResponseCheckMultipleSchemaDataInner) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ResponseCheckMultipleSchemaDataInner) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ResponseCheckMultipleSchemaDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ResponseCheckMultipleSchemaDataInner) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ResponseCheckMultipleSchemaDataInner) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetState

`func (o *ResponseCheckMultipleSchemaDataInner) GetState() int64`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetStateOk() (*int64, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseCheckMultipleSchemaDataInner) SetState(v int64)`

SetState sets State field to given value.

### HasState

`func (o *ResponseCheckMultipleSchemaDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ResponseCheckMultipleSchemaDataInner) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ResponseCheckMultipleSchemaDataInner) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetVoucherType

`func (o *ResponseCheckMultipleSchemaDataInner) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *ResponseCheckMultipleSchemaDataInner) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.

### HasVoucherType

`func (o *ResponseCheckMultipleSchemaDataInner) HasVoucherType() bool`

HasVoucherType returns a boolean if a field has been set.

### GetExpiryDate

`func (o *ResponseCheckMultipleSchemaDataInner) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ResponseCheckMultipleSchemaDataInner) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ResponseCheckMultipleSchemaDataInner) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetCancelDate

`func (o *ResponseCheckMultipleSchemaDataInner) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *ResponseCheckMultipleSchemaDataInner) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.

### HasCancelDate

`func (o *ResponseCheckMultipleSchemaDataInner) HasCancelDate() bool`

HasCancelDate returns a boolean if a field has been set.

### GetConditions

`func (o *ResponseCheckMultipleSchemaDataInner) GetConditions() ResponseCheckMultipleSchemaDataInnerConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetConditionsOk() (*ResponseCheckMultipleSchemaDataInnerConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResponseCheckMultipleSchemaDataInner) SetConditions(v ResponseCheckMultipleSchemaDataInnerConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ResponseCheckMultipleSchemaDataInner) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetRedemptions

`func (o *ResponseCheckMultipleSchemaDataInner) GetRedemptions() ResponseCheckMultipleSchemaDataInnerRedemptions`

GetRedemptions returns the Redemptions field if non-nil, zero value otherwise.

### GetRedemptionsOk

`func (o *ResponseCheckMultipleSchemaDataInner) GetRedemptionsOk() (*ResponseCheckMultipleSchemaDataInnerRedemptions, bool)`

GetRedemptionsOk returns a tuple with the Redemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptions

`func (o *ResponseCheckMultipleSchemaDataInner) SetRedemptions(v ResponseCheckMultipleSchemaDataInnerRedemptions)`

SetRedemptions sets Redemptions field to given value.

### HasRedemptions

`func (o *ResponseCheckMultipleSchemaDataInner) HasRedemptions() bool`

HasRedemptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


