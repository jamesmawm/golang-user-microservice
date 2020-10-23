package data

import (
	"github.com/jamesmawm/golang-user-microservice/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func GetDatabase() *gorm.DB {
	if database != nil {
		return database
	}

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}

	database = db

	return db
}
