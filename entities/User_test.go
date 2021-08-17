package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomerEntity(t *testing.T) {

	user := User{
		Biotype:     Biotype{ID: 1},
		BoneDensity: BoneDensity{ID: 1},
		Role:        Role{ID: 1},
		Name:        "Test",
		LastName:    "Test",
		Email:       "test@test.com",
		Password:    "12345678",
		Birthday:    time.Now(),
	}

	DB.Create(&user)

	assert.NotEmpty(t, user.ID, "ID should not be empty. Customer was not created")

	DB.Delete(&user)

}
