# ResponseReservedSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ReturnCode** | Pointer to **NullableString** | Result code if failed. In case of successful request: value is null | [optional] 
**MessageEn** | Pointer to **string** | Message notification in English | [optional] 
**MessageVi** | Pointer to **string** | Message notification in Vietnamese | [optional] 
**UsedStore** | Pointer to [**ResponseReservedSchemaUsedStore**](ResponseReservedSchemaUsedStore.md) |  | [optional] 
**BillNumber** | Pointer to **string** | Bill number | [optional] 
**Data** | Pointer to [**[]ResponseReservedSchemaDataInner**](ResponseReservedSchemaDataInner.md) | Detail items of voucher, if result is failed, response will return the first voucher code which is invalid | [optional] 

## Methods

### NewResponseReservedSchema

`func NewResponseReservedSchema() *ResponseReservedSchema`

NewResponseReservedSchema instantiates a new ResponseReservedSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseReservedSchemaWithDefaults

`func NewResponseReservedSchemaWithDefaults() *ResponseReservedSchema`

NewResponseReservedSchemaWithDefaults instantiates a new ResponseReservedSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResponseReservedSchema) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseReservedSchema) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseReservedSchema) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseReservedSchema) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetReturnCode

`func (o *ResponseReservedSchema) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ResponseReservedSchema) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ResponseReservedSchema) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ResponseReservedSchema) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ResponseReservedSchema) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ResponseReservedSchema) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetMessageEn

`func (o *ResponseReservedSchema) GetMessageEn() string`

GetMessageEn returns the MessageEn field if non-nil, zero value otherwise.

### GetMessageEnOk

`func (o *ResponseReservedSchema) GetMessageEnOk() (*string, bool)`

GetMessageEnOk returns a tuple with the MessageEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEn

`func (o *ResponseReservedSchema) SetMessageEn(v string)`

SetMessageEn sets MessageEn field to given value.

### HasMessageEn

`func (o *ResponseReservedSchema) HasMessageEn() bool`

HasMessageEn returns a boolean if a field has been set.

### GetMessageVi

`func (o *ResponseReservedSchema) GetMessageVi() string`

GetMessageVi returns the MessageVi field if non-nil, zero value otherwise.

### GetMessageViOk

`func (o *ResponseReservedSchema) GetMessageViOk() (*string, bool)`

GetMessageViOk returns a tuple with the MessageVi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVi

`func (o *ResponseReservedSchema) SetMessageVi(v string)`

SetMessageVi sets MessageVi field to given value.

### HasMessageVi

`func (o *ResponseReservedSchema) HasMessageVi() bool`

HasMessageVi returns a boolean if a field has been set.

### GetUsedStore

`func (o *ResponseReservedSchema) GetUsedStore() ResponseReservedSchemaUsedStore`

GetUsedStore returns the UsedStore field if non-nil, zero value otherwise.

### GetUsedStoreOk

`func (o *ResponseReservedSchema) GetUsedStoreOk() (*ResponseReservedSchemaUsedStore, bool)`

GetUsedStoreOk returns a tuple with the UsedStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStore

`func (o *ResponseReservedSchema) SetUsedStore(v ResponseReservedSchemaUsedStore)`

SetUsedStore sets UsedStore field to given value.

### HasUsedStore

`func (o *ResponseReservedSchema) HasUsedStore() bool`

HasUsedStore returns a boolean if a field has been set.

### GetBillNumber

`func (o *ResponseReservedSchema) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *ResponseReservedSchema) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *ResponseReservedSchema) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *ResponseReservedSchema) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.

### GetData

`func (o *ResponseReservedSchema) GetData() []ResponseReservedSchemaDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseReservedSchema) GetDataOk() (*[]ResponseReservedSchemaDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseReservedSchema) SetData(v []ResponseReservedSchemaDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseReservedSchema) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


