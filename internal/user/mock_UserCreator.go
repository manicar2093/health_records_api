// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package user

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockUserCreator is an autogenerated mock type for the UserCreator type
type MockUserCreator struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, user
func (_m *MockUserCreator) Run(ctx context.Context, user *UserCreatorRequest) (*UserCreatorResponse, error) {
	ret := _m.Called(ctx, user)

	var r0 *UserCreatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UserCreatorRequest) *UserCreatorResponse); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserCreatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UserCreatorRequest) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}