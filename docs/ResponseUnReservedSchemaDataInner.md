# ResponseUnReservedSchemaDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Voucher code | [optional] 
**Value** | Pointer to **NullableInt32** | Value of voucher | [optional] 
**ProductId** | Pointer to **NullableInt32** | Product ID | [optional] 
**VoucherType** | Pointer to **string** | Voucher type, standard or conditional | [optional] 

## Methods

### NewResponseUnReservedSchemaDataInner

`func NewResponseUnReservedSchemaDataInner() *ResponseUnReservedSchemaDataInner`

NewResponseUnReservedSchemaDataInner instantiates a new ResponseUnReservedSchemaDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseUnReservedSchemaDataInnerWithDefaults

`func NewResponseUnReservedSchemaDataInnerWithDefaults() *ResponseUnReservedSchemaDataInner`

NewResponseUnReservedSchemaDataInnerWithDefaults instantiates a new ResponseUnReservedSchemaDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResponseUnReservedSchemaDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResponseUnReservedSchemaDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResponseUnReservedSchemaDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResponseUnReservedSchemaDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetValue

`func (o *ResponseUnReservedSchemaDataInner) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResponseUnReservedSchemaDataInner) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResponseUnReservedSchemaDataInner) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResponseUnReservedSchemaDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ResponseUnReservedSchemaDataInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ResponseUnReservedSchemaDataInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetProductId

`func (o *ResponseUnReservedSchemaDataInner) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ResponseUnReservedSchemaDataInner) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ResponseUnReservedSchemaDataInner) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ResponseUnReservedSchemaDataInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *ResponseUnReservedSchemaDataInner) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ResponseUnReservedSchemaDataInner) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetVoucherType

`func (o *ResponseUnReservedSchemaDataInner) GetVoucherType() string`

GetVoucherType returns the VoucherType field if non-nil, zero value otherwise.

### GetVoucherTypeOk

`func (o *ResponseUnReservedSchemaDataInner) GetVoucherTypeOk() (*string, bool)`

GetVoucherTypeOk returns a tuple with the VoucherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherType

`func (o *ResponseUnReservedSchemaDataInner) SetVoucherType(v string)`

SetVoucherType sets VoucherType field to given value.

### HasVoucherType

`func (o *ResponseUnReservedSchemaDataInner) HasVoucherType() bool`

HasVoucherType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


