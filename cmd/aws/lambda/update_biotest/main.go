package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/manicar2093/charly_team_api/db/connections"
	"github.com/manicar2093/charly_team_api/internal/config"
	"github.com/manicar2093/charly_team_api/internal/handlers/biotestupdater"
	"github.com/manicar2093/charly_team_api/pkg/validators"
)

func main() {
	config.StartConfig()
	service := biotestupdater.NewBiotestUpdater(
		connections.PostgressConnection(),
		validators.NewStructValidator(),
	)
	lambda.Start(
		NewUpdateBiotestAWSLambda(service).Handler,
	)
}
