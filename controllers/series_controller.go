package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateSeries(c *gin.Context) {
	var input dto.CreateSeriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var track models.Track
	if err := database.DB.First(&track, input.TrackID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Track not found", err.Error())
		return
	}

	series := models.Series{
		TrackID:       input.TrackID,
		SeriesName:    input.SeriesName,
		Description:   input.Description,
		Deadline:      input.Deadline,
		OrderIndex:    input.OrderIndex,
		IsCompetition: input.IsCompetition,
	}

	if err := database.DB.Create(&series).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create series", err.Error())
		return
	}

	database.DB.Preload("Track").First(&series, series.ID)
	utils.APIResponse(c, http.StatusCreated, "Series created successfully", series)
}

// SetSeriesVerificationCode (Admin)
func SetSeriesVerificationCode(c *gin.Context) {
	seriesID := c.Param("id")

	var input dto.SetVerificationCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var series models.Series
	if err := database.DB.First(&series, seriesID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Series not found", nil)
		return
	}

	series.VerificationCode = input.Code
	if err := database.DB.Save(&series).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to set verification code", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusOK, "Verification code set successfully", gin.H{"series_id": series.ID, "code_set": true})
}

// VerifySeriesCode (Member)
func VerifySeriesCode(c *gin.Context) {
	seriesIDStr := c.Param("id")
	memberID, _ := c.Get("userID")

	var input dto.VerifyCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var series models.Series
	if err := database.DB.First(&series, seriesIDStr).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Series not found", nil)
		return
	}

	if series.VerificationCode == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Verification is not active for this series", nil)
		return
	}

	if series.VerificationCode != input.Code {
		utils.APIResponse(c, http.StatusBadRequest, "Invalid verification code", nil)
		return
	}

	verification := models.UserSeriesVerification{
		UserID:     memberID.(uint),
		SeriesID:   series.ID,
		VerifiedAt: time.Now(),
	}

	if err := database.DB.Create(&verification).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			database.DB.Preload("User").Preload("Series").First(&verification, "user_id = ? AND series_id = ?", verification.UserID, verification.SeriesID)
			utils.APIResponse(c, http.StatusOK, "You are already verified for this series", verification)
			return
		}
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to verify", err.Error())
		return
	}

	database.DB.Preload("User").Preload("Series").First(&verification, "user_id = ? AND series_id = ?", verification.UserID, verification.SeriesID)

	utils.APIResponse(c, http.StatusCreated, "Verification successful", verification)
}

