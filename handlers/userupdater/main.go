package userupdater

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/internal/apperrors"
	"github.com/manicar2093/charly_team_api/internal/logger"
	"github.com/manicar2093/charly_team_api/validators"
)

type UserUpdater interface {
	Run(ctx context.Context, userData *entities.User) (*UserUpdaterResponse, error)
}

type UserUpdaterImpl struct {
	repo      rel.Repository
	validator validators.ValidatorService
}

func NewUpdateUser(repo rel.Repository, validator validators.ValidatorService) *UserUpdaterImpl {
	return &UserUpdaterImpl{
		repo:      repo,
		validator: validator,
	}
}

func (c *UserUpdaterImpl) Run(ctx context.Context, userData *entities.User) (*UserUpdaterResponse, error) {

	logger.Info(userData)

	if !validators.IsUpdateRequestValid(userData) {
		logger.Error("identifier miss")
		return nil, apperrors.ValidationErrors{{Field: "identifier", Validation: "required"}}
	}

	dataValidation := c.validator.Validate(userData)

	if !dataValidation.IsValid {
		logger.Error(dataValidation.Err)
		return nil, dataValidation.Err
	}

	err := c.repo.Update(ctx, userData)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &UserUpdaterResponse{userData}, nil
}
