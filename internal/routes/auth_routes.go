package routes

import (
	"github.com/elizielx/arcturus-api/config"
	"github.com/elizielx/arcturus-api/db"
	"github.com/elizielx/arcturus-api/internal/models"
	"github.com/elizielx/arcturus-api/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SetupAuthRoutes(e *echo.Echo) {
	e.POST("/login", login)
}

func login(c echo.Context) error {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		panic(err)
	}

	username := c.FormValue("username")
	rawPassword := c.FormValue("password")

	if username == "" || rawPassword == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Both username and password are required")
	}

	var user models.User
	result := db.GetDatabase().Where("username = ?", username).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return echo.ErrUnauthorized
	}

	if !utils.CheckPasswordHash(rawPassword, user.Password) {
		return echo.ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, utils.JwtClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(configuration.JWTSecret))
	if err != nil {
		return err
	}

	response := map[string]interface{}{
		"access_token": tokenString,
		"expires_in":   time.Now().Add(time.Hour * 2).Unix(),
	}

	return c.JSON(http.StatusOK, response)
}
