package models

import (
	"regexp"
	"strings"

	"github.com/m8-XD/real_estate_selling_website/logging"
	"gorm.io/gorm"
)

type REstate struct {
	gorm.Model
	Header      string
	Description string
	Name        string
	Phone       string
	Price       int64
	Type        uint
	Rooms       string
	CreatorId   uint
}

type User struct {
	gorm.Model
	Name string
	Mail string
	Pass string
}

type Types struct {
	gorm.Model
	Name string
}

type PostPreview struct {
	Header string
	Price  int64
	Type   uint
	Rooms  string
}

func (r REstate) Validate() bool {
	// Phone       string
	ok, err := regexp.Match(`\+375\(00\)000-00-00`, []byte(r.Phone))
	if !ok {
		logging.Info("user entered illegal phone number")
		return false
	}
	if err != nil {
		logging.Info("pattern match returned an error while validating a phone")
		logging.Info(err.Error())
		return false
	}
	// Price       int64
	if r.Price < 0 {
		logging.Info("user entered negative price")
		return false
	}
	// Rooms       string
	if !(strings.EqualFold(r.Rooms, "1") || strings.EqualFold(r.Rooms, "2") ||
		strings.EqualFold(r.Rooms, "3") || strings.EqualFold(r.Rooms, "4") ||
		strings.EqualFold(r.Rooms, "5+")) {
		logging.Info("user entered wrong room number")
		return false
	}
	return true
}
