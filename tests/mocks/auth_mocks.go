package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Example using testify/mock
type MockAuthHandlers struct {
	mock.Mock
}

type LoginRequestParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *MockAuthHandlers) Register(ctx string, request []byte) []byte {
	args := m.Called(ctx, request)

	var obj []byte
	if args.Get(0) != nil {
		obj = args.Get(0).([]byte)
	}

	return obj
}

func (m *MockAuthHandlers) MockAuth(ctx string, request []byte) map[string]interface{} {
	args := m.Called(ctx, request)

	var obj map[string]interface{}
	if args.Get(0) != nil {
		obj = args.Get(0).(map[string]interface{})
	}

	return obj
}
