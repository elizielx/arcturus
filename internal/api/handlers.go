package api

import (
	"github.com/elizielx/arcturus-api/config"
	"github.com/elizielx/arcturus-api/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"strings"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		configuration, err := config.LoadConfiguration(".")
		if err != nil {
			panic(err)
		}

		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.ErrUnauthorized
		}

		token, err := jwt.ParseWithClaims(parts[1], &utils.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(configuration.JWTSecret), nil
		})

		if err != nil {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*utils.JwtClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		c.Set("username", claims.Username)
		return next(c)
	}
}
