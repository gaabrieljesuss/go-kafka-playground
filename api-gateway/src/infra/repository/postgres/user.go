package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	"go_kafka_playground/api-gateway/src/core/domain/user"
	"go_kafka_playground/api-gateway/src/core/interfaces/adapters"
	"go_kafka_playground/api-gateway/src/infra"
)

var logger = infra.Logger()

type userRepository struct {
}

func NewUserRepository() adapters.UserAdapter {
	return &userRepository{}
}

func (r *userRepository) Create(user user.User) (*uuid.UUID, errors.Error) {
	// For now, the logic to persist the user in the database will not be implemented

	randomUuid, err := uuid.NewRandom()
	if err != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred generating the UUID: %s", err.Error()))
		return nil, errors.NewUnexpected()
	}

	return &randomUuid, nil
}
