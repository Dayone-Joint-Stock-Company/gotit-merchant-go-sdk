# ResponseCheckMultipleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ReturnCode** | Pointer to **NullableString** | Result code if failed. In case of successful request: value is null | [optional] 
**MessageEn** | Pointer to **string** | Message notification in English | [optional] 
**MessageVi** | Pointer to **string** | Message notification in Vietnamese | [optional] 
**Data** | Pointer to [**[]ResponseCheckMultipleSchemaDataInner**](ResponseCheckMultipleSchemaDataInner.md) | Detail items of voucher, if result is failed, response will return the first voucher code which is invalid | [optional] 

## Methods

### NewResponseCheckMultipleSchema

`func NewResponseCheckMultipleSchema() *ResponseCheckMultipleSchema`

NewResponseCheckMultipleSchema instantiates a new ResponseCheckMultipleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCheckMultipleSchemaWithDefaults

`func NewResponseCheckMultipleSchemaWithDefaults() *ResponseCheckMultipleSchema`

NewResponseCheckMultipleSchemaWithDefaults instantiates a new ResponseCheckMultipleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResponseCheckMultipleSchema) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseCheckMultipleSchema) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseCheckMultipleSchema) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseCheckMultipleSchema) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetReturnCode

`func (o *ResponseCheckMultipleSchema) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ResponseCheckMultipleSchema) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ResponseCheckMultipleSchema) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ResponseCheckMultipleSchema) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ResponseCheckMultipleSchema) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ResponseCheckMultipleSchema) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetMessageEn

`func (o *ResponseCheckMultipleSchema) GetMessageEn() string`

GetMessageEn returns the MessageEn field if non-nil, zero value otherwise.

### GetMessageEnOk

`func (o *ResponseCheckMultipleSchema) GetMessageEnOk() (*string, bool)`

GetMessageEnOk returns a tuple with the MessageEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEn

`func (o *ResponseCheckMultipleSchema) SetMessageEn(v string)`

SetMessageEn sets MessageEn field to given value.

### HasMessageEn

`func (o *ResponseCheckMultipleSchema) HasMessageEn() bool`

HasMessageEn returns a boolean if a field has been set.

### GetMessageVi

`func (o *ResponseCheckMultipleSchema) GetMessageVi() string`

GetMessageVi returns the MessageVi field if non-nil, zero value otherwise.

### GetMessageViOk

`func (o *ResponseCheckMultipleSchema) GetMessageViOk() (*string, bool)`

GetMessageViOk returns a tuple with the MessageVi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVi

`func (o *ResponseCheckMultipleSchema) SetMessageVi(v string)`

SetMessageVi sets MessageVi field to given value.

### HasMessageVi

`func (o *ResponseCheckMultipleSchema) HasMessageVi() bool`

HasMessageVi returns a boolean if a field has been set.

### GetData

`func (o *ResponseCheckMultipleSchema) GetData() []ResponseCheckMultipleSchemaDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseCheckMultipleSchema) GetDataOk() (*[]ResponseCheckMultipleSchemaDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseCheckMultipleSchema) SetData(v []ResponseCheckMultipleSchemaDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseCheckMultipleSchema) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


