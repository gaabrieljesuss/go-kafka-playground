package routes

import (
	"github.com/labstack/echo/v4"
	"go_kafka_playground/api-gateway/src/apps/api/dicontainer"
	"go_kafka_playground/api-gateway/src/apps/api/handlers"
)

type userRouter struct {
	handler handlers.UserHandler
}

func NewAccountRouter() Router {
	service := dicontainer.UserUseCase()
	handler := handlers.NewUserHandler(service)
	return &userRouter{handler}
}

func (a *userRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/admin/user")
	router.POST("", a.handler.Create)
}
