// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package biotestcomparitiondatafinder

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockBiotestComparitionDataFinder is an autogenerated mock type for the BiotestComparitionDataFinder type
type MockBiotestComparitionDataFinder struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, req
func (_m *MockBiotestComparitionDataFinder) Run(ctx context.Context, req *BiotestComparitionDataFinderRequest) (*BiotestComparitionDataFinderResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *BiotestComparitionDataFinderResponse
	if rf, ok := ret.Get(0).(func(context.Context, *BiotestComparitionDataFinderRequest) *BiotestComparitionDataFinderResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BiotestComparitionDataFinderResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *BiotestComparitionDataFinderRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
