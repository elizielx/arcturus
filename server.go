package main

import (
	"github.com/elizielx/arcturus-api/config"
	"github.com/elizielx/arcturus-api/db"
	"github.com/elizielx/arcturus-api/models"
	"github.com/elizielx/arcturus-api/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

var (
	server *echo.Echo
)

func init() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		panic(err)
	}

	db.InitDatabase(configuration)

	server = echo.New()
}

func main() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		panic(err)
	}

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	server.POST("/login", login)
	server.GET("/me", me, isAuthenticated)
	server.Logger.Fatal(server.Start(":" + configuration.ServerPort))
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
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

func me(c echo.Context) error {
	var user models.User
	result := db.GetDatabase().Where("username = ?", c.Get("username")).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return echo.ErrUnauthorized
	}

	user.Password = ""

	return c.JSON(http.StatusOK, user)
}

func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
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

		token, err := jwt.ParseWithClaims(parts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(configuration.JWTSecret), nil
		})

		if err != nil {
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			return echo.ErrUnauthorized
		}

		c.Set("username", claims.Username)
		return next(c)
	}
}
