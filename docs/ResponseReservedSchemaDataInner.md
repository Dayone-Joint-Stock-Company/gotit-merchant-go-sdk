# ResponseReservedSchemaDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Voucher code | [optional] 
**Value** | Pointer to **NullableInt32** | Value of voucher | [optional] 
**State** | Pointer to **NullableInt32** | State of voucher | [optional] 
**ProductId** | Pointer to **NullableInt32** | Product ID | [optional] 
**VoucherType** | Pointer to **string** | Voucher type, standard or conditional | [optional] 
**Conditions** | Pointer to [**ResponseMarkUseMultipleSchemaDataInnerConditions**](ResponseMarkUseMultipleSchemaDataInnerConditions.md) |  | [optional] 
**Redemptions** | Pointer to [**ResponseReservedSchemaDataInnerRedemptions**](ResponseReservedSchemaDataInnerRedemptions.md) |  | [optional] 

## Methods

### NewResponseReservedSchemaDataInner

`func NewResponseReservedSchemaDataInner() *ResponseReservedSchemaDataInner`

NewResponseReservedSchemaDataInner instantiates a new ResponseReservedSchemaDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseReservedSchemaDataInnerWithDefaults

`func NewResponseReservedSchemaDataInnerWithDefaults() *ResponseReservedSchemaDataInner`

NewResponseReservedSchemaDataInnerWithDefaults instantiates a new ResponseReservedSchemaDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResponseReservedSchemaDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResponseReservedSchemaDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResponseReservedSchemaDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResponseReservedSchemaDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetValue

`func (o *ResponseReservedSchemaDataInner) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseReservedSchemaDataInner) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseReservedSchemaDataInner) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponseReservedSchemaDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ResponseReservedSchemaDataInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ResponseReservedSchemaDataInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetState

`func (o *ResponseReservedSchemaDataInner) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseReservedSchemaDataInner) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseReservedSchemaDataInner) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *ResponseReservedSchemaDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ResponseReservedSchemaDataInner) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ResponseReservedSchemaDataInner) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetProductId

`func (o *ResponseReservedSchemaDataInner) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ResponseReservedSchemaDataInner) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ResponseReservedSchemaDataInner) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ResponseReservedSchemaDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ResponseReservedSchemaDataInner) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ResponseReservedSchemaDataInner) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetVoucherType

`func (o *ResponseReservedSchemaDataInner) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *ResponseReservedSchemaDataInner) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *ResponseReservedSchemaDataInner) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.

### HasVoucherType

`func (o *ResponseReservedSchemaDataInner) HasVoucherType() bool`

HasVoucherType returns a boolean if a field has been set.

### GetConditions

`func (o *ResponseReservedSchemaDataInner) GetConditions() ResponseMarkUseMultipleSchemaDataInnerConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResponseReservedSchemaDataInner) GetConditionsOk() (*ResponseMarkUseMultipleSchemaDataInnerConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResponseReservedSchemaDataInner) SetConditions(v ResponseMarkUseMultipleSchemaDataInnerConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ResponseReservedSchemaDataInner) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetRedemptions

`func (o *ResponseReservedSchemaDataInner) GetRedemptions() ResponseReservedSchemaDataInnerRedemptions`

GetRedemptions returns the Redemptions field if non-nil, zero value otherwise.

### GetRedemptionsOk

`func (o *ResponseReservedSchemaDataInner) GetRedemptionsOk() (*ResponseReservedSchemaDataInnerRedemptions, bool)`

GetRedemptionsOk returns a tuple with the Redemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptions

`func (o *ResponseReservedSchemaDataInner) SetRedemptions(v ResponseReservedSchemaDataInnerRedemptions)`

SetRedemptions sets Redemptions field to given value.

### HasRedemptions

`func (o *ResponseReservedSchemaDataInner) HasRedemptions() bool`

HasRedemptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


