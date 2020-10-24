package dto

import "github.com/google/uuid"

type UserDto struct {
	Username string    `json:"username"`
	Password string    `json:"password,omitempty"`
	UID      uuid.UUID `json:"uid"`
}
