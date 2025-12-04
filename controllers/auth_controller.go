package controllers

import (
	"net/http"

	"github.com/Zain0205/gdgoc-subbmission-be-go/database"
	"github.com/Zain0205/gdgoc-subbmission-be-go/dto"
	"github.com/Zain0205/gdgoc-subbmission-be-go/models"
	"github.com/Zain0205/gdgoc-subbmission-be-go/utils"
	"github.com/Zain0205/gdgoc-subbmission-be-go/validation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.Validate(input); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	var userCount int64
	database.DB.Model(&models.User{}).Count(&userCount)

	userRole := "member"
	if userCount == 0 {
		userRole = "admin"
	}

	user := models.User{
		Name:  input.Name,
		Email: input.Email,
		Role:  userRole,
	}

	if err := user.SetPassword(input.Password); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to hash password", err.Error())
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	utils.APIResponse(c, http.StatusCreated, "User registered successfully", user)
}

func Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.Validate(input); err != nil {
		utils.ValidationErrorResponse(c, err)
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.APIResponse(c, http.StatusUnauthorized, "Invalid email or password", nil)
			return
		}
		utils.APIResponse(c, http.StatusInternalServerError, "Database error", err.Error())
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		utils.APIResponse(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	utils.APIResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}
