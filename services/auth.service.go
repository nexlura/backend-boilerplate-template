package services

import (
	"errors"
	"github.com/backend-boilerplate-template/infrastructure"
	"github.com/backend-boilerplate-template/models"
	"github.com/backend-boilerplate-template/requests"
	"github.com/backend-boilerplate-template/utilities"
	"github.com/backend-boilerplate-template/utilities/responses"
	"github.com/gofiber/fiber/v3"
)

func LoginService(c fiber.Ctx, loginPayload requests.Login) (models.ProfileDTO, responses.ResponseError) {
	//Find the user by the email
	user, findUserError := infrastructure.FindUserByParam(loginPayload.Email)
	if findUserError.Error != nil {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        findUserError.Error,
			ErrorCode:    findUserError.ErrorCode,
			ErrorMessage: findUserError.ErrorMessage,
		}
	}

	// first check if the user is not logged in before generating a new token
	isAlreadyLoggedIn := infrastructure.CheckUserLoginState(c, user)
	if isAlreadyLoggedIn {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        errors.New("authenticated"),
			ErrorCode:    responses.StatusOK,
			ErrorMessage: "user is already authenticated",
		}
	}

	//Check passwords
	if utilities.CheckPasswordHash(user.Password, loginPayload.Password) != nil {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        errors.New("password error"),
			ErrorCode:    responses.StatusBadRequest,
			ErrorMessage: "incorrect password",
		}
	}

	//Authenticate user and generate cookie
	if authenticateError := infrastructure.GenerateCookieAndAuthenticate(c, user); authenticateError != nil {
		return models.ProfileDTO{}, responses.ResponseError{
			Error:        authenticateError,
			ErrorCode:    responses.StatusInternalServerError,
			ErrorMessage: authenticateError.Error(),
		}
	}

	return models.ProfileFromDomain(user), responses.ResponseError{}
}
