package controllers

import (
	"errors"
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateSubmission
func CreateSubmission(c *gin.Context) {
	var input dto.CreateSubmissionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	memberID, _ := c.Get("userID")

	var verification models.UserSeriesVerification
	err := database.DB.Where("series_id = ? AND user_id = ?", input.SeriesID, memberID).First(&verification).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.APIResponse(c, http.StatusForbidden, "You must verify your attendance for this series before submitting", nil)
		return
	} else if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error checking verification", err.Error())
		return
	}

	submission := models.Submission{
		SeriesID: input.SeriesID,
		FileURL:  input.FileURL,
		UserID:   memberID.(uint),
		Score:    0, // Default score
	}

	if err := database.DB.Create(&submission).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create submission", err.Error())
		return
	}

	database.DB.Preload("User").Preload("Series").First(&submission, submission.ID)
	utils.APIResponse(c, http.StatusCreated, "Submission created successfully", submission)
}

// GetSubmissionsBySeries
func GetSubmissionsBySeries(c *gin.Context) {
	seriesID := c.Param("seriesId")

	var submissions []models.Submission
	err := database.DB.Preload("User").Where("series_id = ?", seriesID).Find(&submissions).Error
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to fetch submissions", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusOK, "Submissions fetched successfully", submissions)
}

