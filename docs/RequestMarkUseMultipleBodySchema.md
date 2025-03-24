# RequestMarkUseMultipleBodySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | Pointer to **string** | Store pin | [optional] 
**Codes** | Pointer to **[]string** | Array of 10-16 characters Got It voucher codes | [optional] 
**BillNumber** | Pointer to **string** | Bill number will apply vouchers | [optional] 
**TotalBill** | Pointer to **int32** | Total bill amount | [optional] 
**SkipReservedWhenMarkUsed** | Pointer to **bool** | When true the system will execute the flow without reserve | [optional] 
**SkusInfo** | Pointer to [**[]RequestCheckMultipleBodySchemaSkusInfoInner**](RequestCheckMultipleBodySchemaSkusInfoInner.md) | SKU information in bill_number | [optional] 

## Methods

### NewRequestMarkUseMultipleBodySchema

`func NewRequestMarkUseMultipleBodySchema() *RequestMarkUseMultipleBodySchema`

NewRequestMarkUseMultipleBodySchema instantiates a new RequestMarkUseMultipleBodySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMarkUseMultipleBodySchemaWithDefaults

`func NewRequestMarkUseMultipleBodySchemaWithDefaults() *RequestMarkUseMultipleBodySchema`

NewRequestMarkUseMultipleBodySchemaWithDefaults instantiates a new RequestMarkUseMultipleBodySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *RequestMarkUseMultipleBodySchema) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *RequestMarkUseMultipleBodySchema) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *RequestMarkUseMultipleBodySchema) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *RequestMarkUseMultipleBodySchema) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCodes

`func (o *RequestMarkUseMultipleBodySchema) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *RequestMarkUseMultipleBodySchema) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *RequestMarkUseMultipleBodySchema) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *RequestMarkUseMultipleBodySchema) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetBillNumber

`func (o *RequestMarkUseMultipleBodySchema) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *RequestMarkUseMultipleBodySchema) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *RequestMarkUseMultipleBodySchema) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *RequestMarkUseMultipleBodySchema) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.

### GetTotalBill

`func (o *RequestMarkUseMultipleBodySchema) GetTotalBill() int32`

GetTotalBill returns the TotalBill field if non-nil, zero value otherwise.

### GetTotalBillOk

`func (o *RequestMarkUseMultipleBodySchema) GetTotalBillOk() (*int32, bool)`

GetTotalBillOk returns a tuple with the TotalBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBill

`func (o *RequestMarkUseMultipleBodySchema) SetTotalBill(v int32)`

SetTotalBill sets TotalBill field to given value.

### HasTotalBill

`func (o *RequestMarkUseMultipleBodySchema) HasTotalBill() bool`

HasTotalBill returns a boolean if a field has been set.

### GetSkipReservedWhenMarkUsed

`func (o *RequestMarkUseMultipleBodySchema) GetSkipReservedWhenMarkUsed() bool`

GetSkipReservedWhenMarkUsed returns the SkipReservedWhenMarkUsed field if non-nil, zero value otherwise.

### GetSkipReservedWhenMarkUsedOk

`func (o *RequestMarkUseMultipleBodySchema) GetSkipReservedWhenMarkUsedOk() (*bool, bool)`

GetSkipReservedWhenMarkUsedOk returns a tuple with the SkipReservedWhenMarkUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReservedWhenMarkUsed

`func (o *RequestMarkUseMultipleBodySchema) SetSkipReservedWhenMarkUsed(v bool)`

SetSkipReservedWhenMarkUsed sets SkipReservedWhenMarkUsed field to given value.

### HasSkipReservedWhenMarkUsed

`func (o *RequestMarkUseMultipleBodySchema) HasSkipReservedWhenMarkUsed() bool`

HasSkipReservedWhenMarkUsed returns a boolean if a field has been set.

### GetSkusInfo

`func (o *RequestMarkUseMultipleBodySchema) GetSkusInfo() []RequestCheckMultipleBodySchemaSkusInfoInner`

GetSkusInfo returns the SkusInfo field if non-nil, zero value otherwise.

### GetSkusInfoOk

`func (o *RequestMarkUseMultipleBodySchema) GetSkusInfoOk() (*[]RequestCheckMultipleBodySchemaSkusInfoInner, bool)`

GetSkusInfoOk returns a tuple with the SkusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkusInfo

`func (o *RequestMarkUseMultipleBodySchema) SetSkusInfo(v []RequestCheckMultipleBodySchemaSkusInfoInner)`

SetSkusInfo sets SkusInfo field to given value.

### HasSkusInfo

`func (o *RequestMarkUseMultipleBodySchema) HasSkusInfo() bool`

HasSkusInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


