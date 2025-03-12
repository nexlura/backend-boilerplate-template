package controllers

//func Login(c fiber.Ctx) error {
//	var loginRequest requests.Login
//
//	if err := c.BodyParser(&loginRequest); err != nil {
//		log.Printf("Error parsing body: %v", err)
//		return responses.InternalServerError(c, "Error parsing body")
//	}
//
//	//Fetch the user
//	user, fetchError := infrastructure.FetchUserByEmail(loginRequest.Email)
//	if fetchError.Error != nil {
//		return responses.StatusText(c, fetchError.ErrorCode, fetchError.ErrorMessage, nil)
//	}
//
//	// response output
//	output := models.ProfileFromDomain(&user)
//
//	// first check if the user is not logged in before generating a new token
//	isAlreadyLoggedIn := infrastructure.CheckUserLoginState(c, user)
//	if isAlreadyLoggedIn {
//		return responses.ResponseOKWithData(c, output, "you are already authenticated")
//	}
//
//	//Check passwords
//	if utilities.CheckPasswordHash(user.Password, loginRequest.Password) != nil {
//		return responses.BadRequestErrorWithMessage(c, "incorrect password")
//	}
//
//	//Authenticate user and generate cookie
//	if authenticateError := infrastructure.GenerateCookieAndAuthenticate(c, user); authenticateError != nil {
//		return responses.InternalServerError(c, authenticateError.Error())
//	}
//
//	return responses.ResponseOKWithData(
//		c, output, "login successfully")
//}
//
//func Logout(c fiber.Ctx) error {
//
//	// get cookie
//	cookie := c.Cookies("auth_cookie")
//
//	// throw error if cookieErr exists
//	if cookie == "" {
//		return responses.UnauthorizedError(c, "auth cookie not provided")
//	}
//
//	cacheCookie, validatorErr := ValidateAuthCookie(c)
//	if validatorErr != nil {
//		return responses.UnauthorizedError(c, validatorErr.Error())
//	}
//
//	// get the cache values
//	profileId := cacheCookie["Profile"].(map[string]interface{})["id"].(string)
//
//	// clear the cookie from redis cache
//	if deleteErr := utilities.RedisDeleteKey(cookie); deleteErr != nil {
//		return responses.BadRequestError(c, deleteErr.Error())
//	}
//
//	// write cookie to the client
//	utilities.ResetCookie(c)
//
//	// update the user data with the new auth token
//	_, updateError := infrastructure.UpdateUser(models.Profile{AuthToken: "null"}, profileId)
//	if updateError.Error != nil {
//		return updateError.Error
//	}
//
//	return responses.ResponseOK(c, "logout successfully")
//}
//
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
