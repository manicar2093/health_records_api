package biotestimagessaver

import (
	"context"

	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/internal/logger"
	"github.com/manicar2093/charly_team_api/internal/validators"
	"github.com/manicar2093/charly_team_api/internal/validators/nullsql"
)

type BiotestImagesSaver interface {
	Run(ctx context.Context, biotestImages *BiotestImagesSaverRequest) (*BiotestImagesSaverResponse, error)
}

type biotestImagesSaverImpl struct {
	biotestRepo repositories.BiotestRepository
	validator   validators.ValidatorService
}

func NewBiotestImagesSaverImpl(biotestRepo repositories.BiotestRepository, validator validators.ValidatorService) *biotestImagesSaverImpl {
	return &biotestImagesSaverImpl{biotestRepo: biotestRepo, validator: validator}
}

func (c *biotestImagesSaverImpl) Run(ctx context.Context, biotestImages *BiotestImagesSaverRequest) (*BiotestImagesSaverResponse, error) {
	logger.Info(biotestImages)
	validation := c.validator.Validate(biotestImages)

	if !validation.IsValid {
		logger.Error(validation.Err)
		return nil, validation.Err
	}

	biotest, err := c.biotestRepo.FindBiotestByUUID(ctx, biotestImages.BiotestUUID)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	biotest.FrontPicture = nullsql.ValidateStringSQLValid(biotestImages.FrontPicture)
	biotest.BackPicture = nullsql.ValidateStringSQLValid(biotestImages.BackPicture)
	biotest.LeftSidePicture = nullsql.ValidateStringSQLValid(biotestImages.LeftSidePicture)
	biotest.RightSidePicture = nullsql.ValidateStringSQLValid(biotestImages.RightSidePicture)

	if err := c.biotestRepo.UpdateBiotest(ctx, biotest); err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info(biotestImages)

	return &BiotestImagesSaverResponse{BiotestImagesSaved: biotestImages}, nil
}