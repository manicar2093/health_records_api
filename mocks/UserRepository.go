// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/manicar2093/health_records/internal/db/entities"
	mock "github.com/stretchr/testify/mock"

	paginator "github.com/manicar2093/health_records/internal/db/paginator"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// FindAllUsers provides a mock function with given fields: ctx, pageSort
func (_m *UserRepository) FindAllUsers(ctx context.Context, pageSort *paginator.PageSort) (*paginator.Paginator, error) {
	ret := _m.Called(ctx, pageSort)

	var r0 *paginator.Paginator
	if rf, ok := ret.Get(0).(func(context.Context, *paginator.PageSort) *paginator.Paginator); ok {
		r0 = rf(ctx, pageSort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paginator.Paginator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *paginator.PageSort) error); ok {
		r1 = rf(ctx, pageSort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByUUID provides a mock function with given fields: ctx, userUUID
func (_m *UserRepository) FindUserByUUID(ctx context.Context, userUUID string) (*entities.User, error) {
	ret := _m.Called(ctx, userUUID)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, userUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserLikeEmailOrNameOrLastName provides a mock function with given fields: ctx, parameter
func (_m *UserRepository) FindUserLikeEmailOrNameOrLastName(ctx context.Context, parameter string) (*[]entities.User, error) {
	ret := _m.Called(ctx, parameter)

	var r0 *[]entities.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *[]entities.User); ok {
		r0 = rf(ctx, parameter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, parameter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: ctx, user
func (_m *UserRepository) SaveUser(ctx context.Context, user *entities.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *UserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
