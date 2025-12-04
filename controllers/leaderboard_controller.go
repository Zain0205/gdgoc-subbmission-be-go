package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

type LeaderboardResult struct {
	UserID     uint    `json:"user_id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	TotalScore float64 `json:"total_score"`
	Rank       int     `json:"rank"`
}

func GetLeaderboardByTrack(c *gin.Context) {
	trackID := c.Param("trackId")

	var results []LeaderboardResult
	err := database.DB.Table("submissions").
		Select("users.id as user_id, users.name as name, users.email as email, SUM(submissions.score) as total_score").
		Joins("JOIN users ON users.id = submissions.user_id").
		Joins("JOIN series ON series.id = submissions.series_id").
		Where("series.track_id = ? AND series.is_competition = ?", trackID, true).
		Group("users.id, users.name, users.email").
		Order("total_score DESC").
		Scan(&results).Error
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to fetch leaderboard", err.Error())
		return
	}

	for i := range results {
		results[i].Rank = i + 1
	}

	utils.APIResponse(c, http.StatusOK, "Leaderboard fetched successfully", results)
}

