package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (v ValidationErrors) Error() string {
	var messages []string
	for _, err := range v.Errors {
		messages = append(messages, err.Message)
	}
	return strings.Join(messages, ", ")
}

func Validate(data interface{}) error {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	var validationErrors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, ValidationError{
			Field:   strings.ToLower(err.Field()),
			Message: getErrorMessage(err),
		})
	}

	return ValidationErrors{Errors: validationErrors}
}

func getErrorMessage(err validator.FieldError) string {
	field := strings.ToLower(err.Field())

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		if err.Type().String() == "string" {
			return fmt.Sprintf("%s must be at least %s characters", field, err.Param())
		}
		return fmt.Sprintf("%s must be at least %s", field, err.Param())
	case "max":
		if err.Type().String() == "string" {
			return fmt.Sprintf("%s must not exceed %s characters", field, err.Param())
		}
		return fmt.Sprintf("%s must not exceed %s", field, err.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, err.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, err.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, err.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "numeric":
		return fmt.Sprintf("%s must be a number", field)
	case "alpha":
		return fmt.Sprintf("%s must contain only letters", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
