package repository

import (
	"github.com/m8-XD/real_estate_selling_website/db"
	"github.com/m8-XD/real_estate_selling_website/db/models"
	"github.com/m8-XD/real_estate_selling_website/logging"
)

func CreatePost(post *models.REstate) (uint, error) {
	return 0, nil
}

func HouseTypes() []string {
	db := db.Get()
	var types []string
	err := db.Table("types").Select("name").Find(&types).Error
	if err != nil {
		logging.Error(err.Error())
		logging.Error("can't get house types from db (It shouldn't be like that)")
	}
	return types
}

func HouseTypeId(t string) uint {
	db := db.Get()
	var id uint
	db.Table("types").Select("name = ?", t).First(&id)
	return id
}
