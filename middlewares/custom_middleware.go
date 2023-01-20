package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/rnwxyz/wishlist-sewa/constans"
	"github.com/rnwxyz/wishlist-sewa/utils/myerrors"
)

type customMiddleware struct {
	secret string
}

func (j *customMiddleware) JWTMiddleware(role string) echo.MiddlewareFunc {
	config := echojwt.Config{}
	config.SigningMethod = jwt.SigningMethodHS256.Alg()
	config.SigningKey = []byte(j.secret)
	config.KeyFunc = func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != config.SigningMethod {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		if len(config.SigningKeys) > 0 {
			if kid, ok := t.Header["kid"].(string); ok {
				if key, ok := config.SigningKeys[kid]; ok {
					return key, nil
				}
			}
			return nil, fmt.Errorf("unexpected jwt key id=%v", t.Header["kid"])
		}
		return config.SigningKey, nil
	}
	config.ParseTokenFunc = func(c echo.Context, auth string) (interface{}, error) {
		token, err := jwt.Parse(auth, config.KeyFunc)
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("invalid token")
		}
		claims := token.Claims.(jwt.MapClaims)
		if role != "" {
			userRole := claims["role"].(string)
			if userRole != role && userRole != constans.RoleOwner {
				return nil, myerrors.ErrPermission
			}
		}
		return token, nil
	}
	config.ErrorHandler = func(c echo.Context, err error) error {
		if err == myerrors.ErrPermission {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		if err.Error() == "code=400, message=missing or malformed jwt" {
			return err
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return echojwt.WithConfig(config)
}

func NewCustomMiddleware(secret string) customMiddleware {
	return customMiddleware{
		secret: secret,
	}
}
