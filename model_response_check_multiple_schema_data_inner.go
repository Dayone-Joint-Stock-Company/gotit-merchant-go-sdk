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

// checks if the ResponseCheckMultipleSchemaDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCheckMultipleSchemaDataInner{}

// ResponseCheckMultipleSchemaDataInner struct for ResponseCheckMultipleSchemaDataInner
type ResponseCheckMultipleSchemaDataInner struct {
	// Voucher code
	Code *string `json:"code,omitempty"`
	// Value of voucher
	Value NullableInt32 `json:"value,omitempty"`
	// State of voucher
	State NullableInt32 `json:"state,omitempty"`
	// Voucher type, standard or conditional
	VoucherType *string `json:"voucher_type,omitempty"`
	// Expiry date of voucher (YYYY-MM-DD)
	ExpiryDate *string `json:"expiry_date,omitempty"`
	// Date cancel voucher (YYYY-MM-DD)
	CancelDate *string `json:"cancel_date,omitempty"`
	Conditions *ResponseCheckMultipleSchemaDataInnerConditions `json:"conditions,omitempty"`
	Redemptions *ResponseCheckMultipleSchemaDataInnerRedemptions `json:"redemptions,omitempty"`
}

// NewResponseCheckMultipleSchemaDataInner instantiates a new ResponseCheckMultipleSchemaDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCheckMultipleSchemaDataInner() *ResponseCheckMultipleSchemaDataInner {
	this := ResponseCheckMultipleSchemaDataInner{}
	return &this
}

// NewResponseCheckMultipleSchemaDataInnerWithDefaults instantiates a new ResponseCheckMultipleSchemaDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCheckMultipleSchemaDataInnerWithDefaults() *ResponseCheckMultipleSchemaDataInner {
	this := ResponseCheckMultipleSchemaDataInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ResponseCheckMultipleSchemaDataInner) SetCode(v string) {
	o.Code = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCheckMultipleSchemaDataInner) GetValue() int32 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret int32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCheckMultipleSchemaDataInner) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableInt32 and assigns it to the Value field.
func (o *ResponseCheckMultipleSchemaDataInner) SetValue(v int32) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ResponseCheckMultipleSchemaDataInner) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ResponseCheckMultipleSchemaDataInner) UnsetValue() {
	o.Value.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCheckMultipleSchemaDataInner) GetState() int32 {
	if o == nil || IsNil(o.State.Get()) {
		var ret int32
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCheckMultipleSchemaDataInner) GetStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableInt32 and assigns it to the State field.
func (o *ResponseCheckMultipleSchemaDataInner) SetState(v int32) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *ResponseCheckMultipleSchemaDataInner) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *ResponseCheckMultipleSchemaDataInner) UnsetState() {
	o.State.Unset()
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInner) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInner) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *ResponseCheckMultipleSchemaDataInner) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInner) GetExpiryDate() string {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInner) GetExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *ResponseCheckMultipleSchemaDataInner) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetCancelDate returns the CancelDate field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInner) GetCancelDate() string {
	if o == nil || IsNil(o.CancelDate) {
		var ret string
		return ret
	}
	return *o.CancelDate
}

// GetCancelDateOk returns a tuple with the CancelDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInner) GetCancelDateOk() (*string, bool) {
	if o == nil || IsNil(o.CancelDate) {
		return nil, false
	}
	return o.CancelDate, true
}

// HasCancelDate returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasCancelDate() bool {
	if o != nil && !IsNil(o.CancelDate) {
		return true
	}

	return false
}

// SetCancelDate gets a reference to the given string and assigns it to the CancelDate field.
func (o *ResponseCheckMultipleSchemaDataInner) SetCancelDate(v string) {
	o.CancelDate = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInner) GetConditions() ResponseCheckMultipleSchemaDataInnerConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret ResponseCheckMultipleSchemaDataInnerConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInner) GetConditionsOk() (*ResponseCheckMultipleSchemaDataInnerConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given ResponseCheckMultipleSchemaDataInnerConditions and assigns it to the Conditions field.
func (o *ResponseCheckMultipleSchemaDataInner) SetConditions(v ResponseCheckMultipleSchemaDataInnerConditions) {
	o.Conditions = &v
}

// GetRedemptions returns the Redemptions field value if set, zero value otherwise.
func (o *ResponseCheckMultipleSchemaDataInner) GetRedemptions() ResponseCheckMultipleSchemaDataInnerRedemptions {
	if o == nil || IsNil(o.Redemptions) {
		var ret ResponseCheckMultipleSchemaDataInnerRedemptions
		return ret
	}
	return *o.Redemptions
}

// GetRedemptionsOk returns a tuple with the Redemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCheckMultipleSchemaDataInner) GetRedemptionsOk() (*ResponseCheckMultipleSchemaDataInnerRedemptions, bool) {
	if o == nil || IsNil(o.Redemptions) {
		return nil, false
	}
	return o.Redemptions, true
}

// HasRedemptions returns a boolean if a field has been set.
func (o *ResponseCheckMultipleSchemaDataInner) HasRedemptions() bool {
	if o != nil && !IsNil(o.Redemptions) {
		return true
	}

	return false
}

// SetRedemptions gets a reference to the given ResponseCheckMultipleSchemaDataInnerRedemptions and assigns it to the Redemptions field.
func (o *ResponseCheckMultipleSchemaDataInner) SetRedemptions(v ResponseCheckMultipleSchemaDataInnerRedemptions) {
	o.Redemptions = &v
}

func (o ResponseCheckMultipleSchemaDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCheckMultipleSchemaDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !IsNil(o.VoucherType) {
		toSerialize["voucher_type"] = o.VoucherType
	}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if !IsNil(o.CancelDate) {
		toSerialize["cancel_date"] = o.CancelDate
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Redemptions) {
		toSerialize["redemptions"] = o.Redemptions
	}
	return toSerialize, nil
}

type NullableResponseCheckMultipleSchemaDataInner struct {
	value *ResponseCheckMultipleSchemaDataInner
	isSet bool
}

func (v NullableResponseCheckMultipleSchemaDataInner) Get() *ResponseCheckMultipleSchemaDataInner {
	return v.value
}

func (v *NullableResponseCheckMultipleSchemaDataInner) Set(val *ResponseCheckMultipleSchemaDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCheckMultipleSchemaDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCheckMultipleSchemaDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCheckMultipleSchemaDataInner(val *ResponseCheckMultipleSchemaDataInner) *NullableResponseCheckMultipleSchemaDataInner {
	return &NullableResponseCheckMultipleSchemaDataInner{value: val, isSet: true}
}

func (v NullableResponseCheckMultipleSchemaDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCheckMultipleSchemaDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


