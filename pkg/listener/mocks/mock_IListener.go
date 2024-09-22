// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockIListener is an autogenerated mock type for the IListener type
type MockIListener struct {
	mock.Mock
}

type MockIListener_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIListener) EXPECT() *MockIListener_Expecter {
	return &MockIListener_Expecter{mock: &_m.Mock}
}

// End provides a mock function with given fields:
func (_m *MockIListener) End() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for End")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIListener_End_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'End'
type MockIListener_End_Call struct {
	*mock.Call
}

// End is a helper method to define mock.On call
func (_e *MockIListener_Expecter) End() *MockIListener_End_Call {
	return &MockIListener_End_Call{Call: _e.mock.On("End")}
}

func (_c *MockIListener_End_Call) Run(run func()) *MockIListener_End_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIListener_End_Call) Return(_a0 error) *MockIListener_End_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIListener_End_Call) RunAndReturn(run func() error) *MockIListener_End_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields:
func (_m *MockIListener) Register() {
	_m.Called()
}

// MockIListener_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockIListener_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
func (_e *MockIListener_Expecter) Register() *MockIListener_Register_Call {
	return &MockIListener_Register_Call{Call: _e.mock.On("Register")}
}

func (_c *MockIListener_Register_Call) Run(run func()) *MockIListener_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIListener_Register_Call) Return() *MockIListener_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIListener_Register_Call) RunAndReturn(run func()) *MockIListener_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockIListener) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIListener_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockIListener_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockIListener_Expecter) Start() *MockIListener_Start_Call {
	return &MockIListener_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockIListener_Start_Call) Run(run func()) *MockIListener_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIListener_Start_Call) Return(_a0 error) *MockIListener_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIListener_Start_Call) RunAndReturn(run func() error) *MockIListener_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIListener creates a new instance of MockIListener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIListener(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIListener {
	mock := &MockIListener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
