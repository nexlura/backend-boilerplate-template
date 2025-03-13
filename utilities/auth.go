package utilities

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(hashPassword, loginRequestPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(loginRequestPassword))
}

var jwtSecretKey = os.Getenv("JWT_SECRET")

func GenerateJWT(userID int, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

func DecodeJWT(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("error retrieving user ID from token")
		}
		return int(userId), nil
	}

	return 0, errors.New("invalid token")
}

func ValidateAuthCookie(c fiber.Ctx) (map[string]interface{}, error) {
	// get cookie
	cookie := c.Cookies("auth_cookie")

	// throw error if cookieErr exists
	if cookie == "" {
		return nil, errors.New("auth cookie not provided")
	}

	// throw error if the cookie doesn't exist in the cache
	cacheCookieData, cacheCookieErr := GetCacheCookie(cookie)
	if cacheCookieErr != nil {
		return nil, errors.New("invalid cookie provided")
	}

	// validate the cookie
	validatedCookie, validatedCookieErr := ValidateCookie(cookie)

	// throw error if the cookie is invalid
	if validatedCookieErr != nil {
		return nil, errors.New("invalid cookie")
	}

	// throw error for expires session
	if validatedCookie.HasExpired {
		return nil, errors.New("cookies has expired")
	}

	return cacheCookieData, nil
}
