package routes

import (
	"chapter1/config"
	"chapter1/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, ah users.UHandler) {
	e.POST("/login", ah.Login())
	e.POST("/register", ah.Register())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(config.ImportSetting().JWTSecret),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
