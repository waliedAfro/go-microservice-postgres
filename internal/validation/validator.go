package validation

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

// ValidationError represents a single validation error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// GetValidator returns a singleton validator instance.
func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()

		// Use JSON tag names instead of struct field names.
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.Split(fld.Tag.Get("json"), ",")[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// Register custom validations here if needed.
		// validate.RegisterValidation("booktitle", validateBookTitle)
	})

	return validate
}

// Validate validates any struct and returns formatted errors.
func Validate(v interface{}) []ValidationError {

	err := GetValidator().Struct(v)
	if err == nil {
		return nil
	}

	var errors []ValidationError

	for _, e := range err.(validator.ValidationErrors) {

		errors = append(errors, ValidationError{
			Field:   e.Field(),
			Message: buildMessage(e),
		})
	}

	return errors
}

// buildMessage converts validator tags into user-friendly messages.
func buildMessage(err validator.FieldError) string {

	switch err.Tag() {

	case "required":
		return fmt.Sprintf("%s is required", err.Field())

	case "min":
		return fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())

	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", err.Field(), err.Param())

	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())

	case "numeric":
		return fmt.Sprintf("%s must be numeric", err.Field())

	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
