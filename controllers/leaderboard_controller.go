package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

func GetLeaderboardByTrack(c *gin.Context) {
	trackID := c.Param("trackId")

	var results []dto.LeaderboardResult

	err := database.DB.Table("users").
		Select("users.id as user_id, users.name, SUM(submissions.score) as total_score").
		Joins("JOIN submissions ON users.id = submissions.user_id").
		Joins("JOIN series ON submissions.series_id = series.id").
		Where("series.track_id = ?", trackID).
		// Only sum scores from series that are marked as a competition
		Where("series.is_competition = ?", true).
		Group("users.id, users.name").
		Order("total_score DESC").
		Scan(&results).Error
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to generate leaderboard", err.Error())
		return
	}

	for i := range results {
		results[i].Rank = i + 1
	}

	utils.APIResponse(c, http.StatusOK, "Leaderboard fetched successfully", results)
}

