package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/gin-gonic/gin"
)

// SetUserRole (Admin)
func SetUserRole(c *gin.Context) {
	userID := c.Param("id")
	var input dto.UpdateUserRoleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.APIResponse(c, http.StatusNotFound, "User not found", nil)
		return
	}

	// Update role
	user.Role = input.Role
	if err := database.DB.Save(&user).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to update user role", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusOK, "User role updated successfully", user)
}

// GetAllUsers (Admin)
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Select("id, name, email, role, created_at").Where("role = ?", "member").Find(&users).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}
	utils.APIResponse(c, http.StatusOK, "Users fetched successfully", users)
}

