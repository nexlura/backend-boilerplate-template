package utilities

import (
	"errors"
	"fmt"
)

func ValidatePasswordLength(password string) error {
	if len(password) < 4 {
		return errors.New("password much be at least six characters long")
	}

	return nil
}

func ValidatePasswordsLength(password, confirmPassword string) error {
	if len(password) < 6 && len(confirmPassword) < 6 {
		return errors.New("passwords much be at least six characters long")
	}

	return nil
}

func ValidatePasswordsMatch(password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("passwords mismatch")
	}

	return nil
}

func ValidateIfRequestDataRequired(requestData map[string]interface{}) (error, string) {
	for key, item := range requestData {
		if item == "" {
			return errors.New(fmt.Sprintf("The field: %v is required!", key)), "Checks Failed"
		}
	}

	return nil, "Checks Passed"
}

//func CallConcurrently(c *fiber.Ctx, db startup.Database) error {
//
//	var wg sync.WaitGroup
//	var err error
//	var propertyAmenitiesListResult []models.AmenitiesDto
//	var propertyDetailResult models.PropertiesDto
//
//	wg.Add(3)
//
//	go func() {
//		defer wg.Done()
//		var result1 []models.AmenitiesDto
//		result1, err = GetPropertyAmenities(c, db)
//		if err == nil {
//			propertyAmenitiesListResult = result1
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		var result3 models.PropertiesDto
//		result3, err = GetPropertyById(c, db)
//		if err == nil {
//			propertyDetailResult = result3
//		}
//	}()
//
//	wg.Wait()
//
//	if err != nil {
//		return responses.BadRequestError(c, err.Error())
//	}
//
//	// assign the values gotten from the concurrent call
//	propertyDetailResult.Amenities = propertyAmenitiesListResult
//
//	return responses.ResponseOKWithData(c, propertyDetailResult, "property details")
//}
