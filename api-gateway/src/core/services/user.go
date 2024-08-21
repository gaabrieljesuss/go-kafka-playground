package services

import (
	"github.com/google/uuid"
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	"go_kafka_playground/api-gateway/src/core/domain/user"
	"go_kafka_playground/api-gateway/src/core/interfaces/adapters"
	"go_kafka_playground/api-gateway/src/core/interfaces/usecases"
	"go_kafka_playground/api-gateway/src/infra"
	"go_kafka_playground/api-gateway/src/infra/kafka"
)

var logger = infra.Logger()

type userService struct {
	adapter adapters.UserAdapter
}

func NewUserService(repository adapters.UserAdapter) usecases.UserUseCase {
	return &userService{repository}
}

func (s *userService) Create(user user.User) (*uuid.UUID, errors.Error) {
	// Since database persistence will not be implemented, some uniqueness validations will be ignored for now
	userId, err := s.adapter.Create(user)
	if err != nil {
		return nil, err
	}

	message := "User Created: " + userId.String()
	err = kafka.ProduceMessage([]string{"localhost:9092"}, "user-topic", message)
	if err != nil {
		return nil, err
	}

	return userId, nil
}
