package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

// GradeSubmission (Admin)
func GradeSubmission(c *gin.Context) {
	var input dto.CreateScoreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var submission models.Submission
	if err := database.DB.First(&submission, input.SubmissionID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Submission not found", nil)
		return
	}

	submission.Score = input.Score
	submission.Feedback = input.Feedback

	if err := database.DB.Save(&submission).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to update submission score", err.Error())
		return
	}

	database.DB.Preload("User").Preload("Series").First(&submission, submission.ID)
	utils.APIResponse(c, http.StatusOK, "Submission graded successfully", submission)
}
