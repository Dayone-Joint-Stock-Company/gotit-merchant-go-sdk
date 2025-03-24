# RequestUnReservedBodySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | Pointer to **string** | Store pin | [optional] 
**Codes** | Pointer to **[]string** | Array of 10-16 characters Got It voucher codes | [optional] 
**BillNumber** | Pointer to **string** | Bill number will apply vouchers | [optional] 
**BillCreatedAt** | Pointer to **string** | Bill creation time. Format: YYYY-MM-DD HH:MM:SS | [optional] 

## Methods

### NewRequestUnReservedBodySchema

`func NewRequestUnReservedBodySchema() *RequestUnReservedBodySchema`

NewRequestUnReservedBodySchema instantiates a new RequestUnReservedBodySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUnReservedBodySchemaWithDefaults

`func NewRequestUnReservedBodySchemaWithDefaults() *RequestUnReservedBodySchema`

NewRequestUnReservedBodySchemaWithDefaults instantiates a new RequestUnReservedBodySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *RequestUnReservedBodySchema) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *RequestUnReservedBodySchema) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *RequestUnReservedBodySchema) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *RequestUnReservedBodySchema) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetCodes

`func (o *RequestUnReservedBodySchema) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *RequestUnReservedBodySchema) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *RequestUnReservedBodySchema) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *RequestUnReservedBodySchema) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetBillNumber

`func (o *RequestUnReservedBodySchema) GetBillNumber() string`

GetBillNumber returns the BillNumber field if non-nil, zero value otherwise.

### GetBillNumberOk

`func (o *RequestUnReservedBodySchema) GetBillNumberOk() (*string, bool)`

GetBillNumberOk returns a tuple with the BillNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillNumber

`func (o *RequestUnReservedBodySchema) SetBillNumber(v string)`

SetBillNumber sets BillNumber field to given value.

### HasBillNumber

`func (o *RequestUnReservedBodySchema) HasBillNumber() bool`

HasBillNumber returns a boolean if a field has been set.

### GetBillCreatedAt

`func (o *RequestUnReservedBodySchema) GetBillCreatedAt() string`

GetBillCreatedAt returns the BillCreatedAt field if non-nil, zero value otherwise.

### GetBillCreatedAtOk

`func (o *RequestUnReservedBodySchema) GetBillCreatedAtOk() (*string, bool)`

GetBillCreatedAtOk returns a tuple with the BillCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCreatedAt

`func (o *RequestUnReservedBodySchema) SetBillCreatedAt(v string)`

SetBillCreatedAt sets BillCreatedAt field to given value.

### HasBillCreatedAt

`func (o *RequestUnReservedBodySchema) HasBillCreatedAt() bool`

HasBillCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


