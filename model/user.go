package model

import (
	"github.com/go-gorm/gorm"
	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	Username string
	Password string
	UID      uuid.UUID
}
