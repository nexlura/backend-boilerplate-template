package infrastructure

import (
	"github.com/backend-boilerplate-template/models"
	"github.com/backend-boilerplate-template/utilities"
	"github.com/gofiber/fiber/v3"
	"log"
)

func GenerateCookieAndAuthenticate(c fiber.Ctx, user models.Profile) error {
	// create auth cookie
	generatedCookie, generatedCookieErr := utilities.GenerateCookie(user)

	// Check and throw error weather any occurred when generating
	if generatedCookieErr != nil {
		log.Printf("Error generating cookie: %v", generatedCookieErr)
		return generatedCookieErr
	}

	domain := utilities.GetCookieDomain(c.Get("Origin"))

	// write cookie to the client
	c.Cookie(&fiber.Cookie{
		Name:     generatedCookie.CookieName,
		Value:    generatedCookie.CookieValue,
		Expires:  generatedCookie.ExpirationTime,
		Path:     "/",
		Domain:   domain,
		Secure:   false,
		HTTPOnly: true,
	})

	// update the user data with the new auth token
	_, updateError := AlterUser(models.ProfileFrom{ID: user.ID, AuthToken: generatedCookie.CookieValue})
	if updateError.Error != nil {
		return updateError.Error
	}

	return nil
}

func CheckUserLoginState(c fiber.Ctx, user models.Profile) bool {

	// check if the user has an auth token and set it to the cookie
	if user.AuthToken != "" && user.AuthToken != "null" {
		// get the token data from the cache
		cacheCookie, cacheCookieErr := utilities.GetCacheCookie(user.AuthToken)
		if cacheCookieErr != nil {
			return false
		}

		domain := utilities.GetCookieDomain(c.Get("Origin"))

		// write cookie to the client
		c.Cookie(&fiber.Cookie{
			Name:     cacheCookie["CookieName"].(string),
			Value:    cacheCookie["CookieValue"].(string),
			Path:     "/",
			Domain:   domain,
			Secure:   false,
			HTTPOnly: true,
		})

		return true
	}

	return false
}
