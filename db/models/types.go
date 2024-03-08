package models

import "gorm.io/gorm"

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
