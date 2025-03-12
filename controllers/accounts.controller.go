package controllers

import (
	"github.com/backend-boilerplate-template/models"
	"github.com/backend-boilerplate-template/utilities/responses"
	"github.com/gofiber/fiber/v3"
	"log"
)

func CreateAccount(c fiber.Ctx) error {
	var newUser models.Profile

	if err := c.Bind().Body(&newUser); err != nil {
		log.Printf("Error parsing body: %v", err)
		return responses.InternalServerError(c, err)
	}

	//found, _ := infrastructure.FetchUserByEmail(newUser.Email)
	//if found.Email != "" {
	//	return responses.StatusText(c, http.StatusConflict, fmt.Sprintf("user email: %s already exist", found.Email), nil)
	//}
	//
	////hash pass
	//hashedPassword, err := utilities.HashPassword(newUser.Password)
	//
	//if err != nil {
	//	log.Printf("Error hashing password: %v", err)
	//	return responses.InternalServerError(c, err.Error())
	//}
	//newUser.Password = hashedPassword
	//
	///* Create new user */
	//user, saveError := infrastructure.SaveUser(newUser)
	//if saveError.Error != nil {
	//	return responses.StatusText(c, saveError.ErrorCode, saveError.ErrorMessage, nil)
	//}

	//output := models.ProfileFromDomain(&user)

	return responses.ResponseCreated(c, newUser, "new user created")
}

//func GetAllAccounts(c *fiber.Ctx) error {
//	_, validatorErr := ValidateAuthCookie(c)
//	if validatorErr != nil {
//		return responses.UnauthorizedError(c, validatorErr.Error())
//	}
//
//	page, err := strconv.Atoi(c.Query("page", "1"))
//	if err != nil || page < 1 {
//		page = 1
//	}
//
//	// Fetch Users
//	users, fetchError := infrastructure.FetchUsers(page)
//	if fetchError.Error != nil {
//		return responses.StatusText(c, fetchError.ErrorCode, fetchError.ErrorMessage, nil)
//	}
//
//	return responses.ResponseOKWithData(c, users, "account list")
//}
//
//func UpdateAccount(c *fiber.Ctx) error {
//	_, validatorErr := ValidateAuthCookie(c)
//	if validatorErr != nil {
//		return responses.UnauthorizedError(c, validatorErr.Error())
//	}
//
//	file, fileErr := c.FormFile("avatar")
//	user := models.Profile{
//		FirstName:   c.FormValue("first_name"),
//		LastName:    c.FormValue("last_name"),
//		Email:       c.FormValue("email"),
//		Phone:       c.FormValue("phone"),
//		Address:     c.FormValue("address"),
//		Country:     c.FormValue("country"),
//		AccountType: c.FormValue("account_type"),
//		Password:    c.FormValue("password"),
//		Status:      c.FormValue("status"),
//	}
//
//	id := c.Params("id")
//	if id == "" {
//		return responses.BadRequestError(c, "missing request ID")
//	}
//
//	// Check if the account exists
//	fetchResult, fetchError := infrastructure.FetchUserById(id)
//	if fetchError.Error != nil {
//		return responses.StatusText(c, fetchError.ErrorCode, fetchError.ErrorMessage, nil)
//	}
//
//	var avatarPath string
//	if file != nil {
//		if fileErr != nil {
//			return responses.InternalServerError(c, fileErr.Error())
//		}
//
//		// upload files to minio
//		uploadInfo, uploadInfoErr := utilities.MinioUpload(file, "avatars")
//
//		if uploadInfoErr != nil {
//			return responses.InternalServerError(c, uploadInfoErr.Error())
//		}
//
//		// re-assign the path
//		avatarPath = uploadInfo
//
//		// Update the avatar
//		user.Avatar = avatarPath
//	}
//
//	if user.Password != "" {
//		// Hash the new password
//		newHashedPassword, err := utilities.HashPassword(user.Password)
//
//		// Update the user password
//		user.Password = newHashedPassword
//		if err != nil {
//			log.Printf("Error hashing password: %v", err)
//			return responses.InternalServerError(c, utilities.ErrParseJSON)
//		}
//	}
//
//	// invoke the update user
//	updateResult, updateErr := infrastructure.UpdateUser(user, id)
//	if updateErr.Error != nil {
//		// delete the current uploaded file
//		if avatarPath != "" {
//			if deleteErr := utilities.MinioRemoveMedia(avatarPath, "avatars"); deleteErr != nil {
//				return responses.InternalServerError(c, deleteErr.Error())
//			}
//		}
//
//		return responses.StatusText(c, updateErr.ErrorCode, updateErr.ErrorMessage, nil)
//	}
//
//	// delete old avatar from the minio
//	if file != nil && fetchResult.Avatar != "" {
//		if deleteErr := utilities.MinioRemoveMedia(fetchResult.Avatar, "avatars"); deleteErr != nil {
//			return responses.InternalServerError(c, deleteErr.Error())
//		}
//	}
//
//	profile := models.ProfileFromDomain(&updateResult)
//	// Return the response
//	return responses.ResponseOKWithData(c, profile, "user profile updated successfully")
//}
//
//func GetAccount(c *fiber.Ctx) error {
//	_, validatorErr := ValidateAuthCookie(c)
//	if validatorErr != nil {
//		return responses.UnauthorizedError(c, validatorErr.Error())
//	}
//
//	id := c.Params("id")
//
//	if id == "" {
//		return responses.BadRequestError(c, "missing request ID")
//	}
//
//	user, fetchError := infrastructure.FetchUserById(id)
//	if fetchError.Error != nil {
//		return responses.StatusText(c, fetchError.ErrorCode, fetchError.ErrorMessage, nil)
//	}
//
//	return responses.ResponseOKWithData(c, user, "profile details")
//}
//
//func DeleteAccount(c *fiber.Ctx) error {
//	_, validatorErr := ValidateAuthCookie(c)
//	if validatorErr != nil {
//		return responses.UnauthorizedError(c, validatorErr.Error())
//	}
//
//	id := c.Params("id")
//
//	// get cookie
//	cookie := c.Cookies("auth_cookie")
//
//	if id == "" {
//		return responses.BadRequestError(c, "missing request ID in param")
//	}
//
//	// Fetch the user
//	fetchUser, fetchError := infrastructure.FetchUserById(id)
//	if fetchError.Error != nil {
//		return responses.StatusText(c, fetchError.ErrorCode, fetchError.ErrorMessage, nil)
//	}
//
//	// Remove user from db
//	if removeError := infrastructure.RemoveUser(id); removeError.Error != nil {
//		return responses.StatusText(c, removeError.ErrorCode, removeError.ErrorMessage, nil)
//	}
//
//	// Delete the user avatar
//	if fetchUser.Avatar != "" {
//		if deleteErr := utilities.MinioRemoveMedia(fetchUser.Avatar, "avatars"); deleteErr != nil {
//			return responses.InternalServerError(c, deleteErr.Error())
//		}
//	}
//
//	// clear the redis cache if the user was logged in
//	if deleteErr := utilities.RedisDeleteKey(cookie); deleteErr != nil {
//		return responses.BadRequestError(c, deleteErr.Error())
//	}
//
//	// reset the cookie
//	utilities.ResetCookie(c)
//
//	return responses.ResponseOK(c, "account deleted successfully")
//}
