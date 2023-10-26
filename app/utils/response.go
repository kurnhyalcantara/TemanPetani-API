package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseJSON struct{}

func NewResponseJSON() *ResponseJSON {
	return &ResponseJSON{}
}

func (res *ResponseJSON) StatusOK(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": message,
	})
}

func (res *ResponseJSON) StatusOKWithData(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func (res *ResponseJSON) StatusCreated(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func (res *ResponseJSON) StatusBadRequestResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]any{
		"status":  "fail",
		"message": message,
	})
}

func (res *ResponseJSON) StatusNotFoundResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"status":  "fail",
		"message": message,
	})
}

func (res *ResponseJSON) StatusAuthorizationErrorResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusUnauthorized, map[string]any{
		"status":  "fail",
		"message": message,
	})
}

func (res *ResponseJSON) StatusForbiddenResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusForbidden, map[string]any{
		"status":  "fail",
		"message": message,
	})
}

func (res *ResponseJSON) StatusInternalServerError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, map[string]any{
		"status":  "fail",
		"message": "Terjadi kesalahan di server kami: " + message,
	})
}
