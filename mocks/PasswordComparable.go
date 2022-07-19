// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordComparable is an autogenerated mock type for the PasswordComparable type
type PasswordComparable struct {
	mock.Mock
}

// Compare provides a mock function with given fields: hashed, password
func (_m *PasswordComparable) Compare(hashed string, password string) error {
	ret := _m.Called(hashed, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(hashed, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
