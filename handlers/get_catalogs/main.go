package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/models"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	lambda.Start(
		CreateLambdaHandlerWDependencies(
			repositories.NewCatalogRepositoryGorm(
				connections.PostgressConnection(),
			),
			validators.NewStructValidator(),
		),
	)

}

func CreateLambdaHandlerWDependencies(
	catalogsRepository repositories.CatalogRepository,
	validator validators.ValidatorService,
) func(catalogs models.GetCatalogsRequest) *models.Response {

	return func(catalogs models.GetCatalogsRequest) *models.Response {

		isValid, response := validators.CheckValidationErrors(validator.Validate(catalogs))
		if !isValid {
			return response
		}

		gotCatalogs, err := CatalogFactoryLoop(catalogs, catalogsRepository)

		if err != nil {
			return validators.CreateResponseError(err)
		}

		return models.CreateResponse(http.StatusOK, gotCatalogs)

	}

}
