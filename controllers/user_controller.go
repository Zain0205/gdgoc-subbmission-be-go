package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/Zain0205/gdgoc-subbmission-be-go/validation"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	userID := c.GetUint("userID")

	var user models.User
	database.DB.Select("id, name, email, role, avatar_url").First(&user, userID)

	utils.APIResponse(c, http.StatusOK, "Success", user)
}

func UpdateProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	var input dto.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.Validate(input); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	var user models.User
	database.DB.First(&user, userID)

	user.Name = input.Name
	database.DB.Save(&user)

	utils.APIResponse(c, http.StatusOK, "Profile updated", user)
}

func ChangePassword(c *gin.Context) {
	userID := c.GetUint("userID")

	var input dto.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.Validate(input); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	var user models.User
	database.DB.First(&user, userID)

	if err := user.CheckPassword(input.CurrentPassword); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Current password incorrect", nil)
		return
	}

	user.SetPassword(input.NewPassword)
	database.DB.Save(&user)

	utils.APIResponse(c, http.StatusOK, "Password changed", nil)
}

func UpdateAvatar(c *gin.Context) {
	userID := c.GetUint("userID")

	file, err := c.FormFile("avatar")
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Image required", nil)
		return
	}

	if file.Size > 5*1024*1024 {
		utils.APIResponse(c, http.StatusBadRequest, "Max 5MB", nil)
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		utils.APIResponse(c, http.StatusBadRequest, "Only JPG/PNG/WebP", nil)
		return
	}

	filename := utils.RandomString(20) + ext
	path := "./uploads/avatars/" + filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Save failed", nil)
		return
	}

	var user models.User
	database.DB.First(&user, userID)

	if user.AvatarURL != "" {
		oldPath := "./uploads/avatars/" + filepath.Base(user.AvatarURL)
		utils.SafeDeleteFile(oldPath)
	}

	user.AvatarURL = "/uploads/avatars/" + filename
	database.DB.Save(&user)

	utils.APIResponse(c, http.StatusOK, "Avatar updated", gin.H{
		"avatar_url": user.AvatarURL,
	})
}

func SetUserRole(c *gin.Context) {
	userID := c.Param("id")

	var input dto.UpdateUserRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.Validate(input); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	user.Role = input.Role
	database.DB.Save(&user)

	utils.APIResponse(c, http.StatusOK, "Role updated", user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Select("id, name, email, role, created_at, avatar_url").Where("role = ?", "member").Find(&users)

	utils.APIResponse(c, http.StatusOK, "Success", users)
}
