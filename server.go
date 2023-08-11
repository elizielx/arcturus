package main

import (
	"github.com/elizielx/arcturus-api/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":8080"))
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	user := models.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if (bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password))) == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			},
		})

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		response := map[string]interface{}{
			"access_token": tokenString,
			"expires_in":   time.Now().Add(time.Hour * 2).Unix(),
		}

		return c.JSON(http.StatusOK, response)
	}

	return echo.ErrUnauthorized
}
