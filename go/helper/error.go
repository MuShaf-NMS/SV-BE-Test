package helper

import "fmt"

type CustomError struct {
	Code    int
	Message string
	Errors  interface{}
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("%d %s; errors: %s", ce.Code, ce.Message, ce.Errors)
}

// Helper to create custom error
func NewError(statusCode int, msg string, errors interface{}) error {
	return &CustomError{
		Code:    statusCode,
		Message: msg,
		Errors:  errors,
	}
}
