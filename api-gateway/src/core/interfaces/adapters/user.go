package adapters

import (
	"github.com/google/uuid"
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	"go_kafka_playground/api-gateway/src/core/domain/user"
)

type UserAdapter interface {
	Create(user user.User) (*uuid.UUID, errors.Error)
}
