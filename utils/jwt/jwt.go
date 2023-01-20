package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/rnwxyz/wishlist-sewa/constans"
	"github.com/rnwxyz/wishlist-sewa/model"
)

type JWTService interface {
	GenerateToken(user *model.User) (at string, err error)
	GetClaims(c *echo.Context) jwt.MapClaims
}

type jwtServiceImpl struct {
	secretKey string
}

func NewJWTService(secretKey string) JWTService {
	return &jwtServiceImpl{
		secretKey: secretKey,
	}
}

func (j *jwtServiceImpl) GenerateToken(user *model.User) (at string, err error) {
	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(constans.ExpAccessToken).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return result, nil
}

func (j *jwtServiceImpl) GetClaims(c *echo.Context) jwt.MapClaims {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
