package dicontainer

import (
	"go_kafka_playground/api-gateway/src/core/interfaces/usecases"
	"go_kafka_playground/api-gateway/src/core/services"
	"go_kafka_playground/api-gateway/src/infra/repository/postgres"
)

func UserUseCase() usecases.UserUseCase {
	repo := postgres.NewUserRepository()
	return services.NewUserService(repo)
}
