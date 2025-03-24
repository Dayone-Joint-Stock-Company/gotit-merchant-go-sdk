/*
Merchant APIs

Technical document APIs for Merchant APIs

API version: 6.0
Contact: duong.vu@gotit.vn
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gotit_merchant_apis

import (
	"encoding/json"
)

// checks if the RequestReservedBodySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestReservedBodySchema{}

// RequestReservedBodySchema struct for RequestReservedBodySchema
type RequestReservedBodySchema struct {
	// Store pin
	Pin *string `json:"pin,omitempty"`
	// Array of 10-16 characters Got It voucher codes
	Codes []string `json:"codes,omitempty"`
	// Bill number will apply vouchers
	BillNumber *string `json:"bill_number,omitempty"`
	// Total bill amount
	TotalBill *int32 `json:"total_bill,omitempty"`
	// Bill creation time. Format: YYYY-MM-DD HH:MM:SS
	BillCreatedAt *string `json:"bill_created_at,omitempty"`
	// SKU information in bill_number
	SkusInfo []RequestCheckMultipleBodySchemaSkusInfoInner `json:"skus_info,omitempty"`
}

// NewRequestReservedBodySchema instantiates a new RequestReservedBodySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestReservedBodySchema() *RequestReservedBodySchema {
	this := RequestReservedBodySchema{}
	return &this
}

// NewRequestReservedBodySchemaWithDefaults instantiates a new RequestReservedBodySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestReservedBodySchemaWithDefaults() *RequestReservedBodySchema {
	this := RequestReservedBodySchema{}
	return &this
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *RequestReservedBodySchema) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestReservedBodySchema) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *RequestReservedBodySchema) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *RequestReservedBodySchema) SetPin(v string) {
	o.Pin = &v
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *RequestReservedBodySchema) GetCodes() []string {
	if o == nil || IsNil(o.Codes) {
		var ret []string
		return ret
	}
	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestReservedBodySchema) GetCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *RequestReservedBodySchema) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given []string and assigns it to the Codes field.
func (o *RequestReservedBodySchema) SetCodes(v []string) {
	o.Codes = v
}

// GetBillNumber returns the BillNumber field value if set, zero value otherwise.
func (o *RequestReservedBodySchema) GetBillNumber() string {
	if o == nil || IsNil(o.BillNumber) {
		var ret string
		return ret
	}
	return *o.BillNumber
}

// GetBillNumberOk returns a tuple with the BillNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestReservedBodySchema) GetBillNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BillNumber) {
		return nil, false
	}
	return o.BillNumber, true
}

// HasBillNumber returns a boolean if a field has been set.
func (o *RequestReservedBodySchema) HasBillNumber() bool {
	if o != nil && !IsNil(o.BillNumber) {
		return true
	}

	return false
}

// SetBillNumber gets a reference to the given string and assigns it to the BillNumber field.
func (o *RequestReservedBodySchema) SetBillNumber(v string) {
	o.BillNumber = &v
}

// GetTotalBill returns the TotalBill field value if set, zero value otherwise.
func (o *RequestReservedBodySchema) GetTotalBill() int32 {
	if o == nil || IsNil(o.TotalBill) {
		var ret int32
		return ret
	}
	return *o.TotalBill
}

// GetTotalBillOk returns a tuple with the TotalBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestReservedBodySchema) GetTotalBillOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalBill) {
		return nil, false
	}
	return o.TotalBill, true
}

// HasTotalBill returns a boolean if a field has been set.
func (o *RequestReservedBodySchema) HasTotalBill() bool {
	if o != nil && !IsNil(o.TotalBill) {
		return true
	}

	return false
}

// SetTotalBill gets a reference to the given int32 and assigns it to the TotalBill field.
func (o *RequestReservedBodySchema) SetTotalBill(v int32) {
	o.TotalBill = &v
}

// GetBillCreatedAt returns the BillCreatedAt field value if set, zero value otherwise.
func (o *RequestReservedBodySchema) GetBillCreatedAt() string {
	if o == nil || IsNil(o.BillCreatedAt) {
		var ret string
		return ret
	}
	return *o.BillCreatedAt
}

// GetBillCreatedAtOk returns a tuple with the BillCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestReservedBodySchema) GetBillCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.BillCreatedAt) {
		return nil, false
	}
	return o.BillCreatedAt, true
}

// HasBillCreatedAt returns a boolean if a field has been set.
func (o *RequestReservedBodySchema) HasBillCreatedAt() bool {
	if o != nil && !IsNil(o.BillCreatedAt) {
		return true
	}

	return false
}

// SetBillCreatedAt gets a reference to the given string and assigns it to the BillCreatedAt field.
func (o *RequestReservedBodySchema) SetBillCreatedAt(v string) {
	o.BillCreatedAt = &v
}

// GetSkusInfo returns the SkusInfo field value if set, zero value otherwise.
func (o *RequestReservedBodySchema) GetSkusInfo() []RequestCheckMultipleBodySchemaSkusInfoInner {
	if o == nil || IsNil(o.SkusInfo) {
		var ret []RequestCheckMultipleBodySchemaSkusInfoInner
		return ret
	}
	return o.SkusInfo
}

// GetSkusInfoOk returns a tuple with the SkusInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestReservedBodySchema) GetSkusInfoOk() ([]RequestCheckMultipleBodySchemaSkusInfoInner, bool) {
	if o == nil || IsNil(o.SkusInfo) {
		return nil, false
	}
	return o.SkusInfo, true
}

// HasSkusInfo returns a boolean if a field has been set.
func (o *RequestReservedBodySchema) HasSkusInfo() bool {
	if o != nil && !IsNil(o.SkusInfo) {
		return true
	}

	return false
}

// SetSkusInfo gets a reference to the given []RequestCheckMultipleBodySchemaSkusInfoInner and assigns it to the SkusInfo field.
func (o *RequestReservedBodySchema) SetSkusInfo(v []RequestCheckMultipleBodySchemaSkusInfoInner) {
	o.SkusInfo = v
}

func (o RequestReservedBodySchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestReservedBodySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.Codes) {
		toSerialize["codes"] = o.Codes
	}
	if !IsNil(o.BillNumber) {
		toSerialize["bill_number"] = o.BillNumber
	}
	if !IsNil(o.TotalBill) {
		toSerialize["total_bill"] = o.TotalBill
	}
	if !IsNil(o.BillCreatedAt) {
		toSerialize["bill_created_at"] = o.BillCreatedAt
	}
	if !IsNil(o.SkusInfo) {
		toSerialize["skus_info"] = o.SkusInfo
	}
	return toSerialize, nil
}

type NullableRequestReservedBodySchema struct {
	value *RequestReservedBodySchema
	isSet bool
}

func (v NullableRequestReservedBodySchema) Get() *RequestReservedBodySchema {
	return v.value
}

func (v *NullableRequestReservedBodySchema) Set(val *RequestReservedBodySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestReservedBodySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestReservedBodySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestReservedBodySchema(val *RequestReservedBodySchema) *NullableRequestReservedBodySchema {
	return &NullableRequestReservedBodySchema{value: val, isSet: true}
}

func (v NullableRequestReservedBodySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestReservedBodySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


