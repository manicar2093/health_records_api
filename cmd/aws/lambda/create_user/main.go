package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/aws"
	"github.com/manicar2093/charly_team_api/config"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/db/paginator"
	"github.com/manicar2093/charly_team_api/db/repositories"
	"github.com/manicar2093/charly_team_api/handlers/usercreator"
	"github.com/manicar2093/charly_team_api/services"
	"github.com/manicar2093/charly_team_api/validators"
)

func main() {
	config.StartConfig()
	conn := connections.PostgressConnection()
	paginator := paginator.NewPaginable(conn)
	service := usercreator.NewUserCreatorImpl(
		aws.NewCognitoClient(),
		services.PasswordGenerator{},
		repositories.NewUserRepositoryRel(conn, paginator),
		services.UUIDGeneratorImpl{},
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewUserCreatorAWSLambda(service).Handler,
	)
}