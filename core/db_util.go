package core

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type UserModel struct {
	gorm.Model
	Uuid           uuid.UUID
	Username       string `json:"username"`
	Password       string
	FirstName      string
	LastName       string
	Email          string `gorm:"index"`
	Salt           string
	HashedPassword string
	Ctime          int
	Mtime          int
	Status         int
}

func InitDb() (db *gorm.DB, err error) {
	dialect := "sqlite3"
	conn := "db.sqlite"

	db, err = gorm.Open(dialect, conn)
	if err != nil {
		return
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)

	_ = db.AutoMigrate(User{})
	db.CreateTable(&User{})
	return
}
