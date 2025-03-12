package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/backend-boilerplate-template/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"strings"
	"time"
)

type CustomCookie struct {
	Profile        models.ProfileDto
	CookieName     string
	CookieValue    string
	ExpirationTime time.Time
}

func GenerateCookie(profile *models.Profile) (*CustomCookie, error) {
	// Set the token in the session map, along with the session information
	newCustomCookie := &CustomCookie{
		Profile:        models.ProfileFromDomain(profile),
		CookieName:     "auth_cookie",
		CookieValue:    uuid.NewString(),
		ExpirationTime: time.Now().Add(24 * time.Hour), // expires in 24 hours 24 * time.Hour
	}

	// set the newly created cookie to redis and expires in 3 minutes after expiration time * 1443 | 1443*time.Minute
	if setCacheErr := RedisSetCache(
		newCustomCookie.CookieValue, newCustomCookie, 1443*time.Minute); setCacheErr != nil {
		return nil, setCacheErr
	}

	return newCustomCookie, nil
}

func GetCookieDomain(origin string) string {
	var domain = ""

	if strings.Contains(origin, "localhost") {
		domain = "localhost"
	} else if strings.Contains(origin, "propati") {
		domain = "app.propati.xyz"
	}

	return domain
}

type validateCookieStruct struct {
	IsValid    bool
	CacheData  map[string]interface{}
	HasExpired bool
}

// we'll use this method later to determine if the session has expired
func ValidateCookie(key string) (validateCookieStruct, error) {
	cookie, err := GetCacheCookie(key)

	if err != nil {
		return validateCookieStruct{
			IsValid:    false,
			CacheData:  nil,
			HasExpired: false,
		}, err
	}

	// parse the time to string
	expTime, expErr := time.Parse("2006-01-02T15:04:05.999Z", cookie["ExpirationTime"].(string))

	if expErr != nil {
		return validateCookieStruct{
			IsValid:    true,
			CacheData:  nil,
			HasExpired: false,
		}, errors.New(fmt.Sprintf("Time parsing error: %s", expErr))
	}

	// returns true if the cookie has expired
	if time.Now().After(expTime) {
		return validateCookieStruct{
			IsValid:    false,
			CacheData:  cookie,
			HasExpired: true,
		}, nil
	}

	return validateCookieStruct{
		IsValid:    false,
		CacheData:  cookie,
		HasExpired: false,
	}, nil
}

func GetCacheCookie(key string) (map[string]interface{}, error) {
	// get the cache cookie from redis cache
	redisCache, redisCacheErr := RedisGetCache(key)

	// throw error if redisCacheErr exists
	if redisCacheErr != nil {
		return nil, redisCacheErr
	}

	var cacheObject interface{}
	unmarshallErr := json.Unmarshal([]byte(redisCache), &cacheObject)

	if unmarshallErr != nil {
		return nil, unmarshallErr
	}

	cacheData, _ := cacheObject.(map[string]interface{})

	return cacheData, nil
}

func ResetCookie(c fiber.Ctx) {
	domain := GetCookieDomain(c.Get("Origin"))

	// write cookie to the client
	c.Cookie(&fiber.Cookie{
		Name:     "auth_cookie",
		Value:    "",
		MaxAge:   0,
		Path:     "/",
		Domain:   domain,
		Secure:   false,
		HTTPOnly: true,
	})
}
