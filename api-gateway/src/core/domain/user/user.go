package user

import (
	"github.com/google/uuid"
)

var _ User = &user{}

type user struct {
	id       *uuid.UUID
	name     string
	email    string
	password string
}

func (u *user) ID() *uuid.UUID {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Email() string {
	return u.email
}

func (u *user) Password() string {
	return u.password
}
