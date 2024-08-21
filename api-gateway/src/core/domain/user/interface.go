package user

import (
	"github.com/google/uuid"
)

type User interface {
	ID() *uuid.UUID
	Name() string
	Email() string
	Password() string
}
