package controller

import "github.com/labstack/echo/v4"

type UserControllerInterface interface {
	PostUserController(c echo.Context) error
}
