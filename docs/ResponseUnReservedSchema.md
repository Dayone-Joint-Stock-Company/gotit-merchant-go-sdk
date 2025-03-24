# ResponseUnReservedSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ReturnCode** | Pointer to **NullableString** | Result code if failed. In case of successful request: value is null | [optional] 
**MessageEn** | Pointer to **string** | Message notification in English | [optional] 
**MessageVi** | Pointer to **string** | Message notification in Vietnamese | [optional] 
**UsedStore** | Pointer to [**ResponseReservedSchemaUsedStore**](ResponseReservedSchemaUsedStore.md) |  | [optional] 
**BillNumber** | Pointer to **string** | Bill number | [optional] 
**Data** | Pointer to [**[]ResponseUnReservedSchemaDataInner**](ResponseUnReservedSchemaDataInner.md) | Detail items of voucher, if result is failed, response will return the first voucher code which is invalid | [optional] 

## Methods

### NewResponseUnReservedSchema

`func NewResponseUnReservedSchema() *ResponseUnReservedSchema`

NewResponseUnReservedSchema instantiates a new ResponseUnReservedSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseUnReservedSchemaWithDefaults

`func NewResponseUnReservedSchemaWithDefaults() *ResponseUnReservedSchema`

NewResponseUnReservedSchemaWithDefaults instantiates a new ResponseUnReservedSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResponseUnReservedSchema) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseUnReservedSchema) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseUnReservedSchema) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseUnReservedSchema) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetReturnCode

`func (o *ResponseUnReservedSchema) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ResponseUnReservedSchema) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ResponseUnReservedSchema) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ResponseUnReservedSchema) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ResponseUnReservedSchema) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ResponseUnReservedSchema) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetMessageEn

`func (o *ResponseUnReservedSchema) GetMessageEn() string`

GetMessageEn returns the MessageEn field if non-nil, zero value otherwise.

### GetMessageEnOk

`func (o *ResponseUnReservedSchema) GetMessageEnOk() (*string, bool)`

GetMessageEnOk returns a tuple with the MessageEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEn

`func (o *ResponseUnReservedSchema) SetMessageEn(v string)`

SetMessageEn sets MessageEn field to given value.

### HasMessageEn

`func (o *ResponseUnReservedSchema) HasMessageEn() bool`

HasMessageEn returns a boolean if a field has been set.

### GetMessageVi

`func (o *ResponseUnReservedSchema) GetMessageVi() string`

GetMessageVi returns the MessageVi field if non-nil, zero value otherwise.

### GetMessageViOk

`func (o *ResponseUnReservedSchema) GetMessageViOk() (*string, bool)`

GetMessageViOk returns a tuple with the MessageVi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVi

`func (o *ResponseUnReservedSchema) SetMessageVi(v string)`

SetMessageVi sets MessageVi field to given value.

### HasMessageVi

`func (o *ResponseUnReservedSchema) HasMessageVi() bool`

HasMessageVi returns a boolean if a field has been set.

### GetUsedStore

`func (o *ResponseUnReservedSchema) GetUsedStore() ResponseReservedSchemaUsedStore`

GetUsedStore returns the UsedStore field if non-nil, zero value otherwise.

### GetUsedStoreOk

`func (o *ResponseUnReservedSchema) GetUsedStoreOk() (*ResponseReservedSchemaUsedStore, bool)`

GetUsedStoreOk returns a tuple with the UsedStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStore

`func (o *ResponseUnReservedSchema) SetUsedStore(v ResponseReservedSchemaUsedStore)`

SetUsedStore sets UsedStore field to given value.

### HasUsedStore

`func (o *ResponseUnReservedSchema) HasUsedStore() bool`

HasUsedStore returns a boolean if a field has been set.

### GetBillNumber

`func (o *ResponseUnReservedSchema) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *ResponseUnReservedSchema) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *ResponseUnReservedSchema) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *ResponseUnReservedSchema) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.

### GetData

`func (o *ResponseUnReservedSchema) GetData() []ResponseUnReservedSchemaDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseUnReservedSchema) GetDataOk() (*[]ResponseUnReservedSchemaDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseUnReservedSchema) SetData(v []ResponseUnReservedSchemaDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseUnReservedSchema) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


