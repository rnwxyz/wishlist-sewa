package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rnwxyz/wishlist-sewa/internal/services/users"
	jwtServ "github.com/rnwxyz/wishlist-sewa/utils/jwt"
)

type userControlleImpl struct {
	userService users.UserService
	jwtService  jwtServ.JWTService
}

// DeleteUser implements UserController
func (d *userControlleImpl) DeleteUser(c echo.Context) error {
	panic("unimplemented")
}

// GetProfile implements UserController
func (d *userControlleImpl) GetProfile(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	id := uint(claims["user_id"].(float64))
	fmt.Println(id)
	user, err := d.userService.FindUserByID(&id, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get profile",
		"data":    user,
	})
}

// GetUserByID implements UserController
func (d *userControlleImpl) GetUserByID(c echo.Context) error {
	panic("unimplemented")
}

// GetUsers implements UserController
func (d *userControlleImpl) GetUsers(c echo.Context) error {
	panic("unimplemented")
}

// UpdateOwner implements UserController
func (d *userControlleImpl) UpdateOwner(c echo.Context) error {
	panic("unimplemented")
}

// UpdateUser implements UserController
func (d *userControlleImpl) UpdateUser(c echo.Context) error {
	panic("unimplemented")
}

func NewUserController(userService users.UserService, jwtService jwtServ.JWTService) UserController {
	return &userControlleImpl{
		userService: userService,
		jwtService:  jwtService,
	}
}
