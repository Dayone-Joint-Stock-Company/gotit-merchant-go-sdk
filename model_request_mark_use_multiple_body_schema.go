/*
Merchant APIs

Technical document APIs for Merchant APIs

API version: 6.0
Contact: duong.vu@gotit.vn
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gotit_merchant_api

import (
	"encoding/json"
)

// checks if the RequestMarkUseMultipleBodySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestMarkUseMultipleBodySchema{}

// RequestMarkUseMultipleBodySchema struct for RequestMarkUseMultipleBodySchema
type RequestMarkUseMultipleBodySchema struct {
	// Store pin
	Pin *string `json:"pin,omitempty"`
	// Array of 10-16 characters Got It voucher codes
	Codes []string `json:"codes,omitempty"`
	// Bill number will apply vouchers
	BillNumber *string `json:"bill_number,omitempty"`
	// Total bill amount
	TotalBill *int64 `json:"total_bill,omitempty"`
	// When true the system will execute the flow without reserve
	SkipReservedWhenMarkUsed *bool `json:"skip_reserved_when_mark_used,omitempty"`
	// SKU information in bill_number
	SkusInfo []RequestCheckMultipleBodySchemaSkusInfoInner `json:"skus_info,omitempty"`
}

// NewRequestMarkUseMultipleBodySchema instantiates a new RequestMarkUseMultipleBodySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMarkUseMultipleBodySchema() *RequestMarkUseMultipleBodySchema {
	this := RequestMarkUseMultipleBodySchema{}
	return &this
}

// NewRequestMarkUseMultipleBodySchemaWithDefaults instantiates a new RequestMarkUseMultipleBodySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMarkUseMultipleBodySchemaWithDefaults() *RequestMarkUseMultipleBodySchema {
	this := RequestMarkUseMultipleBodySchema{}
	return &this
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *RequestMarkUseMultipleBodySchema) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMarkUseMultipleBodySchema) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *RequestMarkUseMultipleBodySchema) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *RequestMarkUseMultipleBodySchema) SetPin(v string) {
	o.Pin = &v
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *RequestMarkUseMultipleBodySchema) GetCodes() []string {
	if o == nil || IsNil(o.Codes) {
		var ret []string
		return ret
	}
	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMarkUseMultipleBodySchema) GetCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *RequestMarkUseMultipleBodySchema) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given []string and assigns it to the Codes field.
func (o *RequestMarkUseMultipleBodySchema) SetCodes(v []string) {
	o.Codes = v
}

// GetBillNumber returns the BillNumber field value if set, zero value otherwise.
func (o *RequestMarkUseMultipleBodySchema) GetBillNumber() string {
	if o == nil || IsNil(o.BillNumber) {
		var ret string
		return ret
	}
	return *o.BillNumber
}

// GetBillNumberOk returns a tuple with the BillNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMarkUseMultipleBodySchema) GetBillNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BillNumber) {
		return nil, false
	}
	return o.BillNumber, true
}

// HasBillNumber returns a boolean if a field has been set.
func (o *RequestMarkUseMultipleBodySchema) HasBillNumber() bool {
	if o != nil && !IsNil(o.BillNumber) {
		return true
	}

	return false
}

// SetBillNumber gets a reference to the given string and assigns it to the BillNumber field.
func (o *RequestMarkUseMultipleBodySchema) SetBillNumber(v string) {
	o.BillNumber = &v
}

// GetTotalBill returns the TotalBill field value if set, zero value otherwise.
func (o *RequestMarkUseMultipleBodySchema) GetTotalBill() int64 {
	if o == nil || IsNil(o.TotalBill) {
		var ret int64
		return ret
	}
	return *o.TotalBill
}

// GetTotalBillOk returns a tuple with the TotalBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMarkUseMultipleBodySchema) GetTotalBillOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalBill) {
		return nil, false
	}
	return o.TotalBill, true
}

// HasTotalBill returns a boolean if a field has been set.
func (o *RequestMarkUseMultipleBodySchema) HasTotalBill() bool {
	if o != nil && !IsNil(o.TotalBill) {
		return true
	}

	return false
}

// SetTotalBill gets a reference to the given int64 and assigns it to the TotalBill field.
func (o *RequestMarkUseMultipleBodySchema) SetTotalBill(v int64) {
	o.TotalBill = &v
}

// GetSkipReservedWhenMarkUsed returns the SkipReservedWhenMarkUsed field value if set, zero value otherwise.
func (o *RequestMarkUseMultipleBodySchema) GetSkipReservedWhenMarkUsed() bool {
	if o == nil || IsNil(o.SkipReservedWhenMarkUsed) {
		var ret bool
		return ret
	}
	return *o.SkipReservedWhenMarkUsed
}

// GetSkipReservedWhenMarkUsedOk returns a tuple with the SkipReservedWhenMarkUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMarkUseMultipleBodySchema) GetSkipReservedWhenMarkUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipReservedWhenMarkUsed) {
		return nil, false
	}
	return o.SkipReservedWhenMarkUsed, true
}

// HasSkipReservedWhenMarkUsed returns a boolean if a field has been set.
func (o *RequestMarkUseMultipleBodySchema) HasSkipReservedWhenMarkUsed() bool {
	if o != nil && !IsNil(o.SkipReservedWhenMarkUsed) {
		return true
	}

	return false
}

// SetSkipReservedWhenMarkUsed gets a reference to the given bool and assigns it to the SkipReservedWhenMarkUsed field.
func (o *RequestMarkUseMultipleBodySchema) SetSkipReservedWhenMarkUsed(v bool) {
	o.SkipReservedWhenMarkUsed = &v
}

// GetSkusInfo returns the SkusInfo field value if set, zero value otherwise.
func (o *RequestMarkUseMultipleBodySchema) GetSkusInfo() []RequestCheckMultipleBodySchemaSkusInfoInner {
	if o == nil || IsNil(o.SkusInfo) {
		var ret []RequestCheckMultipleBodySchemaSkusInfoInner
		return ret
	}
	return o.SkusInfo
}

// GetSkusInfoOk returns a tuple with the SkusInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMarkUseMultipleBodySchema) GetSkusInfoOk() ([]RequestCheckMultipleBodySchemaSkusInfoInner, bool) {
	if o == nil || IsNil(o.SkusInfo) {
		return nil, false
	}
	return o.SkusInfo, true
}

// HasSkusInfo returns a boolean if a field has been set.
func (o *RequestMarkUseMultipleBodySchema) HasSkusInfo() bool {
	if o != nil && !IsNil(o.SkusInfo) {
		return true
	}

	return false
}

// SetSkusInfo gets a reference to the given []RequestCheckMultipleBodySchemaSkusInfoInner and assigns it to the SkusInfo field.
func (o *RequestMarkUseMultipleBodySchema) SetSkusInfo(v []RequestCheckMultipleBodySchemaSkusInfoInner) {
	o.SkusInfo = v
}

func (o RequestMarkUseMultipleBodySchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestMarkUseMultipleBodySchema) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SkipReservedWhenMarkUsed) {
		toSerialize["skip_reserved_when_mark_used"] = o.SkipReservedWhenMarkUsed
	}
	if !IsNil(o.SkusInfo) {
		toSerialize["skus_info"] = o.SkusInfo
	}
	return toSerialize, nil
}

type NullableRequestMarkUseMultipleBodySchema struct {
	value *RequestMarkUseMultipleBodySchema
	isSet bool
}

func (v NullableRequestMarkUseMultipleBodySchema) Get() *RequestMarkUseMultipleBodySchema {
	return v.value
}

func (v *NullableRequestMarkUseMultipleBodySchema) Set(val *RequestMarkUseMultipleBodySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMarkUseMultipleBodySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMarkUseMultipleBodySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMarkUseMultipleBodySchema(val *RequestMarkUseMultipleBodySchema) *NullableRequestMarkUseMultipleBodySchema {
	return &NullableRequestMarkUseMultipleBodySchema{value: val, isSet: true}
}

func (v NullableRequestMarkUseMultipleBodySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMarkUseMultipleBodySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


