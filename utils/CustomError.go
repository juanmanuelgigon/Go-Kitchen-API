package utils

import "fmt"

type CustomError struct {
	Code    string
	Message string
}

func NewCustomError(code, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
