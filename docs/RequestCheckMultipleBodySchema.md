# RequestCheckMultipleBodySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | Pointer to **string** | Store pin | [optional] 
**Codes** | Pointer to **[]string** | Array of 10-16 characters Got It voucher codes | [optional] 
**BillNumber** | Pointer to **string** | Bill number will apply vouchers | [optional] 
**TotalBill** | Pointer to **int64** | Total bill amount | [optional] 
**SkipReservedWhenMarkUsed** | Pointer to **bool** | When true the system will execute the flow without reserve | [optional] 
**SkusInfo** | Pointer to [**[]RequestCheckMultipleBodySchemaSkusInfoInner**](RequestCheckMultipleBodySchemaSkusInfoInner.md) | SKU information in bill_number | [optional] 

## Methods

### NewRequestCheckMultipleBodySchema

`func NewRequestCheckMultipleBodySchema() *RequestCheckMultipleBodySchema`

NewRequestCheckMultipleBodySchema instantiates a new RequestCheckMultipleBodySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCheckMultipleBodySchemaWithDefaults

`func NewRequestCheckMultipleBodySchemaWithDefaults() *RequestCheckMultipleBodySchema`

NewRequestCheckMultipleBodySchemaWithDefaults instantiates a new RequestCheckMultipleBodySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *RequestCheckMultipleBodySchema) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *RequestCheckMultipleBodySchema) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *RequestCheckMultipleBodySchema) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *RequestCheckMultipleBodySchema) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCodes

`func (o *RequestCheckMultipleBodySchema) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *RequestCheckMultipleBodySchema) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *RequestCheckMultipleBodySchema) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *RequestCheckMultipleBodySchema) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetBillNumber

`func (o *RequestCheckMultipleBodySchema) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *RequestCheckMultipleBodySchema) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *RequestCheckMultipleBodySchema) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *RequestCheckMultipleBodySchema) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.

### GetTotalBill

`func (o *RequestCheckMultipleBodySchema) GetTotalBill() int64`

GetTotalBill returns the TotalBill field if non-nil, zero value otherwise.

### GetTotalBillOk

`func (o *RequestCheckMultipleBodySchema) GetTotalBillOk() (*int64, bool)`

GetTotalBillOk returns a tuple with the TotalBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBill

`func (o *RequestCheckMultipleBodySchema) SetTotalBill(v int64)`

SetTotalBill sets TotalBill field to given value.

### HasTotalBill

`func (o *RequestCheckMultipleBodySchema) HasTotalBill() bool`

HasTotalBill returns a boolean if a field has been set.

### GetSkipReservedWhenMarkUsed

`func (o *RequestCheckMultipleBodySchema) GetSkipReservedWhenMarkUsed() bool`

GetSkipReservedWhenMarkUsed returns the SkipReservedWhenMarkUsed field if non-nil, zero value otherwise.

### GetSkipReservedWhenMarkUsedOk

`func (o *RequestCheckMultipleBodySchema) GetSkipReservedWhenMarkUsedOk() (*bool, bool)`

GetSkipReservedWhenMarkUsedOk returns a tuple with the SkipReservedWhenMarkUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReservedWhenMarkUsed

`func (o *RequestCheckMultipleBodySchema) SetSkipReservedWhenMarkUsed(v bool)`

SetSkipReservedWhenMarkUsed sets SkipReservedWhenMarkUsed field to given value.

### HasSkipReservedWhenMarkUsed

`func (o *RequestCheckMultipleBodySchema) HasSkipReservedWhenMarkUsed() bool`

HasSkipReservedWhenMarkUsed returns a boolean if a field has been set.

### GetSkusInfo

`func (o *RequestCheckMultipleBodySchema) GetSkusInfo() []RequestCheckMultipleBodySchemaSkusInfoInner`

GetSkusInfo returns the SkusInfo field if non-nil, zero value otherwise.

### GetSkusInfoOk

`func (o *RequestCheckMultipleBodySchema) GetSkusInfoOk() (*[]RequestCheckMultipleBodySchemaSkusInfoInner, bool)`

GetSkusInfoOk returns a tuple with the SkusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusInfo

`func (o *RequestCheckMultipleBodySchema) SetSkusInfo(v []RequestCheckMultipleBodySchemaSkusInfoInner)`

SetSkusInfo sets SkusInfo field to given value.

### HasSkusInfo

`func (o *RequestCheckMultipleBodySchema) HasSkusInfo() bool`

HasSkusInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


