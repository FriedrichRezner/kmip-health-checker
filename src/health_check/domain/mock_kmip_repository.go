// Code generated by mockery. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockKMIPRepository is an autogenerated mock type for the KMIPRepository type
type MockKMIPRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx
func (_m *MockKMIPRepository) Create(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields: ctx, id
func (_m *MockKMIPRepository) Destroy(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Destroy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockKMIPRepository creates a new instance of MockKMIPRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockKMIPRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockKMIPRepository {
	mock := &MockKMIPRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}