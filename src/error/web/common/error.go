package common

import (
	"fmt"
)

type AppError struct {
	Message   string
	ErrorInfo interface{}
}

func (err AppError) String() string {
	return fmt.Sprintf("AppError{Message: %v, ErrorInfo: %v}", err.Message, err.ErrorInfo)
}

func (err AppError) Error() string {
	return err.String()
}
