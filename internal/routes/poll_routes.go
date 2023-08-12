package routes

import "github.com/labstack/echo/v4"

func SetupPollRoutes(e *echo.Echo) {
	e.GET("/polls", getPolls)
	e.GET("/polls/:id", getPoll)
	e.POST("/polls", createNewPoll)
	e.POST("/polls/:poll_id/vote/:choice_id", voteOnPoll)
	e.DELETE("/polls/:id", deletePoll)
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

func voteOnPoll(c echo.Context) error {
	return nil
}

func deletePoll(c echo.Context) error {
	return nil
}
