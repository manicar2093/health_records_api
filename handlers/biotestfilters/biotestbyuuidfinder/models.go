package biotestbyuuidfinder

import "github.com/manicar2093/charly_team_api/db/entities"

type BiotestByUUIDRequest struct {
	UUID string `validate:"required"`
}

type BiotestByUUIDResponse struct {
	Biotest *entities.Biotest
}