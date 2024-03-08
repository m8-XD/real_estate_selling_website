package repository

import (
	"strings"

	"github.com/m8-XD/real_estate_selling_website/db"
	"github.com/m8-XD/real_estate_selling_website/db/models"
	"github.com/m8-XD/real_estate_selling_website/logging"
)

func CreateUser(user *models.User) error {
	db := db.Get()
	result := db.Create(user)
	if result.Error != nil {
		logging.Error("cannot crete a user: %v", *user)
		return result.Error
	}
	return nil
}

// returns true if mail is not presented in db
func IsValidMail(mail string) bool {
	db := db.Get()
	user := models.User{}
	result := db.Where("mail = ?", mail).First(&user)
	return result.Error != nil
}

func GetId(mail string, pass string) (bool, uint) {
	db := db.Get()
	user := models.User{}
	result := db.Model(&models.User{}).Where("mail = ?", mail).First(&user)
	if result.Error != nil {
		return false, 0
	}
	if !strings.EqualFold(user.Pass, pass) {
		return false, 0
	}
	return true, user.ID
}
