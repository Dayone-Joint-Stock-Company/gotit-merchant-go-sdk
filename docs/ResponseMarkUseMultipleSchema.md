# ResponseMarkUseMultipleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ReturnCode** | Pointer to **NullableString** | Result code if failed. In case of successful request: value is null | [optional] 
**MessageEn** | Pointer to **string** | Message notification in English | [optional] 
**MessageVi** | Pointer to **string** | Message notification in Vietnamese | [optional] 
**Data** | Pointer to [**[]ResponseMarkUseMultipleSchemaDataInner**](ResponseMarkUseMultipleSchemaDataInner.md) | Detail items of voucher, if result is failed, response will return the first voucher code which is invalid | [optional] 
**TransactionId** | Pointer to **string** | Transaction ID (if mark used successfully) | [optional] 
**BillNumber** | Pointer to **string** | Bill number that vouchers were marked as used for. | [optional] 

## Methods

### NewResponseMarkUseMultipleSchema

`func NewResponseMarkUseMultipleSchema() *ResponseMarkUseMultipleSchema`

NewResponseMarkUseMultipleSchema instantiates a new ResponseMarkUseMultipleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMarkUseMultipleSchemaWithDefaults

`func NewResponseMarkUseMultipleSchemaWithDefaults() *ResponseMarkUseMultipleSchema`

NewResponseMarkUseMultipleSchemaWithDefaults instantiates a new ResponseMarkUseMultipleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResponseMarkUseMultipleSchema) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseMarkUseMultipleSchema) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseMarkUseMultipleSchema) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseMarkUseMultipleSchema) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetReturnCode

`func (o *ResponseMarkUseMultipleSchema) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ResponseMarkUseMultipleSchema) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ResponseMarkUseMultipleSchema) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ResponseMarkUseMultipleSchema) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ResponseMarkUseMultipleSchema) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ResponseMarkUseMultipleSchema) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetMessageEn

`func (o *ResponseMarkUseMultipleSchema) GetMessageEn() string`

GetMessageEn returns the MessageEn field if non-nil, zero value otherwise.

### GetMessageEnOk

`func (o *ResponseMarkUseMultipleSchema) GetMessageEnOk() (*string, bool)`

GetMessageEnOk returns a tuple with the MessageEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageEn

`func (o *ResponseMarkUseMultipleSchema) SetMessageEn(v string)`

SetMessageEn sets MessageEn field to given value.

### HasMessageEn

`func (o *ResponseMarkUseMultipleSchema) HasMessageEn() bool`

HasMessageEn returns a boolean if a field has been set.

### GetMessageVi

`func (o *ResponseMarkUseMultipleSchema) GetMessageVi() string`

GetMessageVi returns the MessageVi field if non-nil, zero value otherwise.

### GetMessageViOk

`func (o *ResponseMarkUseMultipleSchema) GetMessageViOk() (*string, bool)`

GetMessageViOk returns a tuple with the MessageVi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVi

`func (o *ResponseMarkUseMultipleSchema) SetMessageVi(v string)`

SetMessageVi sets MessageVi field to given value.

### HasMessageVi

`func (o *ResponseMarkUseMultipleSchema) HasMessageVi() bool`

HasMessageVi returns a boolean if a field has been set.

### GetData

`func (o *ResponseMarkUseMultipleSchema) GetData() []ResponseMarkUseMultipleSchemaDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseMarkUseMultipleSchema) GetDataOk() (*[]ResponseMarkUseMultipleSchemaDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseMarkUseMultipleSchema) SetData(v []ResponseMarkUseMultipleSchemaDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseMarkUseMultipleSchema) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseMarkUseMultipleSchema) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseMarkUseMultipleSchema) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseMarkUseMultipleSchema) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseMarkUseMultipleSchema) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetBillNumber

`func (o *ResponseMarkUseMultipleSchema) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *ResponseMarkUseMultipleSchema) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *ResponseMarkUseMultipleSchema) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *ResponseMarkUseMultipleSchema) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


