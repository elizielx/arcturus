package routes

import (
	"github.com/elizielx/arcturus-api/db"
	"github.com/elizielx/arcturus-api/internal/api"
	"github.com/elizielx/arcturus-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupUserRoutes(e *echo.Echo) {
	e.GET("/me", me, api.IsAuthenticated)
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
