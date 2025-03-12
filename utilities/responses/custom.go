package responses

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
)

const (
	StatusOK                  = 200 // RFC 9110, 15.3.1
	StatusCreated             = 201 // RFC 9110, 15.3.2
	StatusMovedPermanently    = 301 // RFC 9110, 15.4.2
	StatusFound               = 302 // RFC 9110, 15.4.3
	StatusBadRequest          = 400 // RFC 9110, 15.5.1
	StatusUnauthorized        = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired     = 402 // RFC 9110, 15.5.3
	StatusForbidden           = 403 // RFC 9110, 15.5.4
	StatusNotFound            = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed    = 405 // RFC 9110, 15.5.6
	StatusRequestTimeout      = 408 // RFC 9110, 15.5.9
	StatusConflict            = 409 // RFC 9110, 15.5.10
	StatusInternalServerError = 500 // RFC 9110, 15.6.1
	StatusBadGateway          = 502 // RFC 9110, 15.6.3
)

type ResponseOKWithDataFormat struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseOKFormat struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseErrorFormat struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseErrorWithMessageFormat struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseOKWithData(c fiber.Ctx, data interface{}, message string) error {
	response := ResponseOKWithDataFormat{
		Success: true,
		Code:    StatusOK,
		Data:    data,
		Message: message,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func ResponseCreated(c fiber.Ctx, data interface{}, message string) error {
	response := ResponseOKWithDataFormat{
		Success: true,
		Code:    StatusCreated,
		Data:    data,
		Message: message,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func ResponseOK(c fiber.Ctx, message string) error {
	response := ResponseOKFormat{
		Success: true,
		Code:    StatusOK,
		Message: message,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func RecordNotFoundError(c fiber.Ctx, err string) error {
	response := ResponseErrorWithMessageFormat{
		Success: false,
		Code:    StatusNotFound,
		Message: err,
	}

	return c.Status(http.StatusNotFound).JSON(response)
}

func BadRequestError(c fiber.Ctx, err string) error {
	response := ResponseErrorFormat{
		Success: false,
		Code:    StatusBadRequest,
		Message: err,
	}

	return c.Status(http.StatusBadRequest).JSON(response)
}

func BadRequestErrorWithMessage(c fiber.Ctx, err string) error {
	response := ResponseErrorWithMessageFormat{
		Success: false,
		Code:    StatusBadRequest,
		Message: err,
	}

	return c.Status(http.StatusBadRequest).JSON(response)
}

func DataConflictError(c fiber.Ctx, err string) error {
	response := ResponseErrorWithMessageFormat{
		Success: false,
		Code:    StatusConflict,
		Message: err,
	}

	return c.Status(http.StatusConflict).JSON(response)
}

func UnauthorizedError(c fiber.Ctx, err string) error {
	response := ResponseErrorWithMessageFormat{
		Success: false,
		Code:    StatusUnauthorized,
		Message: err,
	}

	return c.Status(http.StatusUnauthorized).JSON(response)
}

func InternalServerError(c fiber.Ctx, err error) error {
	response := ResponseErrorWithMessageFormat{
		Success: false,
		Code:    StatusInternalServerError,
		Message: err.Error(),
	}

	return c.Status(http.StatusInternalServerError).JSON(response)
}
