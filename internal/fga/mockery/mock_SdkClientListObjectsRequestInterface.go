// Code generated by mockery v2.40.1. DO NOT EDIT.

package client

import (
	context "context"

	client "github.com/openfga/go-sdk/client"

	mock "github.com/stretchr/testify/mock"

	openfga "github.com/openfga/go-sdk"
)

// MockSdkClientListObjectsRequestInterface is an autogenerated mock type for the SdkClientListObjectsRequestInterface type
type MockSdkClientListObjectsRequestInterface struct {
	mock.Mock
}

type MockSdkClientListObjectsRequestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSdkClientListObjectsRequestInterface) EXPECT() *MockSdkClientListObjectsRequestInterface_Expecter {
	return &MockSdkClientListObjectsRequestInterface_Expecter{mock: &_m.Mock}
}

// Body provides a mock function with given fields: body
func (_m *MockSdkClientListObjectsRequestInterface) Body(body client.ClientListObjectsRequest) client.SdkClientListObjectsRequestInterface {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 client.SdkClientListObjectsRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientListObjectsRequest) client.SdkClientListObjectsRequestInterface); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientListObjectsRequestInterface)
		}
	}

	return r0
}

// MockSdkClientListObjectsRequestInterface_Body_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Body'
type MockSdkClientListObjectsRequestInterface_Body_Call struct {
	*mock.Call
}

// Body is a helper method to define mock.On call
//   - body client.ClientListObjectsRequest
func (_e *MockSdkClientListObjectsRequestInterface_Expecter) Body(body interface{}) *MockSdkClientListObjectsRequestInterface_Body_Call {
	return &MockSdkClientListObjectsRequestInterface_Body_Call{Call: _e.mock.On("Body", body)}
}

func (_c *MockSdkClientListObjectsRequestInterface_Body_Call) Run(run func(body client.ClientListObjectsRequest)) *MockSdkClientListObjectsRequestInterface_Body_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientListObjectsRequest))
	})
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_Body_Call) Return(_a0 client.SdkClientListObjectsRequestInterface) *MockSdkClientListObjectsRequestInterface_Body_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_Body_Call) RunAndReturn(run func(client.ClientListObjectsRequest) client.SdkClientListObjectsRequestInterface) *MockSdkClientListObjectsRequestInterface_Body_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields:
func (_m *MockSdkClientListObjectsRequestInterface) Execute() (*openfga.ListObjectsResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *openfga.ListObjectsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*openfga.ListObjectsResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *openfga.ListObjectsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openfga.ListObjectsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSdkClientListObjectsRequestInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockSdkClientListObjectsRequestInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockSdkClientListObjectsRequestInterface_Expecter) Execute() *MockSdkClientListObjectsRequestInterface_Execute_Call {
	return &MockSdkClientListObjectsRequestInterface_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockSdkClientListObjectsRequestInterface_Execute_Call) Run(run func()) *MockSdkClientListObjectsRequestInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_Execute_Call) Return(_a0 *openfga.ListObjectsResponse, _a1 error) *MockSdkClientListObjectsRequestInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_Execute_Call) RunAndReturn(run func() (*openfga.ListObjectsResponse, error)) *MockSdkClientListObjectsRequestInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthorizationModelIdOverride provides a mock function with given fields:
func (_m *MockSdkClientListObjectsRequestInterface) GetAuthorizationModelIdOverride() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorizationModelIdOverride")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorizationModelIdOverride'
type MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call struct {
	*mock.Call
}

// GetAuthorizationModelIdOverride is a helper method to define mock.On call
func (_e *MockSdkClientListObjectsRequestInterface_Expecter) GetAuthorizationModelIdOverride() *MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call {
	return &MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call{Call: _e.mock.On("GetAuthorizationModelIdOverride")}
}

func (_c *MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call) Run(run func()) *MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call) Return(_a0 *string) *MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call) RunAndReturn(run func() *string) *MockSdkClientListObjectsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(run)
	return _c
}

// GetBody provides a mock function with given fields:
func (_m *MockSdkClientListObjectsRequestInterface) GetBody() *client.ClientListObjectsRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 *client.ClientListObjectsRequest
	if rf, ok := ret.Get(0).(func() *client.ClientListObjectsRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientListObjectsRequest)
		}
	}

	return r0
}

