package repository

import (
	"fmt"
	"strings"

	"github.com/m8-XD/real_estate_selling_website/db"
	"github.com/m8-XD/real_estate_selling_website/db/models"
	"github.com/m8-XD/real_estate_selling_website/logging"
	"gorm.io/gorm"
)

const (
	pageSize int = 5
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

func AllPostsByFilter(filter map[string][]string, sortBy string, page int) []models.PostPreview {
	db := db.Get()
	ok := validatePostFilter(filter)
	if !ok {
		logging.Info("someone passed illegal values to a filter, returning default filter")
		filter = make(map[string][]string, 0)
	}
	query := db.Table("post").Select("header", "price", "type", "rooms").Session(&gorm.Session{})
	populateConditionQuery(query, filter)
	query.Order(sortBy)
	query.Offset((page - 1) * pageSize)
	query.Limit(pageSize)
	var posts []models.PostPreview
	query.Find(&posts)

	return posts
}

func validatePostFilter(filter map[string][]string) bool {
	for param := range filter {
		if !(strings.EqualFold(param, "price") || strings.EqualFold(param, "type") || strings.EqualFold(param, "rooms")) {
			return false
		}
	}
	return true
}

func populateConditionQuery(query *gorm.DB, filter map[string][]string) {
	for key, values := range filter {
		if len(values) != 0 {
			query.Where(fmt.Sprintf("%s IN ?", key), values)
		}
	}
}
