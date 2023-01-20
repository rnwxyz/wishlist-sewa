package users

import "github.com/labstack/echo/v4"

type UserController interface {
	GetProfile(c echo.Context) error
	GetUserByID(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	UpdateOwner(c echo.Context) error
}
