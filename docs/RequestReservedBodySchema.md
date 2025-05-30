# RequestReservedBodySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | Pointer to **string** | Store pin | [optional] 
**Codes** | Pointer to **[]string** | Array of 10-16 characters Got It voucher codes | [optional] 
**BillNumber** | Pointer to **string** | Bill number will apply vouchers | [optional] 
**TotalBill** | Pointer to **int64** | Total bill amount | [optional] 
**BillCreatedAt** | Pointer to **string** | Bill creation time. Format: YYYY-MM-DD HH:MM:SS | [optional] 
**SkusInfo** | Pointer to [**[]RequestCheckMultipleBodySchemaSkusInfoInner**](RequestCheckMultipleBodySchemaSkusInfoInner.md) | SKU information in bill_number | [optional] 

## Methods

### NewRequestReservedBodySchema

`func NewRequestReservedBodySchema() *RequestReservedBodySchema`

NewRequestReservedBodySchema instantiates a new RequestReservedBodySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestReservedBodySchemaWithDefaults

`func NewRequestReservedBodySchemaWithDefaults() *RequestReservedBodySchema`

NewRequestReservedBodySchemaWithDefaults instantiates a new RequestReservedBodySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *RequestReservedBodySchema) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *RequestReservedBodySchema) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *RequestReservedBodySchema) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *RequestReservedBodySchema) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCodes

`func (o *RequestReservedBodySchema) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *RequestReservedBodySchema) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *RequestReservedBodySchema) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *RequestReservedBodySchema) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetBillNumber

`func (o *RequestReservedBodySchema) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *RequestReservedBodySchema) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *RequestReservedBodySchema) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *RequestReservedBodySchema) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.

### GetTotalBill

`func (o *RequestReservedBodySchema) GetTotalBill() int64`

GetTotalBill returns the TotalBill field if non-nil, zero value otherwise.

### GetTotalBillOk

`func (o *RequestReservedBodySchema) GetTotalBillOk() (*int64, bool)`

GetTotalBillOk returns a tuple with the TotalBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBill

`func (o *RequestReservedBodySchema) SetTotalBill(v int64)`

SetTotalBill sets TotalBill field to given value.

### HasTotalBill

`func (o *RequestReservedBodySchema) HasTotalBill() bool`

HasTotalBill returns a boolean if a field has been set.

### GetBillCreatedAt

`func (o *RequestReservedBodySchema) GetBillCreatedAt() string`

GetBillCreatedAt returns the BillCreatedAt field if non-nil, zero value otherwise.

### GetBillCreatedAtOk

`func (o *RequestReservedBodySchema) GetBillCreatedAtOk() (*string, bool)`

GetBillCreatedAtOk returns a tuple with the BillCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCreatedAt

`func (o *RequestReservedBodySchema) SetBillCreatedAt(v string)`

SetBillCreatedAt sets BillCreatedAt field to given value.

### HasBillCreatedAt

`func (o *RequestReservedBodySchema) HasBillCreatedAt() bool`

HasBillCreatedAt returns a boolean if a field has been set.

### GetSkusInfo

`func (o *RequestReservedBodySchema) GetSkusInfo() []RequestCheckMultipleBodySchemaSkusInfoInner`

GetSkusInfo returns the SkusInfo field if non-nil, zero value otherwise.

### GetSkusInfoOk

`func (o *RequestReservedBodySchema) GetSkusInfoOk() (*[]RequestCheckMultipleBodySchemaSkusInfoInner, bool)`

GetSkusInfoOk returns a tuple with the SkusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusInfo

`func (o *RequestReservedBodySchema) SetSkusInfo(v []RequestCheckMultipleBodySchemaSkusInfoInner)`

SetSkusInfo sets SkusInfo field to given value.

### HasSkusInfo

`func (o *RequestReservedBodySchema) HasSkusInfo() bool`

HasSkusInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


