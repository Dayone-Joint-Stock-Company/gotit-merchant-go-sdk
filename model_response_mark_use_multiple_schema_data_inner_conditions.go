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

// checks if the ResponseMarkUseMultipleSchemaDataInnerConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMarkUseMultipleSchemaDataInnerConditions{}

// ResponseMarkUseMultipleSchemaDataInnerConditions Include information involve with voucher type is conditional
type ResponseMarkUseMultipleSchemaDataInnerConditions struct {
	// List of redeemable SKUs of the voucher code. For voucher type = conditional, bill number must contain at least 1 redeemable SKU of the voucher.
	RedeemableSkus []string `json:"redeemable_skus,omitempty"`
}

// NewResponseMarkUseMultipleSchemaDataInnerConditions instantiates a new ResponseMarkUseMultipleSchemaDataInnerConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMarkUseMultipleSchemaDataInnerConditions() *ResponseMarkUseMultipleSchemaDataInnerConditions {
	this := ResponseMarkUseMultipleSchemaDataInnerConditions{}
	return &this
}

// NewResponseMarkUseMultipleSchemaDataInnerConditionsWithDefaults instantiates a new ResponseMarkUseMultipleSchemaDataInnerConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMarkUseMultipleSchemaDataInnerConditionsWithDefaults() *ResponseMarkUseMultipleSchemaDataInnerConditions {
	this := ResponseMarkUseMultipleSchemaDataInnerConditions{}
	return &this
}

// GetRedeemableSkus returns the RedeemableSkus field value if set, zero value otherwise.
func (o *ResponseMarkUseMultipleSchemaDataInnerConditions) GetRedeemableSkus() []string {
	if o == nil || IsNil(o.RedeemableSkus) {
		var ret []string
		return ret
	}
	return o.RedeemableSkus
}

// GetRedeemableSkusOk returns a tuple with the RedeemableSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMarkUseMultipleSchemaDataInnerConditions) GetRedeemableSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.RedeemableSkus) {
		return nil, false
	}
	return o.RedeemableSkus, true
}

// HasRedeemableSkus returns a boolean if a field has been set.
func (o *ResponseMarkUseMultipleSchemaDataInnerConditions) HasRedeemableSkus() bool {
	if o != nil && !IsNil(o.RedeemableSkus) {
		return true
	}

	return false
}

// SetRedeemableSkus gets a reference to the given []string and assigns it to the RedeemableSkus field.
func (o *ResponseMarkUseMultipleSchemaDataInnerConditions) SetRedeemableSkus(v []string) {
	o.RedeemableSkus = v
}

func (o ResponseMarkUseMultipleSchemaDataInnerConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMarkUseMultipleSchemaDataInnerConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RedeemableSkus) {
		toSerialize["redeemable_skus"] = o.RedeemableSkus
	}
	return toSerialize, nil
}

type NullableResponseMarkUseMultipleSchemaDataInnerConditions struct {
	value *ResponseMarkUseMultipleSchemaDataInnerConditions
	isSet bool
}

func (v NullableResponseMarkUseMultipleSchemaDataInnerConditions) Get() *ResponseMarkUseMultipleSchemaDataInnerConditions {
	return v.value
}

func (v *NullableResponseMarkUseMultipleSchemaDataInnerConditions) Set(val *ResponseMarkUseMultipleSchemaDataInnerConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMarkUseMultipleSchemaDataInnerConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMarkUseMultipleSchemaDataInnerConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMarkUseMultipleSchemaDataInnerConditions(val *ResponseMarkUseMultipleSchemaDataInnerConditions) *NullableResponseMarkUseMultipleSchemaDataInnerConditions {
	return &NullableResponseMarkUseMultipleSchemaDataInnerConditions{value: val, isSet: true}
}

func (v NullableResponseMarkUseMultipleSchemaDataInnerConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMarkUseMultipleSchemaDataInnerConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


