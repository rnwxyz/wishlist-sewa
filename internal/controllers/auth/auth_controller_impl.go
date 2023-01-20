package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rnwxyz/wishlist-sewa/dto/requests"
	"github.com/rnwxyz/wishlist-sewa/internal/services/auth"
)

type authControllerImpl struct {
	authService auth.AuthService
}

// Login implements AuthController
func (d *authControllerImpl) Login(c echo.Context) error {
	var credential requests.Credential
	if err := c.Bind(&credential); err != nil {
		return err
	}
	if err := c.Validate(credential); err != nil {
		return err
	}
	token, err := d.authService.Login(&credential, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"data":    token,
	})
}

// UserRegister implements AuthController
func (d *authControllerImpl) UserRegister(c echo.Context) error {
	var user requests.UserRegister
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}
	if err := d.authService.UserRegister(&user, c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "register success",
	})
}

func NewAuthController(authService auth.AuthService) AuthController {
	return &authControllerImpl{
		authService: authService,
	}
}
