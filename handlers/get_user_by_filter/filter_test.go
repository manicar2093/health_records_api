package main

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/mocks"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserFilterTest struct {
	suite.Suite
	repo                         *reltest.Repository
	paginator                    *mocks.Paginable
	ctx                          context.Context
	notFoundError, ordinaryError error
}

func (c *UserFilterTest) SetupTest() {
	c.repo = reltest.New()
	c.ctx = context.Background()
	c.ordinaryError = errors.New("An ordinary error :O")
	c.notFoundError = rel.NotFoundError{}
	c.paginator = &mocks.Paginable{}

}

func (c *UserFilterTest) TearDownTest() {
	c.repo.AssertExpectations(c.T())
	c.paginator.AssertExpectations(c.T())
}

func (c *UserFilterTest) TestFilterUserByID() {

	userIDRequested := 1

	request := map[string]interface{}{
		"user_id": userIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("id", userIDRequested),
	).Result(
		entities.User{
			ID: int32(userIDRequested),
		},
	)

	got, err := FindUserByID(c.ctx, c.repo, request, c.paginator)

	c.Nil(err, "FindUserByID return an error")

	userGot, ok := got.(entities.User)
	c.True(ok, "unexpected answare type")
	c.Equal(userGot.ID, int32(userIDRequested), "unexpected user id")

}

func (c *UserFilterTest) TestFilterUserByIDValidatioError() {

	request := map[string]interface{}{}

	_, err := FindUserByID(c.ctx, c.repo, request, c.paginator)

	validationError, isValidationError := err.(apperrors.ValidationError)

	c.True(isValidationError, "bad type of error ")

	c.Equal(validationError.Validation, "required")
	c.Equal(validationError.Field, "user_id")

}

func (c *UserFilterTest) TestFilterUserByIDNotFound() {

	userIDRequested := 1

	request := map[string]interface{}{
		"user_id": userIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("id", userIDRequested),
	).Return(c.notFoundError)

	_, err := FindUserByID(c.ctx, c.repo, request, c.paginator)

	_, isHandableNotFoundError := err.(apperrors.UserNotFound)

	c.True(isHandableNotFoundError, "unexpected error type")

}

func (c *UserFilterTest) TestFilterUserByIDUnhandledError() {

	userIDRequested := 1

	request := map[string]interface{}{
		"user_id": userIDRequested,
	}

	c.repo.ExpectFind(
		where.Eq("id", userIDRequested),
	).Return(c.ordinaryError)

	_, err := FindUserByID(c.ctx, c.repo, request, c.paginator)

	c.NotNil(err, "should return error")

}

func (c *UserFilterTest) TestFindUserByEmail() {

	userEmailRequested := "testing@testing.com"

	request := map[string]interface{}{
		"email": userEmailRequested,
	}

	c.repo.ExpectFind(
		where.Like("email", "%"+userEmailRequested+"%"),
	).Result(
		entities.User{
			Email: userEmailRequested,
		},
	)

	got, err := FindUserByEmail(c.ctx, c.repo, request, c.paginator)

	c.Nil(err, "FindUserByID return an error")

	userGot, ok := got.(entities.User)
	c.True(ok, "unexpected answare type")
	c.Equal(userGot.Email, userEmailRequested, "unexpected user id")

}

func (c *UserFilterTest) TestFindUserByEmailValidationError() {

	request := map[string]interface{}{}

	_, err := FindUserByEmail(c.ctx, c.repo, request, c.paginator)

	_, isValidationError := err.(apperrors.ValidationError)
	c.True(isValidationError, "unexpected error type")

}

func (c *UserFilterTest) TestFindUserByEmailNotFoundError() {

	userEmailRequested := "testing@testing.com"

	request := map[string]interface{}{
		"email": userEmailRequested,
	}

	c.repo.ExpectFind(
		where.Like("email", "%"+userEmailRequested+"%"),
	).Return(c.notFoundError)

	_, err := FindUserByEmail(c.ctx, c.repo, request, c.paginator)

	_, isHandableNotFoundError := err.(apperrors.UserNotFound)

	c.True(isHandableNotFoundError, "unexpected error type")

}

func (c *UserFilterTest) TestFindUserByEmailUnhandledError() {

	userEmailRequested := "testing@testing.com"

	request := map[string]interface{}{
		"email": userEmailRequested,
	}

	c.repo.ExpectFind(
		where.Like("email", "%"+userEmailRequested+"%"),
	).Return(c.ordinaryError)

	_, err := FindUserByEmail(c.ctx, c.repo, request, c.paginator)

	c.NotNil(err, "should not return error")

}

func (c *UserFilterTest) TestFindAllUsers() {

	userPageRequested := 2

	request := map[string]interface{}{
		"page_number": userPageRequested,
	}

	c.paginator.On(
		"CreatePaginator",
		c.ctx,
		entities.UserTable,
		mock.Anything,
		userPageRequested,
	).Return(&models.Paginator{}, nil)

	got, err := FindAllUsers(c.ctx, c.repo, request, c.paginator)

	c.Nil(err, "FindUserByID return an error")

	_, ok := got.(*models.Paginator)
	c.True(ok, "unexpected answare type")

}

func (c *UserFilterTest) TestFindAllUsersValidationError() {

	request := map[string]interface{}{}

	_, err := FindAllUsers(c.ctx, c.repo, request, c.paginator)

	_, isValidationError := err.(apperrors.ValidationError)
	c.True(isValidationError, "unexpected error type")

}

func TestUserFilter(t *testing.T) {
	suite.Run(t, new(UserFilterTest))
}