# ResponseMarkUseMultipleSchemaDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Voucher code | [optional] 
**Value** | Pointer to **NullableInt32** | Value of voucher | [optional] 
**State** | Pointer to **NullableInt32** | State of voucher | [optional] 
**VoucherType** | Pointer to **string** | Voucher type, standard or conditional | [optional] 
**Conditions** | Pointer to [**ResponseMarkUseMultipleSchemaDataInnerConditions**](ResponseMarkUseMultipleSchemaDataInnerConditions.md) |  | [optional] 
**Redemptions** | Pointer to [**ResponseMarkUseMultipleSchemaDataInnerRedemptions**](ResponseMarkUseMultipleSchemaDataInnerRedemptions.md) |  | [optional] 

## Methods

### NewResponseMarkUseMultipleSchemaDataInner

`func NewResponseMarkUseMultipleSchemaDataInner() *ResponseMarkUseMultipleSchemaDataInner`

NewResponseMarkUseMultipleSchemaDataInner instantiates a new ResponseMarkUseMultipleSchemaDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMarkUseMultipleSchemaDataInnerWithDefaults

`func NewResponseMarkUseMultipleSchemaDataInnerWithDefaults() *ResponseMarkUseMultipleSchemaDataInner`

NewResponseMarkUseMultipleSchemaDataInnerWithDefaults instantiates a new ResponseMarkUseMultipleSchemaDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResponseMarkUseMultipleSchemaDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetValue

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponseMarkUseMultipleSchemaDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ResponseMarkUseMultipleSchemaDataInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetState

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *ResponseMarkUseMultipleSchemaDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ResponseMarkUseMultipleSchemaDataInner) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetVoucherType

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.

### HasVoucherType

`func (o *ResponseMarkUseMultipleSchemaDataInner) HasVoucherType() bool`

HasVoucherType returns a boolean if a field has been set.

### GetConditions

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetConditions() ResponseMarkUseMultipleSchemaDataInnerConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetConditionsOk() (*ResponseMarkUseMultipleSchemaDataInnerConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetConditions(v ResponseMarkUseMultipleSchemaDataInnerConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ResponseMarkUseMultipleSchemaDataInner) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetRedemptions

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetRedemptions() ResponseMarkUseMultipleSchemaDataInnerRedemptions`

GetRedemptions returns the Redemptions field if non-nil, zero value otherwise.

### GetRedemptionsOk

`func (o *ResponseMarkUseMultipleSchemaDataInner) GetRedemptionsOk() (*ResponseMarkUseMultipleSchemaDataInnerRedemptions, bool)`

GetRedemptionsOk returns a tuple with the Redemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptions

`func (o *ResponseMarkUseMultipleSchemaDataInner) SetRedemptions(v ResponseMarkUseMultipleSchemaDataInnerRedemptions)`

SetRedemptions sets Redemptions field to given value.

### HasRedemptions

`func (o *ResponseMarkUseMultipleSchemaDataInner) HasRedemptions() bool`

HasRedemptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


