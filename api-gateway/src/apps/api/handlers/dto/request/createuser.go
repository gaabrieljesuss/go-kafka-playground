package request

import (
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	user2 "go_kafka_playground/api-gateway/src/core/domain/user"
)

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *CreateUser) ToDomain() (user2.User, errors.Error) {
	userBuilder := user2.NewBuilder()

	userBuilder.WithName(c.Name).
		WithEmail(c.Email).
		WithPassword(c.Password)

	_user, err := userBuilder.Build()
	if err != nil {
		return nil, err
	}

	return _user, nil
}
