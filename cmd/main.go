package main

import (
	"github.com/elizielx/arcturus-api/config"
	"github.com/elizielx/arcturus-api/db"
	"github.com/elizielx/arcturus-api/internal/routes"
	"github.com/labstack/echo/v4"
)

var (
	server *echo.Echo
)

func main() {
	configuration, err := config.LoadConfiguration(".")
	if err != nil {
		panic(err)
	}

	db.InitDatabase(configuration)

	server = echo.New()
	routes.SetupAuthRoutes(server)
	routes.SetupUserRoutes(server)
	routes.SetupPollRoutes(server)
	server.Logger.Fatal(server.Start(":" + configuration.ServerPort))
}
