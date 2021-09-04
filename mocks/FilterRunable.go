// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	filters "github.com/manicar2093/charly_team_api/db/filters"
	mock "github.com/stretchr/testify/mock"
)

// FilterRunable is an autogenerated mock type for the FilterRunable type
type FilterRunable struct {
	mock.Mock
}

// IsFound provides a mock function with given fields:
func (_m *FilterRunable) IsFound() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Run provides a mock function with given fields: filterParameters
func (_m *FilterRunable) Run(filterParameters *filters.FilterParameters) (interface{}, error) {
	ret := _m.Called(filterParameters)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*filters.FilterParameters) interface{}); ok {
		r0 = rf(filterParameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*filters.FilterParameters) error); ok {
		r1 = rf(filterParameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}