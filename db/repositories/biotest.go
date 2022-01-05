package repositories

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/paginator"
)

type BiotestRepositoryRel struct {
	repo      rel.Repository
	paginator paginator.Paginable
}

func NewBiotestRepositoryRel(repo rel.Repository, paginator paginator.Paginable) *BiotestRepositoryRel {
	return &BiotestRepositoryRel{repo: repo, paginator: paginator}
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
	biotestsFoundHolder := []entities.Biotest{}
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
	userFound, err := c.findUser(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	biotestsDetails := []BiotestDetails{}
	c.repo.FindAll(
		ctx,
		&biotestsDetails,
		where.Eq("customer_id", userFound.ID),
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
		sort.Asc("created_at"),
	)

	if len(biotestsDetails) == 0 {
		return nil, NotFoundError{Entity: "BiotestComparitionData", Identifier: userUUID}
	}

	firstBiotest := entities.Biotest{}
	c.repo.Find(
		ctx,
		&firstBiotest,
		where.Eq("customer_id", userFound.ID),
		sort.Asc("created_at"),
	)
	lastBiotest := entities.Biotest{}
	err = c.repo.Find(
		ctx,
		&lastBiotest,
		where.Eq("customer_id", userFound.ID),
		sort.Desc("created_at"),
	)

	if err != nil {
		switch err.(type) {
		case rel.NotFoundError:
			return &BiotestComparisionResponse{
				FirstBiotest:       &firstBiotest,
				AllBiotestsDetails: &biotestsDetails,
			}, nil
		}
		return nil, err
	}

	return &BiotestComparisionResponse{
		FirstBiotest:       &firstBiotest,
		LastBiotest:        &lastBiotest,
		AllBiotestsDetails: &biotestsDetails,
	}, nil
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