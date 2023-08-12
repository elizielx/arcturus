package routes

import "github.com/labstack/echo/v4"

func SetupPollRoutes(e *echo.Echo) {
	e.GET("/polls", getPolls)
	e.GET("/polls/:id", getPoll)
	e.POST("/polls", createNewPoll)
}

func getPolls(c echo.Context) error {
	return nil
}

func getPoll(c echo.Context) error {
	return nil
}

func createNewPoll(c echo.Context) error {
	return nil
}
