package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func Get() *gorm.DB {
	if db == nil {
		return create()
	}
	return db
}

func create() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=restate port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("filed to connect to database, error : " + err.Error())
	}
	return db
}
