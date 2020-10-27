package data

import (
	"log"
	"sync"

	"github.com/jamesmawm/golang-user-microservice/config"

	"github.com/jamesmawm/golang-user-microservice/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}
var database *gorm.DB

// GetDatabase checks if there is already an instance of Database and returns that or
// creates a new one
func GetDatabase() *gorm.DB {
	if database == nil {
		lock.Lock()
		defer lock.Unlock()
		if database == nil {
			database = newDatabase()
		}
	}

	return database
}

func newDatabase() *gorm.DB {
	dbConfig := config.App.Database

	var db *gorm.DB
	var err error

	switch dbConfig.Connection {
	case "mysql":
		dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ")/" + dbConfig.Name + "?charset=utf8&parseTime=True&loc=Local"
		log.Printf("Found configuration for connection type \"mysql\". Trying to connect to \"%s\".", dsn)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "sqlite":
		log.Printf("Found configuration for connection type \"sqlite\". Trying to open \"%s\".", dbConfig.Name)
		db, err = gorm.Open(sqlite.Open(dbConfig.Name), &gorm.Config{})
		break
	default:
		log.Printf("Failed to find a driver for the connection type \"%s\". As fallback an in-memory sqlite database is used.", dbConfig.Connection)
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}

	return db
}
