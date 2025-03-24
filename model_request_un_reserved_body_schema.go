/*
Merchant APIs

Technical document APIs for Merchant APIs

API version: 6.0
Contact: duong.vu@gotit.vn
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gotit-merchant-apis

import (
	"encoding/json"
)

// checks if the RequestUnReservedBodySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestUnReservedBodySchema{}

// RequestUnReservedBodySchema struct for RequestUnReservedBodySchema
type RequestUnReservedBodySchema struct {
	// Store pin
	Pin *string `json:"pin,omitempty"`
	// Array of 10-16 characters Got It voucher codes
	Codes []string `json:"codes,omitempty"`
	// Bill number will apply vouchers
	BillNumber *string `json:"bill_number,omitempty"`
	// Bill creation time. Format: YYYY-MM-DD HH:MM:SS
	BillCreatedAt *string `json:"bill_created_at,omitempty"`
}

// NewRequestUnReservedBodySchema instantiates a new RequestUnReservedBodySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestUnReservedBodySchema() *RequestUnReservedBodySchema {
	this := RequestUnReservedBodySchema{}
	return &this
}

// NewRequestUnReservedBodySchemaWithDefaults instantiates a new RequestUnReservedBodySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestUnReservedBodySchemaWithDefaults() *RequestUnReservedBodySchema {
	this := RequestUnReservedBodySchema{}
	return &this
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *RequestUnReservedBodySchema) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUnReservedBodySchema) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *RequestUnReservedBodySchema) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *RequestUnReservedBodySchema) SetPin(v string) {
	o.Pin = &v
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *RequestUnReservedBodySchema) GetCodes() []string {
	if o == nil || IsNil(o.Codes) {
		var ret []string
		return ret
	}
	return o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUnReservedBodySchema) GetCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *RequestUnReservedBodySchema) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given []string and assigns it to the Codes field.
func (o *RequestUnReservedBodySchema) SetCodes(v []string) {
	o.Codes = v
}

// GetBillNumber returns the BillNumber field value if set, zero value otherwise.
func (o *RequestUnReservedBodySchema) GetBillNumber() string {
	if o == nil || IsNil(o.BillNumber) {
		var ret string
		return ret
	}
	return *o.BillNumber
}

// GetBillNumberOk returns a tuple with the BillNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUnReservedBodySchema) GetBillNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BillNumber) {
		return nil, false
	}
	return o.BillNumber, true
}

// HasBillNumber returns a boolean if a field has been set.
func (o *RequestUnReservedBodySchema) HasBillNumber() bool {
	if o != nil && !IsNil(o.BillNumber) {
		return true
	}

	return false
}

// SetBillNumber gets a reference to the given string and assigns it to the BillNumber field.
func (o *RequestUnReservedBodySchema) SetBillNumber(v string) {
	o.BillNumber = &v
}

// GetBillCreatedAt returns the BillCreatedAt field value if set, zero value otherwise.
func (o *RequestUnReservedBodySchema) GetBillCreatedAt() string {
	if o == nil || IsNil(o.BillCreatedAt) {
		var ret string
		return ret
	}
	return *o.BillCreatedAt
}

// GetBillCreatedAtOk returns a tuple with the BillCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUnReservedBodySchema) GetBillCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.BillCreatedAt) {
		return nil, false
	}
	return o.BillCreatedAt, true
}

// HasBillCreatedAt returns a boolean if a field has been set.
func (o *RequestUnReservedBodySchema) HasBillCreatedAt() bool {
	if o != nil && !IsNil(o.BillCreatedAt) {
		return true
	}

	return false
}

// SetBillCreatedAt gets a reference to the given string and assigns it to the BillCreatedAt field.
func (o *RequestUnReservedBodySchema) SetBillCreatedAt(v string) {
	o.BillCreatedAt = &v
}

func (o RequestUnReservedBodySchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestUnReservedBodySchema) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BillCreatedAt) {
		toSerialize["bill_created_at"] = o.BillCreatedAt
	}
	return toSerialize, nil
}

type NullableRequestUnReservedBodySchema struct {
	value *RequestUnReservedBodySchema
	isSet bool
}

func (v NullableRequestUnReservedBodySchema) Get() *RequestUnReservedBodySchema {
	return v.value
}

func (v *NullableRequestUnReservedBodySchema) Set(val *RequestUnReservedBodySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestUnReservedBodySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestUnReservedBodySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestUnReservedBodySchema(val *RequestUnReservedBodySchema) *NullableRequestUnReservedBodySchema {
	return &NullableRequestUnReservedBodySchema{value: val, isSet: true}
}

func (v NullableRequestUnReservedBodySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestUnReservedBodySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


