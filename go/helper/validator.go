package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// func customValidator(fl validator.FieldLevel) bool {

// 	if fl.Field().String() == "invalid" {
// 		return false
// 	}

// 	return true
// }

var Validator = validator.New()

func RegisterCustomValidation() {
	// Validator.RegisterValidation()
	// Validator.RegisterCustomTypeFunc()
}

// Helper to validate struct
func Validate(s interface{}) error {
	return Validator.Struct(s)
}

// Helper to extrac errors from struct error
func ValidationError(err error) []string {
	errMsg := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errMsg = append(errMsg, fmt.Sprintf("%s %s", e.Field(), e.ActualTag()))
	}
	return errMsg
}
