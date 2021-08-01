package entities

import (
	"os"
	"testing"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB
var dns = "postgres://development:development@localhost:3456/charly_team_db_dev?sslmode=disable"

func TestMain(m *testing.M) {
	db, err := gorm.Open(postgres.Open(dns))
	if err != nil {
		panic(err)
	}
	DB = db
	os.Exit(m.Run())

}