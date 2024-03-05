package models

import "gorm.io/gorm"

type REstate struct {
	gorm.Model
	Header      string
	Description string
	OwnerName   string
	OwnerPhone  string
	CreatorId   uint
}

type User struct {
	gorm.Model
	Name string
	Mail string
	Pass string
}
