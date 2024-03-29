// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	cognitoidentityprovider "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	mock "github.com/stretchr/testify/mock"
)

// CongitoClient is an autogenerated mock type for the CongitoClient type
type CongitoClient struct {
	mock.Mock
}

// AdminCreateUser provides a mock function with given fields: input
func (_m *CongitoClient) AdminCreateUser(input *cognitoidentityprovider.AdminCreateUserInput) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
	ret := _m.Called(input)

	var r0 *cognitoidentityprovider.AdminCreateUserOutput
	if rf, ok := ret.Get(0).(func(*cognitoidentityprovider.AdminCreateUserInput) *cognitoidentityprovider.AdminCreateUserOutput); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cognitoidentityprovider.AdminCreateUserOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cognitoidentityprovider.AdminCreateUserInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
