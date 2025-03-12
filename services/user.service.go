package services

import (
	"errors"
	"fmt"
	"github.com/backend-boilerplate-template/infrastructure"
	"github.com/backend-boilerplate-template/models"
	"github.com/backend-boilerplate-template/utilities"
	"github.com/backend-boilerplate-template/utilities/responses"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
)

func CreateUserService(newUserPayload models.Profile) (models.ProfileDTO, responses.ResponseError) {
	// Fetch the user with matching email
	fetcherOne, _ := infrastructure.FetchUserByParam(newUserPayload.Email)
	fetcherTwo, _ := infrastructure.FetchUserByParam(newUserPayload.Phone)

	// Check weather the found email or phone is not empty and throw error
	if !utilities.IsStringEmpty(fetcherOne.ID) || !utilities.IsStringEmpty(fetcherTwo.ID) {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        errors.New("record match"),
			ErrorCode:    responses.StatusConflict,
			ErrorMessage: fmt.Sprintf("user with data: %s%s already exist", fetcherOne.Email, fetcherTwo.Phone),
		}
	}

	//hash password
	hashedPassword, err := utilities.HashPassword(newUserPayload.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        errors.New("hashing error"),
			ErrorCode:    responses.StatusInternalServerError,
			ErrorMessage: "error hashing password",
		}
	}

	// set the user password to the hashed
	newUserPayload.Password = hashedPassword

	// assign a default status if status is empty
	if newUserPayload.Status == "" {
		newUserPayload.Status = "active"
	}

	// Pass the newUserPayload for db operation
	infraResult, infraError := infrastructure.SaveUser(newUserPayload)

	// Return db error if any
	if infraError.Error != nil {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        infraError.Error,
			ErrorCode:    infraError.ErrorCode,
			ErrorMessage: infraError.ErrorMessage,
		}
	}

	// Return response
	return models.ProfileFromDomain(infraResult), responses.ResponseError{}
}

func ListUsersService(c fiber.Ctx) ([]models.ProfileDTO, error) {
	//_, validatorErr := ValidateAuthCookie(c)
	//if validatorErr != nil {
	//	return responses.UnauthorizedError(c, validatorErr.Error())
	//}

	// Convert the page number from string to int
	page, err := strconv.Atoi(c.Query("page", "1"))

	// Set page number to 1 if there is an error or the page number is lesser than 1
	if err != nil || page < 1 {
		page = 1
	}

	// Pass the page to fetch the users
	infraResult, infraError := infrastructure.FetchUsers(page)
	if infraError.Error != nil {
		return nil, infraError.Error
	}

	// Return the response
	return models.ProfileFromDomainList(infraResult), nil
}

func GetUserService(c fiber.Ctx) (models.ProfileDTO, responses.ResponseError) {
	//_, validatorErr := ValidateAuthCookie(c)
	//if validatorErr != nil {
	//	return responses.UnauthorizedError(c, validatorErr.Error())
	//}

	// Get the param from the url
	param := c.Params("param")

	// Throw error if the param is empty
	if utilities.IsStringEmpty(param) {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        errors.New("missing param"),
			ErrorCode:    responses.StatusBadRequest,
			ErrorMessage: "missing param in the request",
		}
	}

	// Pass the para for db operation
	infraResult, infraError := infrastructure.FetchUserByParam(param)
	if infraError.Error != nil {
		return models.ProfileDTO{}, infraError
	}

	// Return the response
	return models.ProfileFromDomain(infraResult), responses.ResponseError{}
}

func UpdateUserService(c fiber.Ctx, userPayload models.ProfileFrom) (models.ProfileDTO, responses.ResponseError) {
	//_, validatorErr := ValidateAuthCookie(c)
	//if validatorErr != nil {
	//	return responses.UnauthorizedError(c, validatorErr.Error())
	//}

	// Get the id passed in the param
	paramId := c.Params("id")

	// Throw error if the id param is empty
	if utilities.IsStringEmpty(paramId) {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        errors.New("params missing"),
			ErrorCode:    responses.StatusBadRequest,
			ErrorMessage: "missing id in the request",
		}
	}

	// Assign the param_id to the user payload ID
	userPayload.ID = paramId

	// Check if the user exists
	_, fetchError := infrastructure.FetchUserByParam(userPayload.ID)
	if fetchError.Error != nil {
		return models.ProfileDTO{}, fetchError
	}

	if !utilities.IsStringEmpty(userPayload.Password) {
		// Hash the new password
		newHashedPassword, err := utilities.HashPassword(userPayload.Password)

		// Update the user password
		userPayload.Password = newHashedPassword
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			return models.ProfileDTO{}, responses.ResponseError{
				Error:        errors.New("hashing error"),
				ErrorCode:    responses.StatusInternalServerError,
				ErrorMessage: "error hashing password",
			}
		}
	}

	// Pass the payload to the db for operation
	infraResult, infraError := infrastructure.AlterUser(userPayload)
	if infraError.Error != nil {
		// delete the current uploaded file
		//if avatarPath != "" {
		//	if deleteErr := utilities.MinioRemoveMedia(avatarPath, "avatars"); deleteErr != nil {
		//		return responses.InternalServerError(c, deleteErr.Error())
		//	}
		//}

		return models.ProfileDTO{}, infraError
	}

	//// delete old avatar from the minio
	//if file != nil && fetchResult.Avatar != "" {
	//	if deleteErr := utilities.MinioRemoveMedia(fetchResult.Avatar, "avatars"); deleteErr != nil {
	//		return responses.InternalServerError(c, deleteErr.Error())
	//	}
	//}

	// Return response
	return models.ProfileFromDomain(infraResult), responses.ResponseError{}
}

func DeleteUserService(c fiber.Ctx) (models.ProfileDTO, responses.ResponseError) {
	//_, validatorErr := ValidateAuthCookie(c)
	//if validatorErr != nil {
	//	return responses.UnauthorizedError(c, validatorErr.Error())
	//}

	// Get the id passed in the param
	paramId := c.Params("id")

	// Throw error if the id param is empty
	if utilities.IsStringEmpty(paramId) {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        errors.New("params missing"),
			ErrorCode:    responses.StatusBadRequest,
			ErrorMessage: "missing id in the request",
		}
	}

	// Check if the user exists
	fetchUser, fetchError := infrastructure.FetchUserByParam(paramId)
	if fetchError.Error != nil {
		return models.ProfileDTO{}, fetchError
	}

	// Pass the user id for db remove operation
	if removeError := infrastructure.RemoveUser(paramId); removeError.Error != nil {
		return models.ProfileDTO{}, removeError
	}

	// Delete the user avatar
	if !utilities.IsStringEmpty(fetchUser.Avatar) {
		if avatarDeleteErr := utilities.MinioRemoveMedia(fetchUser.Avatar, "avatars"); avatarDeleteErr != nil {
			return models.ProfileDTO{}, responses.ResponseError{
				Error:        errors.New("user avatar delete error"),
				ErrorCode:    responses.StatusInternalServerError,
				ErrorMessage: avatarDeleteErr.Error(),
			}
		}
	}

	// Return response
	return models.ProfileFromDomain(fetchUser), responses.ResponseError{}
}
