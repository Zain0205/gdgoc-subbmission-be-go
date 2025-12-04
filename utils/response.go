package utils

import (
	"github.com/Zain0205/gdgoc-subbmission-be-go/validation"
	"github.com/gin-gonic/gin"
)

func APIResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

func ValidationErrorResponse(c *gin.Context, err error) {
	if validationErr, ok := err.(validation.ValidationErrors); ok {
		if len(validationErr.Errors) == 1 {
			c.JSON(400, gin.H{
				"message": "Validation failed",
				"data": gin.H{
					"field":   validationErr.Errors[0].Field,
					"message": validationErr.Errors[0].Message,
				},
			})
			return
		}

		c.JSON(400, gin.H{
			"message": "Validation failed",
			"data": gin.H{
				"errors": validationErr.Errors,
			},
		})
		return
	}

	c.JSON(400, gin.H{
		"message": "Invalid request",
		"data":    nil,
	})
}
