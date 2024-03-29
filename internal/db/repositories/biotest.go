package repositories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/db/paginator"
	"github.com/manicar2093/health_records/internal/services"
)

type BiotestRepositoryRel struct {
	repo      rel.Repository
	paginator paginator.Paginable
	uuidGen   services.UUIDGenerator
}

func NewBiotestRepositoryRel(
	repo rel.Repository,
	paginator paginator.Paginable,
	uuidGen services.UUIDGenerator,
) BiotestRepository {
	return &BiotestRepositoryRel{
		repo:      repo,
		paginator: paginator,
		uuidGen:   uuidGen,
	}
}

func (c *BiotestRepositoryRel) FindBiotestByUUID(
	ctx context.Context,
	biotestUUID string,
) (*entities.Biotest, error) {
	biotest := entities.Biotest{}
	if err := c.repo.Find(ctx, &biotest, where.Eq("biotest_uuid", biotestUUID)); err != nil {
		switch err.(type) {
		case rel.NotFoundError:
			return nil, NotFoundError{Entity: "Biotest", Identifier: biotestUUID}
		}
		return nil, err
	}
	return &biotest, nil
}

func (c *BiotestRepositoryRel) GetAllUserBiotestByUserUUID(
	ctx context.Context,
	pageSort *paginator.PageSort,
	userUUID string,
) (*paginator.Paginator, error) {
	// TODO: Inject userRepo instance to do this.
	userFound, err := c.findUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	biotestsFoundHolder := []entities.Biotest{}
	pageSort.SetFiltersQueries(
		where.Eq("customer_id", userFound.ID),
		sort.Asc("created_at"),
	)
	return c.paginator.CreatePagination(
		ctx,
		entities.BiotestTable,
		&biotestsFoundHolder,
		pageSort,
	)

}

func (c *BiotestRepositoryRel) GetAllUserBiotestByUserUUIDAsCatalog(
	ctx context.Context,
	pageSort *paginator.PageSort,
	userUUID string,
) (*paginator.Paginator, error) {
	// TODO: Inject userRepo instance to do this.

	userFound, err := c.findUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	biotestsFoundHolder := []BiotestDetails{}
	pageSort.SetFiltersQueries(
		where.Eq("customer_id", userFound.ID),
		sort.Asc("created_at"),
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	)
	return c.paginator.CreatePagination(
		ctx,
		entities.BiotestTable,
		&biotestsFoundHolder,
		pageSort,
	)
}

func (c *BiotestRepositoryRel) GetComparitionDataByUserUUID(
	ctx context.Context,
	userUUID string,
) (*BiotestComparisionResponse, error) {
	var (
		biotestsDetails []BiotestDetails
		firstBiotest    entities.Biotest
		lastBiotest     entities.Biotest
	)
	userFound, err := c.findUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	c.repo.FindAll(
		ctx,
		&biotestsDetails,
		where.Eq("customer_id", userFound.ID),
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
		sort.Asc("created_at"),
	)

	switch len(biotestsDetails) {
	case 0:
		return nil, NotFoundError{Entity: "BiotestComparitionData", Identifier: userUUID}
	case 1:
		c.repo.Find(
			ctx,
			&firstBiotest,
			where.Eq("customer_id", userFound.ID),
			sort.Asc("created_at"),
		)
		return &BiotestComparisionResponse{
			FirstBiotest:       &firstBiotest,
			LastBiotest:        nil,
			AllBiotestsDetails: &biotestsDetails,
		}, nil
	default:
		c.repo.Find(
			ctx,
			&firstBiotest,
			where.Eq("customer_id", userFound.ID),
			sort.Asc("created_at"),
		)
		c.repo.Find(
			ctx,
			&lastBiotest,
			where.Eq("customer_id", userFound.ID),
			sort.Desc("created_at"),
		)
		return &BiotestComparisionResponse{
			FirstBiotest:       &firstBiotest,
			LastBiotest:        &lastBiotest,
			AllBiotestsDetails: &biotestsDetails,
		}, nil

	}
}

func (c *BiotestRepositoryRel) SaveBiotest(
	ctx context.Context,
	biotest *entities.Biotest,
) error {
	err := c.repo.Transaction(ctx, func(ctx context.Context) error {
		biotest.BiotestUUID = c.uuidGen.New()

		if err := c.repo.Insert(ctx, biotest); err != nil {
			return err
		}

		biotest.HigherMuscleDensity.BiotestID = &biotest.ID
		biotest.LowerMuscleDensity.BiotestID = &biotest.ID
		biotest.SkinFolds.BiotestID = &biotest.ID

		if err := c.repo.Insert(ctx, &biotest.HigherMuscleDensity); err != nil {
			return err
		}

		if err := c.repo.Insert(ctx, &biotest.LowerMuscleDensity); err != nil {
			return err
		}

		if err := c.repo.Insert(ctx, &biotest.SkinFolds); err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		biotest.BiotestUUID = ""
	}

	return err
}

func (c *BiotestRepositoryRel) UpdateBiotest(
	ctx context.Context,
	biotest *entities.Biotest,
) error {
	return c.repo.Transaction(ctx, func(ctx context.Context) error {
		if err := c.repo.Update(ctx, &biotest.HigherMuscleDensity); err != nil {
			return err
		}

		if err := c.repo.Update(ctx, &biotest.LowerMuscleDensity); err != nil {
			return err
		}

		if err := c.repo.Update(ctx, &biotest.SkinFolds); err != nil {
			return err
		}
		if err := c.repo.Update(ctx, biotest); err != nil {
			return err
		}
		return nil
	})
}

func (c *BiotestRepositoryRel) findUser(ctx context.Context, userUUID string) (*entities.User, error) {
	userFound := entities.User{}

	if err := c.repo.Find(ctx, &userFound, where.Eq("user_uuid", userUUID)); err != nil {
		switch err.(type) {
		case rel.NotFoundError:
			return nil, NotFoundError{Entity: "User", Identifier: userUUID}

		}
		return nil, err
	}
	return &userFound, nil
}
