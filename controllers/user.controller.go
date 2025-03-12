package controllers

import (
	"github.com/backend-boilerplate-template/models"
	"github.com/backend-boilerplate-template/services"
	"github.com/backend-boilerplate-template/utilities/responses"
	"github.com/gofiber/fiber/v3"
	"log"
)

func CreateUser(c fiber.Ctx) error {
	// Get the user payload
	var newUserPayload models.Profile

	// Bind the payload and return error if any
	if err := c.Bind().Body(&newUserPayload); err != nil {
		log.Printf("Error parsing body: %v", err)
		return responses.InternalServerError(c, err)
	}

	// Pass payload to service for business logic handling
	result, resultError := services.CreateUserService(newUserPayload)

	// Throw return error if any
	if resultError.Error != nil {
		return responses.DynamicStatus(c, resultError.ErrorCode, resultError.ErrorMessage, nil)
	}

	// Return final response
	return responses.ResponseCreated(c, result, "new user created")
}

func GetUsers(c fiber.Ctx) error {
	// Fetch the users
	result, resultError := services.ListUsersService(c)

	// Throw return error if any
	if resultError != nil {
		return responses.InternalServerError(c, resultError)
	}

	// Return final response
	return responses.ResponseOKWithData(c, result, "users list")
}

func GetUser(c fiber.Ctx) error {
	// Fetch the user
	result, resultError := services.GetUserService(c)

	if resultError.Error != nil {
		return responses.DynamicStatus(c, resultError.ErrorCode, resultError.ErrorMessage, nil)
	}

	return responses.ResponseOKWithData(c, result, "profile details")
}

func UpdateUser(c fiber.Ctx) error {
	// Get the user payload
	var userPayload models.ProfileFrom

	// Bind the payload and return error if any
	if err := c.Bind().Form(&userPayload); err != nil {
		log.Printf("Error parsing body: %v", err)
		return responses.InternalServerError(c, err)
	}

	// Pass payload to service for business logic handling
	result, resultError := services.UpdateUserService(c, userPayload)

	// Throw return error if any
	if resultError.Error != nil {
		return responses.DynamicStatus(c, resultError.ErrorCode, resultError.ErrorMessage, nil)
	}

	// Return the response
	return responses.ResponseOKWithData(c, result, "updated profile")
}

func DeleteUser(c fiber.Ctx) error {
	// Pass payload to service for business logic handling
	result, resultError := services.DeleteUserService(c)

	// Throw return error if any
	if resultError.Error != nil {
		return responses.DynamicStatus(c, resultError.ErrorCode, resultError.ErrorMessage, nil)
	}

	return responses.ResponseOKWithData(c, result, "user deleted")
}
