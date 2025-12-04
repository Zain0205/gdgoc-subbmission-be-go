package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

type SeriesResponse struct {
	models.Series
	IsVerified    bool   `json:"is_verified"`
	IsSubmitted   bool   `json:"is_submitted"`
	SubmissionURL string `json:"submission_url,omitempty"`
}

type TrackResponse struct {
	models.Track
	Series []SeriesResponse `json:"series"`
}

func CreateTrack(c *gin.Context) {
	var input dto.CreateTrackInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	adminID, _ := c.Get("userID")

	track := models.Track{
		TrackName:   input.TrackName,
		Description: input.Description,
		CreatedByID: adminID.(uint),
		TrackType:   input.TrackType,
	}

	if track.TrackType == "" {
		track.TrackType = "STUDY_JAM"
	}

	if err := database.DB.Create(&track).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create track", err.Error())
		return
	}

	database.DB.Preload("CreatedBy").First(&track, track.ID)
	utils.APIResponse(c, http.StatusCreated, "Track created successfully", track)
}

func GetAllTracks(c *gin.Context) {
	var tracks []models.Track
	if err := database.DB.Preload("Series").Find(&tracks).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to fetch tracks", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if exists {
		var response []TrackResponse
		for _, track := range tracks {
			response = append(response, enrichTrackWithStatus(track, userID.(uint)))
		}
		utils.APIResponse(c, http.StatusOK, "Tracks fetched successfully", response)
		return
	}

	utils.APIResponse(c, http.StatusOK, "Tracks fetched successfully", tracks)
}

func GetTrackWithSeries(c *gin.Context) {
	trackID := c.Param("id")

	var track models.Track
	if err := database.DB.Preload("Series").Preload("CreatedBy").First(&track, trackID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Track not found", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if exists {
		enrichedTrack := enrichTrackWithStatus(track, userID.(uint))
		utils.APIResponse(c, http.StatusOK, "Track fetched successfully", enrichedTrack)
		return
	}

	utils.APIResponse(c, http.StatusOK, "Track fetched successfully", track)
}

func enrichTrackWithStatus(track models.Track, userID uint) TrackResponse {
	var verifications []models.UserSeriesVerification
	verifiedMap := make(map[uint]bool)
	database.DB.Where("user_id = ?", userID).Find(&verifications)
	for _, v := range verifications {
		verifiedMap[v.SeriesID] = true
	}

	var submissions []models.Submission
	submissionMap := make(map[uint]string)
	database.DB.Where("user_id = ?", userID).Find(&submissions)
	for _, s := range submissions {
		submissionMap[s.SeriesID] = s.FileURL
	}

	var seriesResponses []SeriesResponse
	for _, s := range track.Series {
		isVerified := verifiedMap[s.ID]
		submissionURL, isSubmitted := submissionMap[s.ID]

		seriesResponses = append(seriesResponses, SeriesResponse{
			Series:        s,
			IsVerified:    isVerified,
			IsSubmitted:   isSubmitted,
			SubmissionURL: submissionURL,
		})
	}

	return TrackResponse{
		Track:  track,
		Series: seriesResponses,
	}
}

