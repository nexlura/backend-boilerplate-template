package tests

import (
	"bytes"
	"encoding/json"
	"github.com/backend-boilerplate-template/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	// Define Fiber app.
	app := fiber.New()
	type expectedStruct struct {
		Id        string `json:"id"`
		Token     string `json:"token"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		TimeZone  string `json:"time_zone"`
	}

	requestData := []byte(`{
				"email": "user@test.com",
				"password_hash": "123456",
				"first_name": "Test",
				"last_name": "User",
				"time_zone": "GMT"
			}`)

	tests := []struct {
		description  string
		route        string
		requestData  []byte
		expectedCode int
		responseData interface{}
	}{
		// Test case 1
		{
			description:  "return object on success",
			route:        "/register",
			expectedCode: 200,
			responseData: []byte(`{
				"id": "1",
				"token": "xxxxx",
				"email":     "user@test.com",
				"first_name": "Test",
				"last_name":  "User",
				"time_zone":  "GMT"
			}`),
		},
		// Test case 2
		{
			description:  "return error on invalid data",
			route:        "/register",
			expectedCode: 500,
			responseData: []byte(`{"error": "internal server error"}`),
		},
	}
	// instantiate the mocks
	authMockHandler := new(mocks.MockAuthHandlers)

	// test http method
	app.Post("/register", func(c *fiber.Ctx) error {
		// invoke the register func
		res := authMockHandler.Register(mock.Anything, requestData)

		var ex interface{}
		if err := json.Unmarshal(res, &ex); err != nil {
			panic(err)
		}

		return c.JSON(ex)
	})

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			authMockHandler.On("Register", mock.Anything, requestData).Return(tt.responseData)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(requestData))

			res, _ := app.Test(req, 1)
			body, _ := io.ReadAll(res.Body)

			// convert the response body to map
			bodyOut := expectedStruct{}
			json.Unmarshal(body, &bodyOut)

			// assertions
			assert.Equal(t, bodyOut.FirstName, "Test")
			assert.Equal(t, tests[0].expectedCode, 200)
			assert.Equal(t, tests[1].expectedCode, 500)
			assert.Equal(t, tests[1].responseData, []byte(`{"error": "internal server error"}`))

			authMockHandler.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {
	// Define Fiber app.
	app := fiber.New()

	tests := []struct {
		description  string
		route        string
		requestData  []byte
		expectedCode int
		responseData map[string]interface{}
	}{
		// Test case 1
		{
			description:  "return success on login",
			route:        "/login",
			expectedCode: 200,
			requestData: []byte(`{
				"email": "user@test.com",
				"password_hash": "123456"
			}`),
			responseData: map[string]interface{}{
				"token":      "xxxxx",
				"email":      "user@test.com",
				"first_name": "Test",
				"last_name":  "User",
				"time_zone":  "GMT",
			},
		},
		// Test case 2
		{
			description:  "return error on login",
			route:        "/login",
			expectedCode: 500,
			requestData: []byte(`{
				"email": "user@test.com",
				"password_hash": "1234"
			}`),
			responseData: map[string]interface{}{"error": "invalid credentials"},
		},
	}

	// instantiate the mocks
	authMockHandler := new(mocks.MockAuthHandlers)

	// Create route with POST method for test
	app.Post("/login", func(c *fiber.Ctx) error {
		// get the request body
		body := c.Body()

		// make the request to the mock func
		res := authMockHandler.MockAuth(mock.Anything, body)

		return c.JSON(res)
	})

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			authMockHandler.On("MockAuth", mock.Anything, tt.requestData).Return(tt.responseData)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(tt.requestData))

			res, _ := app.Test(req, 1)

			responseBody, _ := io.ReadAll(res.Body)

			// convert the response body to map
			var output map[string]interface{}
			json.Unmarshal(responseBody, &output)

			assert.Equal(t, tt.responseData, output)
		})
	}
}

func TestForgot(t *testing.T) {
	// Define Fiber app.
	app := fiber.New()

	tests := []struct {
		description  string
		route        string
		requestData  []byte
		expectedCode int
		responseData map[string]interface{}
	}{
		// Test case 1
		{
			description:  "return success on request",
			route:        "/forgot-password",
			expectedCode: 200,
			requestData:  []byte(`{"email": "user@test.com"}`),
			responseData: map[string]interface{}{"success": "reset email sent."},
		},
		// Test case 2
		{
			description:  "return error on invalid email request",
			route:        "/forgot-password",
			expectedCode: 500,
			requestData:  []byte(`{"email": "user1@test.com"}`),
			responseData: map[string]interface{}{"error": "invalid credentials"},
		},
	}

	// instantiate the mocks
	authMockHandler := new(mocks.MockAuthHandlers)

	// Create route with POST method for test
	app.Post("/forgot-password", func(c *fiber.Ctx) error {
		// get the request body
		body := c.Body()

		// make the request to the mock func
		res := authMockHandler.MockAuth(mock.Anything, body)

		return c.JSON(res)
	})

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			authMockHandler.On("MockAuth", mock.Anything, tt.requestData).Return(tt.responseData)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest("POST", tt.route, bytes.NewBuffer(tt.requestData))

			res, _ := app.Test(req, 1)

			responseBody, _ := io.ReadAll(res.Body)

			// convert the response body to map
			var output map[string]interface{}
			json.Unmarshal(responseBody, &output)

			assert.Equal(t, tt.responseData, output)
		})
	}
}

func TestResetPassword(t *testing.T) {
	// Define Fiber app.
	app := fiber.New()

	tests := []struct {
		description  string
		route        string
		requestData  []byte
		responseData map[string]interface{}
	}{
		// Test case 1
		{
			description: "return success on request",
			route:       "/reset-password",
			requestData: []byte(`{
				"token": "xxxxx",
				"password": "123456",
				"password_confirmation": "123456"
			}`),
			responseData: map[string]interface{}{"success": "password reset"},
		},
		// Test case 2
		{
			description: "return error on password mismatch request",
			route:       "/reset-password",
			requestData: []byte(`{
				"token": "xxxxx",
				"password": "123456",
				"password_confirmation": "12345"
			}`),
			responseData: map[string]interface{}{"error": "password mismatch"},
		},
		// Test case 3
		{
			description: "return error on token passed",
			route:       "/reset-password",
			requestData: []byte(`{
				"token": "",
				"password": "123456",
				"password_confirmation": "123456"
			}`),
			responseData: map[string]interface{}{"error": "could not update client"},
		},
	}

	// instantiate the mocks
	authMockHandler := new(mocks.MockAuthHandlers)

	// Create route with POST method for test
	app.Patch("/reset-password", func(c *fiber.Ctx) error {
		// get the request body
		body := c.Body()

		// make the request to the mock func
		res := authMockHandler.MockAuth(mock.Anything, body)

		return c.JSON(res)
	})

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			authMockHandler.On("MockAuth", mock.Anything, tt.requestData).Return(tt.responseData)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest("PATCH", tt.route, bytes.NewBuffer(tt.requestData))

			res, _ := app.Test(req, 1)

			responseBody, _ := io.ReadAll(res.Body)

			// convert the response body to map
			var output map[string]interface{}
			json.Unmarshal(responseBody, &output)

			assert.Equal(t, tt.responseData, output)
		})
	}
}
