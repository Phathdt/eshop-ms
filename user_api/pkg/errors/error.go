package errors

import (
	"fmt"

	"user_api/pkg/sdkcm"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var (
	ErrClearCache = sdkcm.CustomError("ErrClearCache", "error when clear cache")
)

func AppRecover(logger *logrus.Logger) {
	if err := recover(); err != nil {
		if appErr, ok := err.(sdkcm.AppError); ok {
			appErr.RootCause = appErr.RootError()
			logger.Error(appErr.RootCause)
		} else {
			if fieldErrors, ok := err.(validator.ValidationErrors); ok {
				message := getMessageError(fieldErrors)

				err := sdkcm.CustomError("ValidateError", message)

				logger.Error(err.Error())
			} else if e, ok := err.(error); ok {
				logger.Error(e.Error())
			} else {
				logger.Error(err)
			}
		}
	}
}

func getMessageError(fieldErrors []validator.FieldError) string {
	fieldError := fieldErrors[0]

	//TODO: add more tag
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is a required field", fieldError.Field())
	case "max":
		return fmt.Sprintf("%s must be a maximum of %s in length", fieldError.Field(), fieldError.Param())
	case "min":
		return fmt.Sprintf("%s must be a minimum of %s in length", fieldError.Field(), fieldError.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", fieldError.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid Email", fieldError.Field())
	case "oneof":
		return fmt.Sprintf("%s must be one of enums %s", fieldError.Field(), fieldError.Param())
	default:
		return fmt.Sprintf("something wrong on %s; %s", fieldError.Field(), fieldError.Tag())
	}
}
