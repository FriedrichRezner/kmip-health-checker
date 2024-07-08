// Code generated by mockery. DO NOT EDIT.

package infrastructure

import (
	context "context"

	kmip "github.com/gemalto/kmip-go"
	mock "github.com/stretchr/testify/mock"
)

// MockKMIPClient is an autogenerated mock type for the KMIPClient type
type MockKMIPClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, msg
func (_m *MockKMIPClient) Create(ctx context.Context, msg kmip.RequestMessage) (*kmip.CreateResponsePayload, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *kmip.CreateResponsePayload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, kmip.RequestMessage) (*kmip.CreateResponsePayload, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, kmip.RequestMessage) *kmip.CreateResponsePayload); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kmip.CreateResponsePayload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, kmip.RequestMessage) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields: ctx, msg
func (_m *MockKMIPClient) Destroy(ctx context.Context, msg kmip.RequestMessage) (*kmip.DestroyResponsePayload, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for Destroy")
	}

	var r0 *kmip.DestroyResponsePayload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, kmip.RequestMessage) (*kmip.DestroyResponsePayload, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, kmip.RequestMessage) *kmip.DestroyResponsePayload); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kmip.DestroyResponsePayload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, kmip.RequestMessage) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockKMIPClient creates a new instance of MockKMIPClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockKMIPClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockKMIPClient {
	mock := &MockKMIPClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}