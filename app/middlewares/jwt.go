package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kurnhyalcantara/TemanPetani-API/app/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.GetAppConfig().JWT_SECRET_KEY),
		ErrorHandler:  jwtError,
		SigningMethod: "HS256",
	})
}

func GenerateAccessToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetAppConfig().JWT_SECRET_KEY))
}

func GenerateRefreshToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetAppConfig().JWT_REFRESH_KEY)
}

func jwtError(c echo.Context, err error) error {
	if strings.Contains(err.Error(), "Missing or malformed JWT") {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"msg": err.Error(),
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]any{
		"msg": err.Error(),
	})
}
