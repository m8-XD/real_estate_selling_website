package repository

import (
	"github.com/m8-XD/real_estate_selling_website/db"
	"github.com/m8-XD/real_estate_selling_website/db/models"
	"github.com/m8-XD/real_estate_selling_website/logging"
)

func CreateUser(user *models.User) (uint, error) {
	db := db.Get()
	result := db.Create(user)
	if result.Error != nil {
		logging.Error("cannot crete a user: %v", *user)
		return 0, result.Error
	}
	return user.ID, nil
}
