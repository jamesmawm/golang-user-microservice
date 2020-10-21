package dto

import "github.com/google/uuid"

type UserDto struct {
	Username string
	UID      uuid.UUID
}
