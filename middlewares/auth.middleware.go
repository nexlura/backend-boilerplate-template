package middlewares

import (
	"github.com/backend-boilerplate-template/utilities"
	"github.com/backend-boilerplate-template/utilities/responses"
	"github.com/gofiber/fiber/v3"
	"regexp"
)

func AuthMiddleware(c fiber.Ctx) error {
	//fmt.Println("auth middleware running...")
	// Get app path
	path := c.OriginalURL()

	// setup regex and check if string match
	re := regexp.MustCompile(`generate|properties`)
	method := string(c.Request().Header.Method())

	// require auth cookie if path doesn't match any in the regex
	if method != "GET" {
		if !re.MatchString(path) {
			// get cookie
			cookie := c.Cookies("auth_cookie")

			// throw error if cookieErr exists
			if cookie == "" {
				return responses.UnauthorizedError(c, "please include a valid cookie in your header when making a request")
			}

			// throw error if the cookie doesn't exist in the cache
			_, cacheCookieErr := utilities.GetCacheCookie(cookie)
			if cacheCookieErr != nil {
				return responses.InternalServerError(c, cacheCookieErr.Error())
			}

			// validate the cookie
			validatedCookie, validatedCookieErr := utilities.ValidateCookie(cookie)

			// throw error if the cookie is invalid
			if validatedCookieErr != nil {
				return responses.UnauthorizedError(c, "Invalid cookie! You are not authenticated to perform any action.")
			}

			// throw error for expires session
			if validatedCookie.HasExpired {
				return responses.UnauthorizedError(c, "Your session expired! Please login to get access.")
			}
		}
	}

	return c.Next()
}