// MockSdkClientListObjectsRequestInterface_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type MockSdkClientListObjectsRequestInterface_GetBody_Call struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *MockSdkClientListObjectsRequestInterface_Expecter) GetBody() *MockSdkClientListObjectsRequestInterface_GetBody_Call {
	return &MockSdkClientListObjectsRequestInterface_GetBody_Call{Call: _e.mock.On("GetBody")}
}

func (_c *MockSdkClientListObjectsRequestInterface_GetBody_Call) Run(run func()) *MockSdkClientListObjectsRequestInterface_GetBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetBody_Call) Return(_a0 *client.ClientListObjectsRequest) *MockSdkClientListObjectsRequestInterface_GetBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetBody_Call) RunAndReturn(run func() *client.ClientListObjectsRequest) *MockSdkClientListObjectsRequestInterface_GetBody_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockSdkClientListObjectsRequestInterface) GetContext() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockSdkClientListObjectsRequestInterface_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSdkClientListObjectsRequestInterface_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockSdkClientListObjectsRequestInterface_Expecter) GetContext() *MockSdkClientListObjectsRequestInterface_GetContext_Call {
	return &MockSdkClientListObjectsRequestInterface_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockSdkClientListObjectsRequestInterface_GetContext_Call) Run(run func()) *MockSdkClientListObjectsRequestInterface_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetContext_Call) Return(_a0 context.Context) *MockSdkClientListObjectsRequestInterface_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetContext_Call) RunAndReturn(run func() context.Context) *MockSdkClientListObjectsRequestInterface_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockSdkClientListObjectsRequestInterface) GetOptions() *client.ClientListObjectsOptions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOptions")
	}

	var r0 *client.ClientListObjectsOptions
	if rf, ok := ret.Get(0).(func() *client.ClientListObjectsOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientListObjectsOptions)
		}
	}

	return r0
}

// MockSdkClientListObjectsRequestInterface_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockSdkClientListObjectsRequestInterface_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockSdkClientListObjectsRequestInterface_Expecter) GetOptions() *MockSdkClientListObjectsRequestInterface_GetOptions_Call {
	return &MockSdkClientListObjectsRequestInterface_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockSdkClientListObjectsRequestInterface_GetOptions_Call) Run(run func()) *MockSdkClientListObjectsRequestInterface_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetOptions_Call) Return(_a0 *client.ClientListObjectsOptions) *MockSdkClientListObjectsRequestInterface_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_GetOptions_Call) RunAndReturn(run func() *client.ClientListObjectsOptions) *MockSdkClientListObjectsRequestInterface_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: options
func (_m *MockSdkClientListObjectsRequestInterface) Options(options client.ClientListObjectsOptions) client.SdkClientListObjectsRequestInterface {
	ret := _m.Called(options)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 client.SdkClientListObjectsRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientListObjectsOptions) client.SdkClientListObjectsRequestInterface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientListObjectsRequestInterface)
		}
	}

	return r0
}

// MockSdkClientListObjectsRequestInterface_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type MockSdkClientListObjectsRequestInterface_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - options client.ClientListObjectsOptions
func (_e *MockSdkClientListObjectsRequestInterface_Expecter) Options(options interface{}) *MockSdkClientListObjectsRequestInterface_Options_Call {
	return &MockSdkClientListObjectsRequestInterface_Options_Call{Call: _e.mock.On("Options", options)}
}

func (_c *MockSdkClientListObjectsRequestInterface_Options_Call) Run(run func(options client.ClientListObjectsOptions)) *MockSdkClientListObjectsRequestInterface_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientListObjectsOptions))
	})
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_Options_Call) Return(_a0 client.SdkClientListObjectsRequestInterface) *MockSdkClientListObjectsRequestInterface_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListObjectsRequestInterface_Options_Call) RunAndReturn(run func(client.ClientListObjectsOptions) client.SdkClientListObjectsRequestInterface) *MockSdkClientListObjectsRequestInterface_Options_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSdkClientListObjectsRequestInterface creates a new instance of MockSdkClientListObjectsRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSdkClientListObjectsRequestInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSdkClientListObjectsRequestInterface {
	mock := &MockSdkClientListObjectsRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}