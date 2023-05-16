/*
Copyright 2021 The cmp authors .

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// StudentSpecApplyConfiguration represents an declarative configuration of the StudentSpec type for use
// with apply.
type StudentSpecApplyConfiguration struct {
	name   *string `json:"name,omitempty"`
	school *string `json:"school,omitempty"`
}

// StudentSpecApplyConfiguration constructs an declarative configuration of the StudentSpec type for use with
// apply.
func StudentSpec() *StudentSpecApplyConfiguration {
	return &StudentSpecApplyConfiguration{}
}

// Withname sets the name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the name field is set to the value of the last call.
func (b *StudentSpecApplyConfiguration) Withname(value string) *StudentSpecApplyConfiguration {
	b.name = &value
	return b
}

// Withschool sets the school field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the school field is set to the value of the last call.
func (b *StudentSpecApplyConfiguration) Withschool(value string) *StudentSpecApplyConfiguration {
	b.school = &value
	return b
}
