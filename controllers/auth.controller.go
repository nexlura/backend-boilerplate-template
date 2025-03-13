package controllers

import (
	"github.com/backend-boilerplate-template/requests"
	"github.com/backend-boilerplate-template/services"
	"github.com/backend-boilerplate-template/utilities/responses"
	"github.com/gofiber/fiber/v3"
	"log"
)

func Login(c fiber.Ctx) error {
	// Get the request payload
	var loginRequest requests.Login

	// Bind the request data
	if err := c.Bind().Body(&loginRequest); err != nil {
		log.Printf("Error parsing body: %v", err)
		return responses.InternalServerError(c, err)
	}

	// Pass the request to the login service
	result, resultError := services.LoginService(c, loginRequest)

	// Check weather the result error exits and throw it
	if resultError.Error != nil {
		return responses.DynamicStatus(c, resultError.ErrorCode, resultError.ErrorMessage, nil)
	}

	// return the response
	return responses.ResponseOKWithData(c, result, "login successfully")
}

func Logout(c fiber.Ctx) error {
	// Invoke the logout service
	_, resultError := services.LogoutService(c)

	// Throw error if any occurs
	if resultError.Error != nil {
		return responses.DynamicStatus(c, resultError.ErrorCode, resultError.ErrorMessage, nil)
	}

	return responses.ResponseOK(c, "logout successfully")
}

//func ForgotPassword(c fiber.Ctx) error {
//	var request struct {
//		Email string `json:"email"`
//	}
//
//	if err := c.BodyParser(&request); err != nil {
//		log.Printf("Error parsing body: %v", err)
//		return responses.InternalServerError(c, utilities.ErrParseJSON)
//	}
//
//	//Fetch the user
//	user, fetchError := infrastructure.FetchUserByEmail(request.Email)
//	if fetchError.Error != nil {
//		return responses.StatusText(c, fetchError.ErrorCode, fetchError.ErrorMessage, nil)
//	}
//
//	// generate the key
//	resetLinkToken := utilities.GenerateUUID()
//
//	// set the cache to expire in 10 minutes
//	utilities.RedisSetCache(resetLinkToken, request.Email, 30*time.Minute)
//
//	pars := utilities.GmailStruct{
//		RecipientEmail: user.Email,
//		RecipientName:  user.FirstName,
//		LinkToken:      resetLinkToken,
//	}
//
//	// send email to the user
//	utilities.GmailSendResetPasswordEmail(pars)
//
//	return responses.ResponseOK(c, "reset email sent")
//}
//
//func ResetPassword(c fiber.Ctx) error {
//	var request requests.ResetPassword
//
//	if err := c.BodyParser(&request); err != nil {
//		log.Printf("Error parsing body: %v", err)
//		return responses.InternalServerError(c, utilities.ErrParseJSON)
//	}
//
//	// get the token's value from the cache
//	cache, cacheErr := utilities.RedisGetCache(request.Token)
//
//	if cacheErr != nil {
//		return responses.RecordNotFoundError(c, "provided token not found")
//	}
//
//	// check that passwords mismatch
//	if passwordsMatchErr := utilities.ValidatePasswordsMatch(request.Password, request.PasswordConfirmation); passwordsMatchErr != nil {
//		return responses.BadRequestError(c, passwordsMatchErr.Error())
//	}
//
//	//Fetch the user
//	user, fetchError := infrastructure.FetchUserByEmail(strings.Trim(cache, "\""))
//	if fetchError.Error != nil {
//		return responses.StatusText(c, fetchError.ErrorCode, fetchError.ErrorMessage, nil)
//	}
//
//	// Hash the new password
//	newHashedPassword, err := utilities.HashPassword(request.Password)
//	if err != nil {
//		log.Printf("Error hashing password: %v", err)
//		return responses.InternalServerError(c, utilities.ErrParseJSON)
//	}
//
//	user.Password = newHashedPassword
//	//Update the user record
//	_, updateErr := infrastructure.UpdateUser(user, user.ID)
//	if updateErr.Error != nil {
//		return responses.StatusText(c, updateErr.ErrorCode, updateErr.ErrorMessage, nil)
//	}
//
//	// Delete token's value from the cache
//	redisErr := utilities.RedisDeleteKey(request.Token)
//	if redisErr != nil {
//		return responses.InternalServerError(c, redisErr.Error())
//	}
//
//	return responses.ResponseOK(c, "password reset successful")
//}
