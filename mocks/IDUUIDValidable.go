// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IDUUIDValidable is an autogenerated mock type for the IDUUIDValidable type
type IDUUIDValidable struct {
	mock.Mock
}

// GetID provides a mock function with given fields:
func (_m *IDUUIDValidable) GetID() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// GetUUID provides a mock function with given fields:
func (_m *IDUUIDValidable) GetUUID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
