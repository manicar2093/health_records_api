package biotestfilters

import (
	"fmt"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/sort"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/filters"
)

func GetBiotestComparision(params *filters.FilterParameters) (interface{}, error) {
	valuesAsMap := params.Values.(map[string]interface{})

	userUUID, ok := valuesAsMap["user_uuid"].(string)
	if !ok {
		return nil, apperrors.ValidationError{Field: "user_uuid", Validation: "required"}
	}

	var userFound entities.User

	err := params.Repo.Find(params.Ctx, &userFound, where.Eq("user_uuid", userUUID))

	if err != nil {
		if _, ok := err.(rel.NotFoundError); ok {
			return nil, apperrors.NotFoundError{
				Message: fmt.Sprintf("User with uuid '%s' does not exist", userUUID),
			}
		}
		return nil, err
	}

	userID := userFound.ID

	var comparisionData BiotestComparisionResponse

	err = params.Repo.FindAll(
		params.Ctx,
		&comparisionData.AllBiotestsDetails,
		rel.Select("biotest_uuid", "created_at").From(entities.BiotestTable),
	)

	if err != nil {
		return nil, err
	}

	if len(comparisionData.AllBiotestsDetails) == 0 {
		return nil, apperrors.NotFoundError{
			Message: "User has no biotests",
		}
	}

	err = params.Repo.Find(
		params.Ctx,
		&comparisionData.FirstBiotest,
		where.Eq("customer_id", userID),
		sort.Asc("created_at"),
	)

	if err != nil {
		return nil, err
	}

	err = params.Repo.Find(
		params.Ctx,
		&comparisionData.LastBiotest,
		where.Eq("customer_id", userID),
		sort.Desc("created_at"),
	)

	if _, ok := err.(rel.NotFoundError); err != nil && !ok {
		return nil, err
	}

	return comparisionData, nil

}