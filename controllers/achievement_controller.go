package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateAchievementType(c *gin.Context) {
	var input dto.CreateAchievementTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	achievType := models.AchievementType{Name: input.Name}
	if err := database.DB.Create(&achievType).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create achievement type", err.Error())
		return
	}
	utils.APIResponse(c, http.StatusCreated, "Achievement type created", achievType)
}

func GetAchievementTypes(c *gin.Context) {
	var types []models.AchievementType
	database.DB.Find(&types)
	utils.APIResponse(c, http.StatusOK, "Achievement types fetched", types)
}

func CreateAchievement(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	typeIDStr := c.PostForm("achievement_type_id")
	iconURL := c.PostForm("icon_url")

	if name == "" || description == "" {
		utils.APIResponse(c, http.StatusBadRequest, "Name and Description are required", nil)
		return
	}

	file, err := c.FormFile("icon_file")
	if err == nil {
		uploadPath := "uploads/badges"
		if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
			os.MkdirAll(uploadPath, 0o755)
		}

		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		dst := filepath.Join(uploadPath, filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			utils.APIResponse(c, http.StatusInternalServerError, "Failed to save file", err.Error())
			return
		}

		iconURL = fmt.Sprintf("/uploads/badges/%s", filename)
	}

	typeID, _ := strconv.Atoi(typeIDStr)
	if typeID == 0 {
		typeID = 1
	}

	var count int64
	database.DB.Model(&models.AchievementType{}).Where("id = ?", typeID).Count(&count)
	if count == 0 {
		defaultType := models.AchievementType{ID: uint(typeID), Name: "General"}
		if err := database.DB.Create(&defaultType).Error; err != nil {
			utils.APIResponse(c, http.StatusInternalServerError, "Failed to create default category", err.Error())
			return
		}
	}

	achiev := models.Achievement{
		Name:              name,
		Description:       description,
		IconURL:           iconURL,
		AchievementTypeID: uint(typeID),
	}

	if err := database.DB.Create(&achiev).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create achievement", err.Error())
		return
	}

	database.DB.Preload("Type").First(&achiev, achiev.ID)
	utils.APIResponse(c, http.StatusCreated, "Achievement created", achiev)
}

func GetAchievements(c *gin.Context) {
	var achievs []models.Achievement
	database.DB.Preload("Type").Find(&achievs)
	utils.APIResponse(c, http.StatusOK, "Achievements fetched", achievs)
}

func UpdateAchievement(c *gin.Context) {
	id := c.Param("id")
	var achiev models.Achievement
	if err := database.DB.First(&achiev, id).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Achievement not found", nil)
		return
	}

	var input dto.UpdateAchievementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := database.DB.Model(&achiev).Updates(input).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to update achievement", err.Error())
		return
	}

	database.DB.Preload("Type").First(&achiev, achiev.ID)
	utils.APIResponse(c, http.StatusOK, "Achievement updated", achiev)
}

func AwardAchievementToUser(c *gin.Context) {
	var input dto.AwardAchievementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var user models.User
	if err := database.DB.First(&user, input.UserID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	var achiev models.Achievement
	if err := database.DB.First(&achiev, input.AchievementID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "Achievement not found", nil)
		return
	}

	userAchiev := models.UserAchievement{
		UserID:        input.UserID,
		AchievementID: input.AchievementID,
		EarnedAt:      time.Now(),
	}

	if err := database.DB.FirstOrCreate(&userAchiev).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to award achievement", err.Error())
		return
	}

	database.DB.Preload("User").Preload("Achievement.Type").First(&userAchiev, "user_id = ? AND achievement_id = ?", userAchiev.UserID, userAchiev.AchievementID)
	utils.APIResponse(c, http.StatusCreated, "Achievement awarded", userAchiev)
}

func RevokeAchievementFromUser(c *gin.Context) {
	var input dto.AwardAchievementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	result := database.DB.Delete(&models.UserAchievement{}, "user_id = ? AND achievement_id = ?", input.UserID, input.AchievementID)
	if result.Error != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error revoking achievement", result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		utils.APIResponse(c, http.StatusNotFound, "User does not have this achievement", nil)
		return
	}

	utils.APIResponse(c, http.StatusOK, "Achievement revoked", nil)
}

func DeleteAchievement(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Achievement{}, id).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to delete achievement", err.Error())
		return
	}
	utils.APIResponse(c, http.StatusOK, "Achievement deleted", nil)
}
