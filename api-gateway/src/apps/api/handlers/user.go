package handlers

import (
	"github.com/labstack/echo/v4"
	"go_kafka_playground/api-gateway/src/apps/api/handlers/dto/request"
	"go_kafka_playground/api-gateway/src/apps/api/handlers/dto/response"
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	"go_kafka_playground/api-gateway/src/core/interfaces/usecases"
	"net/http"
)

type UserHandler interface {
	Create(echo.Context) error
}

type userHandler struct {
	service usecases.UserUseCase
}

func NewUserHandler(service usecases.UserUseCase) UserHandler {
	return &userHandler{service}
}

// Create
// @ID User.Create
// @Summary Create a new user account
// @Description Route responsible for creating a new user account
// @Security bearerAuth
// @Accept json
// @Param json body request.CreateUser true "JSON with all the data necessary to create a user account."
// @Tags Admin
// @Produce json
// @Success 201 {object} response.ID "Request made successfully."
// @Failure 400 {object} response.ErrorMessage "Poorly formulated request."
// @Failure 422 {object} response.ErrorMessage "Some data entered could not be processed correctly."
// @Failure 500 {object} response.ErrorMessage "An unexpected error has occurred. Please contact support."
// @Router /admin/user [post]
func (h *userHandler) Create(context echo.Context) error {

	var dto request.CreateUser
	bindError := context.Bind(&dto)
	if bindError != nil {
		return response.ErrorBuilder().NewFromDomain(errors.NewBadRequest("Invalid request information"))
	}

	if data, err := dto.ToDomain(); err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	} else if id, err := h.service.Create(data); err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	} else {
		return context.JSON(http.StatusCreated, map[string]interface{}{
			"id": id.String(),
		})
	}
}
