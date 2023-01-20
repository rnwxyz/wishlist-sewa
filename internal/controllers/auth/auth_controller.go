package auth

import "github.com/labstack/echo/v4"

type AuthController interface {
	Login(c echo.Context) error
	UserRegister(c echo.Context) error
}
