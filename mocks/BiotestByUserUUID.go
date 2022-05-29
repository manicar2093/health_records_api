// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	context "context"

	biotestfilters "github.com/manicar2093/charly_team_api/internal/biotestfilters"

	mock "github.com/stretchr/testify/mock"
)

// BiotestByUserUUID is an autogenerated mock type for the BiotestByUserUUID type
type BiotestByUserUUID struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, req
func (_m *BiotestByUserUUID) Run(ctx context.Context, req *biotestfilters.BiotestByUserUUIDRequest) (*biotestfilters.BiotestByUserUUIDResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *biotestfilters.BiotestByUserUUIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *biotestfilters.BiotestByUserUUIDRequest) *biotestfilters.BiotestByUserUUIDResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biotestfilters.BiotestByUserUUIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *biotestfilters.BiotestByUserUUIDRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
