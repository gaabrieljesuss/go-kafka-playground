package user

import (
	"github.com/google/uuid"
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	validator2 "go_kafka_playground/api-gateway/src/utils/validator"
	"net/mail"
	"strings"
)

type builder struct {
	errorMessages []string
	fields        []string
	user          *user
}

func NewBuilder() *builder {
	return &builder{
		errorMessages: []string{},
		fields:        []string{},
		user:          &user{},
	}
}

func (b *builder) WithID(id *uuid.UUID) *builder {
	if !validator2.IsUUIDValid(*id) {
		b.errorMessages = append(b.errorMessages, "Invalid user id")
		b.fields = append(b.fields, "User ID")
	}

	b.user.id = id
	return b
}

func (b *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if validator2.IsTextBlank(name) {
		b.errorMessages = append(b.errorMessages, "The user name cannot be empty")
		b.fields = append(b.fields, "User name")
	}

	b.user.name = name
	return b
}

func (b *builder) WithEmail(email string) *builder {
	if addr, _ := mail.ParseAddress(email); addr == nil {
		b.errorMessages = append(b.errorMessages, "Invalid user email")
		b.fields = append(b.errorMessages, "User email")
	}

	b.user.email = email
	return b
}

func (b *builder) WithPassword(password string) *builder {
	password = strings.TrimSpace(password)
	if !validator2.IsPasswordValid(password) {
		b.errorMessages = append(b.errorMessages, "Invalid user password. The password must be between 8 and 50 characters")
		b.fields = append(b.fields, "User password")
	}

	b.user.password = password
	return b
}

func (b *builder) Build() (User, errors.Error) {
	if len(b.errorMessages) != 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{"Fields": b.fields})
	}

	return b.user, nil
}
