package response

import (
	"go_kafka_playground/api-gateway/src/core/domain/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorMessage struct {
	error
	Code          int            `json:"status_code,omitempty"`
	Message       string         `json:"message"`
	InvalidFields []InvalidField `json:"invalid_fields,omitempty"`
}

type InvalidField struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}

type errorBuilder struct{}

var unprocessableEntityError = &echo.HTTPError{
	Code: http.StatusUnprocessableEntity,
}
var badRequestError = &echo.HTTPError{
	Code: http.StatusBadRequest,
}
var internalServerError = &echo.HTTPError{
	Code:    http.StatusInternalServerError,
	Message: "Unexpected error, please contact support",
}

func ErrorBuilder() *errorBuilder {
	return &errorBuilder{}
}

func (e *errorBuilder) NewFromDomain(err errors.Error) *echo.HTTPError {
	if err.CausedByValidation() {
		return e.unprocessableEntityErrorWithMessage(err.String())
	} else if err.CausedInternally() {
		return internalServerError
	} else if err.CausedByBadRequest() {
		return e.badRequestErrorWithMessage(err.String())
	}
	return &echo.HTTPError{
		Code:    badRequestError.Code,
		Message: err.String(),
	}
}

func (*errorBuilder) badRequestErrorWithMessage(message string) *echo.HTTPError {
	err := badRequestError
	err.Message = message
	return err
}

func (*errorBuilder) unprocessableEntityErrorWithMessage(message string) *echo.HTTPError {
	err := unprocessableEntityError
	err.Message = message
	return err
}

func (e *ErrorMessage) Error() echo.HTTPError {
	return echo.HTTPError{
		Message: e.Message,
		Code:    e.Code,
	}
}
