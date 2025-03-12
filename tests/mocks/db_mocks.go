package mocks

import "github.com/stretchr/testify/mock"

// Example using testify/mock
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Save(data string) interface{} {
	args := m.Called(data)
	return args.Get(0)
}

func (m *MockDatabase) Fetch(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDatabase) List(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDatabase) Update(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDatabase) Delete(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDatabase) FetchById(data string) error {
	args := m.Called(data)
	return args.Error(0)
}
