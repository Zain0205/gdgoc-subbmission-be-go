package controllers

import (
	"net/http"
	"strings"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateSubmission(c *gin.Context) {
	var input dto.CreateSubmissionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userID, _ := c.Get("userID")

	submission := models.Submission{
		UserID:   userID.(uint),
		SeriesID: uint(input.SeriesID),
		FileURL:  input.FileURL,
	}

	if err := database.DB.Create(&submission).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			utils.APIResponse(c, http.StatusBadRequest, "You have already submitted for this series", nil)
			return
		}
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create submission", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusCreated, "Submission created successfully", submission)
}

func UpdateSubmission(c *gin.Context) {
	var input dto.CreateSubmissionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userID, _ := c.Get("userID")

	var submission models.Submission
	if err := database.DB.Where("user_id = ? AND series_id = ?", userID, input.SeriesID).First(&submission).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Submission not found", nil)
		return
	}

	submission.FileURL = input.FileURL

	if err := database.DB.Save(&submission).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to update submission", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusOK, "Submission updated successfully", submission)
}

func GetSubmissionsBySeries(c *gin.Context) {
	seriesID := c.Param("seriesId")

	var submissions []models.Submission
	if err := database.DB.Preload("User").Where("series_id = ?", seriesID).Find(&submissions).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to fetch submissions", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusOK, "Submissions fetched successfully", submissions)
}

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
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to grade submission", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusOK, "Submission graded successfully", submission)
}

