package utility

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ErrorResponse{
		Message: message,
	})
}

func NotFoundError(c echo.Context, message string) error {
	return NewErrorResponse(c, http.StatusNotFound, message)
}

func BadRequestError(c echo.Context, message string) error {
	return NewErrorResponse(c, http.StatusBadRequest, message)
}

func InternalServerError(c echo.Context, message string) error {
	return NewErrorResponse(c, http.StatusInternalServerError, message)
}
