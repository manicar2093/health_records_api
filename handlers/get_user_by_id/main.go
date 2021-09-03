package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/entities"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			connections.PostgressConnection(),
			validators.NewStructValidator(),
		),
	)

}

func CreateLambdaHandlerWDependencies(
	repo rel.Repository,
	validator validators.ValidatorService,
) func(ctx context.Context, req GetUserByID) *models.Response {

	return func(ctx context.Context, req GetUserByID) *models.Response {

		isValid, response := validators.CheckValidationErrors(validator.Validate(req))
		if !isValid {
			return response
		}

		var user entities.User

		err := repo.Find(ctx, &user, where.Eq("id", req.UserID))
		if err != nil {
			if _, ok := err.(rel.NotFoundError); ok {
				return models.CreateResponse(http.StatusNotFound, models.ErrorReponse{Error: err.Error()})
			}
			return models.CreateResponseFromError(err)
		}

		return models.CreateResponse(http.StatusOK, user)

	}

}