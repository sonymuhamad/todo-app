package httpserver

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTTPError interface {
	ToHTTPError() (int, []HTTPResponseError)
}

type HTTPResponseError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type UnprocessableError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

func (e UnprocessableError) Error() string {
	return e.Message
}

func (e UnprocessableError) ToHTTPError() (int, []HTTPResponseError) {
	return http.StatusUnprocessableEntity, []HTTPResponseError{
		{
			Field:   e.Field,
			Message: e.Message,
		},
	}
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var report *echo.HTTPError
	ok := errors.As(err, &report)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	switch origErr := err.(type) {
	case HTTPError:
		code, _ := origErr.ToHTTPError()
		report.Code = code

	_:
		c.JSON(origErr.ToHTTPError())
	case validator.ValidationErrors:
		respErrors := make([]HTTPResponseError, 0)
		report.Code = http.StatusBadRequest

		for _, validationErr := range origErr {
			respError := HTTPResponseError{
				Field: validationErr.Field(),
			}

			switch validationErr.Tag() {
			case "required":
				respError.Message = fmt.Sprintf("%s is required",
					validationErr.Field())
			case "email":
				respError.Message = fmt.Sprintf("%s is not valid email",
					validationErr.Field())
			case "gte":
				respError.Message = fmt.Sprintf("%s value must be greater than %s",
					validationErr.Field(), validationErr.Param())
			case "lte":
				respError.Message = fmt.Sprintf("%s value must be lower than %s",
					validationErr.Field(), validationErr.Param())
			case "min":
				respError.Message = fmt.Sprintf("%s value must be have at least %s character",
					validationErr.Field(), validationErr.Param())
			default:
				respError.Message = "Invalid parameter"
			}

			respErrors = append(respErrors, respError)
		}

	_:
		c.JSON(http.StatusBadRequest, respErrors)
	default:
	_:
		c.JSON(http.StatusInternalServerError, origErr.Error())
	}

	c.Logger().Error(report)
}
