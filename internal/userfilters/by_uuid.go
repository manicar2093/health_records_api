package userfilters

import (
	"context"

	"github.com/manicar2093/charly_team_api/internal/db/repositories"
	"github.com/manicar2093/charly_team_api/pkg/logger"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

type UserByUUIDFinder interface {
	Run(ctx context.Context, req *UserByUUIDFinderRequest) (*UserByUUIDFinderResponse, error)
}

type userByUUIDFinderImpl struct {
	userRepo  repositories.UserRepository
	validator validators.ValidatorService
}

func NewUserByUUIDFinderImpl(
	userRepo repositories.UserRepository,
	validator validators.ValidatorService,
) *userByUUIDFinderImpl {
	return &userByUUIDFinderImpl{userRepo: userRepo, validator: validator}
}

func (c *userByUUIDFinderImpl) Run(ctx context.Context, req *UserByUUIDFinderRequest) (*UserByUUIDFinderResponse, error) {
	logger.Info(req)
	if validation := c.validator.Validate(req); !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	user, err := c.userRepo.FindUserByUUID(ctx, req.UserUUID)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &UserByUUIDFinderResponse{UserFound: user}, nil
}