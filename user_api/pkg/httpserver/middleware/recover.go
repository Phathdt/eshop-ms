package middleware

import (
	"fmt"
	"net/http"

	"user_api/pkg/sdkcm"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func Recover(logger *logrus.Logger) fiber.Handler {
	// Return new handler
	return func(ctx *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				if appErr, ok := err.(sdkcm.AppError); ok {
					appErr.RootCause = appErr.RootError()
					logger.Error(appErr.RootCause)

					if appErr.RootCause != nil {
						appErr.Log = appErr.RootCause.Error()
					}

					if err = ctx.Status(appErr.StatusCode).JSON(&fiber.Map{
						"errors": appErr,
					}); err != nil {
						return
					}

				} else {
					var appErr sdkcm.AppError

					if fieldErrors, ok := err.(validator.ValidationErrors); ok {
						message := getMessageError(fieldErrors)

						err := sdkcm.CustomError("ValidateError", message)

						appErr := sdkcm.ErrCustom(err, err)

						logger.Error(err.Error())

						if err := ctx.Status(appErr.StatusCode).JSON(&fiber.Map{
							"errors": appErr,
						}); err != nil {
							return
						}

					} else if e, ok := err.(error); ok {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: "internal server errors"}
						logger.Error(e.Error())

						if err := ctx.Status(appErr.StatusCode).JSON(&fiber.Map{
							"errors": appErr,
						}); err != nil {
							return
						}

					} else {
						appErr = sdkcm.AppError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("%s", err)}
						logger.Error(fmt.Sprintf("%s", err))

						if err := ctx.Status(appErr.StatusCode).JSON(&fiber.Map{
							"errors": appErr,
						}); err != nil {
							return
						}

					}
				}
			}
		}()

		// Return err if existed, else move to next handler
		return ctx.Next()
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
