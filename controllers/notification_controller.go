package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
	userID := c.GetUint("userID")
	var notifs []models.Notification
	database.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifs)
	utils.APIResponse(c, http.StatusOK, "Success", notifs)
}

func MarkAsRead(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("userID")

	var n models.Notification
	if err := database.DB.First(&n, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Notification not found", nil)
		return
	}
	n.IsRead = true
	database.DB.Save(&n)
	utils.APIResponse(c, http.StatusOK, "Marked as read", nil)
}
