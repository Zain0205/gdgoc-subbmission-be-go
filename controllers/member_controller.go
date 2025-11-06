package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

func GetMyAchievements(c *gin.Context) {
	memberID, _ := c.Get("userID")

	var myAchievements []models.UserAchievement

	err := database.DB.Preload("Achievement.Type").Where("user_id = ?", memberID).Find(&myAchievements).Error
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to fetch achievements", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusOK, "My achievements fetched", myAchievements)
}
