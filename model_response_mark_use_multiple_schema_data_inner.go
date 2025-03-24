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

// checks if the ResponseMarkUseMultipleSchemaDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMarkUseMultipleSchemaDataInner{}

// ResponseMarkUseMultipleSchemaDataInner struct for ResponseMarkUseMultipleSchemaDataInner
type ResponseMarkUseMultipleSchemaDataInner struct {
	// Voucher code
	Code *string `json:"code,omitempty"`
	// Value of voucher
	Value NullableInt32 `json:"value,omitempty"`
	// State of voucher
	State NullableInt32 `json:"state,omitempty"`
	// Voucher type, standard or conditional
	VoucherType *string `json:"voucher_type,omitempty"`
	Conditions *ResponseMarkUseMultipleSchemaDataInnerConditions `json:"conditions,omitempty"`
	Redemptions *ResponseMarkUseMultipleSchemaDataInnerRedemptions `json:"redemptions,omitempty"`
}

// NewResponseMarkUseMultipleSchemaDataInner instantiates a new ResponseMarkUseMultipleSchemaDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMarkUseMultipleSchemaDataInner() *ResponseMarkUseMultipleSchemaDataInner {
	this := ResponseMarkUseMultipleSchemaDataInner{}
	return &this
}

// NewResponseMarkUseMultipleSchemaDataInnerWithDefaults instantiates a new ResponseMarkUseMultipleSchemaDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMarkUseMultipleSchemaDataInnerWithDefaults() *ResponseMarkUseMultipleSchemaDataInner {
	this := ResponseMarkUseMultipleSchemaDataInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ResponseMarkUseMultipleSchemaDataInner) SetCode(v string) {
	o.Code = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseMarkUseMultipleSchemaDataInner) GetValue() int32 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret int32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseMarkUseMultipleSchemaDataInner) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableInt32 and assigns it to the Value field.
func (o *ResponseMarkUseMultipleSchemaDataInner) SetValue(v int32) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ResponseMarkUseMultipleSchemaDataInner) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ResponseMarkUseMultipleSchemaDataInner) UnsetValue() {
	o.Value.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseMarkUseMultipleSchemaDataInner) GetState() int32 {
	if o == nil || IsNil(o.State.Get()) {
		var ret int32
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseMarkUseMultipleSchemaDataInner) GetStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableInt32 and assigns it to the State field.
func (o *ResponseMarkUseMultipleSchemaDataInner) SetState(v int32) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *ResponseMarkUseMultipleSchemaDataInner) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *ResponseMarkUseMultipleSchemaDataInner) UnsetState() {
	o.State.Unset()
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *ResponseMarkUseMultipleSchemaDataInner) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetConditions() ResponseMarkUseMultipleSchemaDataInnerConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret ResponseMarkUseMultipleSchemaDataInnerConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetConditionsOk() (*ResponseMarkUseMultipleSchemaDataInnerConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given ResponseMarkUseMultipleSchemaDataInnerConditions and assigns it to the Conditions field.
func (o *ResponseMarkUseMultipleSchemaDataInner) SetConditions(v ResponseMarkUseMultipleSchemaDataInnerConditions) {
	o.Conditions = &v
}

// GetRedemptions returns the Redemptions field value if set, zero value otherwise.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetRedemptions() ResponseMarkUseMultipleSchemaDataInnerRedemptions {
	if o == nil || IsNil(o.Redemptions) {
		var ret ResponseMarkUseMultipleSchemaDataInnerRedemptions
		return ret
	}
	return *o.Redemptions
}

// GetRedemptionsOk returns a tuple with the Redemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) GetRedemptionsOk() (*ResponseMarkUseMultipleSchemaDataInnerRedemptions, bool) {
	if o == nil || IsNil(o.Redemptions) {
		return nil, false
	}
	return o.Redemptions, true
}

// HasRedemptions returns a boolean if a field has been set.
func (o *ResponseMarkUseMultipleSchemaDataInner) HasRedemptions() bool {
	if o != nil && !IsNil(o.Redemptions) {
		return true
	}

	return false
}

// SetRedemptions gets a reference to the given ResponseMarkUseMultipleSchemaDataInnerRedemptions and assigns it to the Redemptions field.
func (o *ResponseMarkUseMultipleSchemaDataInner) SetRedemptions(v ResponseMarkUseMultipleSchemaDataInnerRedemptions) {
	o.Redemptions = &v
}

func (o ResponseMarkUseMultipleSchemaDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMarkUseMultipleSchemaDataInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Redemptions) {
		toSerialize["redemptions"] = o.Redemptions
	}
	return toSerialize, nil
}

type NullableResponseMarkUseMultipleSchemaDataInner struct {
	value *ResponseMarkUseMultipleSchemaDataInner
	isSet bool
}

func (v NullableResponseMarkUseMultipleSchemaDataInner) Get() *ResponseMarkUseMultipleSchemaDataInner {
	return v.value
}

func (v *NullableResponseMarkUseMultipleSchemaDataInner) Set(val *ResponseMarkUseMultipleSchemaDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMarkUseMultipleSchemaDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMarkUseMultipleSchemaDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMarkUseMultipleSchemaDataInner(val *ResponseMarkUseMultipleSchemaDataInner) *NullableResponseMarkUseMultipleSchemaDataInner {
	return &NullableResponseMarkUseMultipleSchemaDataInner{value: val, isSet: true}
}

func (v NullableResponseMarkUseMultipleSchemaDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMarkUseMultipleSchemaDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


