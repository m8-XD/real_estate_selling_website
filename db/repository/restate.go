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
	db := db.Get()
	err := db.Table("post").Create(post).Error
	if err != nil {
		return 0, err
	}
	return post.ID, nil
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

func HouseTypeId(t string) string {
	db := db.Get()
	var id uint
	db.Table("types").Select("id").Where("name = ?", t).First(&id)
	return fmt.Sprint(id)
}

func AllPostsByFilter(filter map[string][]string, sortBy string, page int) []models.PostPreview {
	db := db.Get()
	ok := validatePostFilter(filter)
	if !ok {
		logging.Info("someone passed illegal values to a filter, returning default filter")
		filter = make(map[string][]string, 0)
	}

	query := db.Table("post").Session(&gorm.Session{}).Select("post.header", "post.price", "types.name", "post.rooms")
	populateConditionQuery(query, filter)

	query.Order(sortBy)
	query.Offset((page - 1) * pageSize)
	query.Limit(pageSize)
	query.Joins("JOIN types ON post.type = types.id")

	var posts []models.PostPreview

	query.Scan(&posts)

	logging.Info("fetching posts for main page")
	logging.Info("%v", posts)

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
