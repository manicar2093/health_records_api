package main

import (
	"context"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/apperrors"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/db/repositories"
)

// FilterFunc represents a result getter
type FilterFunc func(
	ctx context.Context,
	repo rel.Repository,
	values interface{},
	paginator repositories.Paginable,
) (interface{}, error)

var userFilterRegistered = make(map[string]FilterFunc)

func init() {
	userFilterRegistered["find_all_users"] = FindAllUsers
	userFilterRegistered["find_user_by_email"] = FindUserByEmail
	userFilterRegistered["finde_user_by_id"] = FindUserByID
}

func FindUserByID(
	ctx context.Context,
	repo rel.Repository,
	values interface{},
	paginator repositories.Paginable,
) (interface{}, error) {

	valuesAsMap := values.(map[string]interface{})

	userID, ok := valuesAsMap["user_id"].(int)
	if !ok {
		return nil, apperrors.ValidationError{Field: "user_id", Validation: "required"}
	}

	var userFound entities.User

	err := repo.Find(ctx, &userFound, where.Eq("id", userID))
	if err != nil {
		return nil, err
	}
	return userFound, nil
}

func FindUserByEmail(
	ctx context.Context,
	repo rel.Repository,
	values interface{},
	paginator repositories.Paginable,
) (interface{}, error) {

	valuesAsMap := values.(map[string]interface{})

	userEmail, ok := valuesAsMap["email"].(string)
	if !ok {
		return nil, apperrors.ValidationError{Field: "email", Validation: "required"}
	}

	var userFound entities.User

	err := repo.Find(ctx, &userFound, where.Like("email", "%"+userEmail+"%"))

	if err != nil {
		if _, ok := err.(rel.NotFoundError); ok {
			return nil, &apperrors.UserNotFound{}
		}
		return nil, err
	}

	return userFound, nil

}

func FindAllUsers(
	ctx context.Context,
	repo rel.Repository,
	values interface{},
	paginator repositories.Paginable,
) (interface{}, error) {

	valuesAsMap := values.(map[string]interface{})

	pageNumber, ok := valuesAsMap["page_number"].(int)
	if !ok {
		return nil, apperrors.ValidationError{Field: "page_number", Validation: "required"}
	}

	var usersFound []entities.User

	return paginator.CreatePaginator(
		ctx,
		entities.User{}.Table(),
		&usersFound,
		pageNumber,
	)

}
