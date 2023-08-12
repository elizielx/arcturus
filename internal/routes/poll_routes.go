package routes

import (
	"encoding/json"
	"github.com/elizielx/arcturus-api/db"
	"github.com/elizielx/arcturus-api/internal/api"
	"github.com/elizielx/arcturus-api/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SetupPollRoutes(e *echo.Echo) {
	e.GET("/polls", getPolls, api.IsAuthenticated)
	e.GET("/polls/:poll_id", getPoll, api.IsAuthenticated)
	e.POST("/polls", createNewPoll, api.IsAuthenticated)
	e.POST("/polls/:poll_id/vote/:choice_id", voteOnPoll, api.IsAuthenticated)
	e.DELETE("/polls/:poll_id", deletePoll, api.IsAuthenticated)
}

func getPolls(c echo.Context) error {
	username := c.Get("username")

	var user models.User
	getUser := db.GetDatabase().Where("username = ?", username).First(&user)
	if getUser.Error != nil || getUser.RowsAffected == 0 {
		return echo.ErrUnauthorized
	}

	var polls []models.Poll

	if user.Role != models.ADMIN {
		result := db.GetDatabase().Where("created_by = ?", user.ID).Find(&polls)
		if result.Error != nil {
			return echo.ErrInternalServerError
		}
	} else {
		result := db.GetDatabase().Find(&polls)
		if result.Error != nil {
			return echo.ErrInternalServerError
		}
	}

	type PollResponse struct {
		ID          uint64    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Deadline    time.Time `json:"deadline"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CreatedBy   uint64    `json:"created_by"`
	}

	var pollResponses []PollResponse
	for _, poll := range polls {
		pollResponses = append(pollResponses, PollResponse{
			ID:          poll.ID,
			Title:       poll.Title,
			Description: poll.Description,
			Deadline:    poll.Deadline,
			CreatedAt:   poll.CreatedAt,
			UpdatedAt:   poll.UpdatedAt,
			CreatedBy:   poll.CreatedBy,
		})
	}

	response := make(map[string]interface{})
	response["polls"] = pollResponses

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONBlob(http.StatusOK, jsonResponse)
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
