/*
Metric Ruleset API

 Metric ruleset API 

API for creating, updating, retrieving, and deleting metric rulesets, which serve as the driving mechanisms behind metrics pipeline management in Splunk Observability Cloud.  ## Requirements  * You must have an organization access token with the API permission or a session token to use the API. * You have to have the Splunk Observability Cloud admin, power, or read_only role to use the `GET /metricruleset/{id}` operations. * You have to have the Splunk Observability Cloud admin or power role to use the `POST /metricruleset`, `PUT /metricruleset/{id}`, `DELETE /metricruleset/{id}`, and `POST /metricruleset/generateAggregationMetricName` operations.

API version: 3.3.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metric_ruleset

import (
	"encoding/json"
)

// ExceptionRuleRestorationFields Fields of the exceptionRules restoration element. 
type ExceptionRuleRestorationFields struct {
	// ID of the restoration job. 
	RestorationId *string `json:"restorationId,omitempty"`
	// Time from which the restoration job will restore archived data, in the form of *nix time in milliseconds.
	StartTime *int64 `json:"startTime,omitempty"`
	// Time to which the restoration job will restore archived data, in the form of *nix time in milliseconds.
	StopTime *int64 `json:"stopTime,omitempty"`
}

// NewExceptionRuleRestorationFields instantiates a new ExceptionRuleRestorationFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionRuleRestorationFields() *ExceptionRuleRestorationFields {
	this := ExceptionRuleRestorationFields{}
	return &this
}

// NewExceptionRuleRestorationFieldsWithDefaults instantiates a new ExceptionRuleRestorationFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionRuleRestorationFieldsWithDefaults() *ExceptionRuleRestorationFields {
	this := ExceptionRuleRestorationFields{}
	return &this
}

// GetRestorationId returns the RestorationId field value if set, zero value otherwise.
func (o *ExceptionRuleRestorationFields) GetRestorationId() string {
	if o == nil || isNil(o.RestorationId) {
		var ret string
		return ret
	}
	return *o.RestorationId
}

// GetRestorationIdOk returns a tuple with the RestorationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionRuleRestorationFields) GetRestorationIdOk() (*string, bool) {
	if o == nil || isNil(o.RestorationId) {
		return nil, false
	}
	return o.RestorationId, true
}

// HasRestorationId returns a boolean if a field has been set.
func (o *ExceptionRuleRestorationFields) HasRestorationId() bool {
	if o != nil && !isNil(o.RestorationId) {
		return true
	}

	return false
}

// SetRestorationId gets a reference to the given string and assigns it to the RestorationId field.
func (o *ExceptionRuleRestorationFields) SetRestorationId(v string) {
	o.RestorationId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ExceptionRuleRestorationFields) GetStartTime() int64 {
	if o == nil || isNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionRuleRestorationFields) GetStartTimeOk() (*int64, bool) {
	if o == nil || isNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ExceptionRuleRestorationFields) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *ExceptionRuleRestorationFields) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *ExceptionRuleRestorationFields) GetStopTime() int64 {
	if o == nil || isNil(o.StopTime) {
		var ret int64
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionRuleRestorationFields) GetStopTimeOk() (*int64, bool) {
	if o == nil || isNil(o.StopTime) {
		return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *ExceptionRuleRestorationFields) HasStopTime() bool {
	if o != nil && !isNil(o.StopTime) {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given int64 and assigns it to the StopTime field.
func (o *ExceptionRuleRestorationFields) SetStopTime(v int64) {
	o.StopTime = &v
}

func (o ExceptionRuleRestorationFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RestorationId) && len(*o.RestorationId) > 0 {
		toSerialize["restorationId"] = o.RestorationId
	}
	if !isNil(o.StartTime) && *o.StartTime > int64(0) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.StopTime) && *o.StopTime > int64(0) {
		toSerialize["stopTime"] = o.StopTime
	}
	return json.Marshal(toSerialize)
}

type NullableExceptionRuleRestorationFields struct {
	value *ExceptionRuleRestorationFields
	isSet bool
}

func (v NullableExceptionRuleRestorationFields) Get() *ExceptionRuleRestorationFields {
	return v.value
}

func (v *NullableExceptionRuleRestorationFields) Set(val *ExceptionRuleRestorationFields) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionRuleRestorationFields) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionRuleRestorationFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionRuleRestorationFields(val *ExceptionRuleRestorationFields) *NullableExceptionRuleRestorationFields {
	return &NullableExceptionRuleRestorationFields{value: val, isSet: true}
}

func (v NullableExceptionRuleRestorationFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionRuleRestorationFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


