// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/manicar2093/health_records/internal/user"
	mock "github.com/stretchr/testify/mock"
)

// AvatarUpdater is an autogenerated mock type for the AvatarUpdater type
type AvatarUpdater struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, req
func (_m *AvatarUpdater) Run(ctx context.Context, req *user.AvatarUpdaterRequest) (*user.AvatarUpdaterResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *user.AvatarUpdaterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *user.AvatarUpdaterRequest) *user.AvatarUpdaterResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.AvatarUpdaterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.AvatarUpdaterRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
