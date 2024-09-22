// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	logger "ask/pkg/logger"

	mock "github.com/stretchr/testify/mock"
)

// MockLogger is an autogenerated mock type for the Logger type
type MockLogger struct {
	mock.Mock
}

type MockLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLogger) EXPECT() *MockLogger_Expecter {
	return &MockLogger_Expecter{mock: &_m.Mock}
}

// AddField provides a mock function with given fields: key, value
func (_m *MockLogger) AddField(key string, value interface{}) logger.Logger {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for AddField")
	}

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func(string, interface{}) logger.Logger); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// MockLogger_AddField_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddField'
type MockLogger_AddField_Call struct {
	*mock.Call
}

// AddField is a helper method to define mock.On call
//   - key string
//   - value interface{}
func (_e *MockLogger_Expecter) AddField(key interface{}, value interface{}) *MockLogger_AddField_Call {
	return &MockLogger_AddField_Call{Call: _e.mock.On("AddField", key, value)}
}

func (_c *MockLogger_AddField_Call) Run(run func(key string, value interface{})) *MockLogger_AddField_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockLogger_AddField_Call) Return(_a0 logger.Logger) *MockLogger_AddField_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_AddField_Call) RunAndReturn(run func(string, interface{}) logger.Logger) *MockLogger_AddField_Call {
	_c.Call.Return(run)
	return _c
}

// Debug provides a mock function with given fields: msg
func (_m *MockLogger) Debug(msg string) {
	_m.Called(msg)
}

// MockLogger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type MockLogger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//   - msg string
func (_e *MockLogger_Expecter) Debug(msg interface{}) *MockLogger_Debug_Call {
	return &MockLogger_Debug_Call{Call: _e.mock.On("Debug", msg)}
}

func (_c *MockLogger_Debug_Call) Run(run func(msg string)) *MockLogger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLogger_Debug_Call) Return() *MockLogger_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Debug_Call) RunAndReturn(run func(string)) *MockLogger_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Debugf provides a mock function with given fields: msg, args
func (_m *MockLogger) Debugf(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Debugf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debugf'
type MockLogger_Debugf_Call struct {
	*mock.Call
}

// Debugf is a helper method to define mock.On call
//   - msg string
//   - args ...interface{}
func (_e *MockLogger_Expecter) Debugf(msg interface{}, args ...interface{}) *MockLogger_Debugf_Call {
	return &MockLogger_Debugf_Call{Call: _e.mock.On("Debugf",
		append([]interface{}{msg}, args...)...)}
}

func (_c *MockLogger_Debugf_Call) Run(run func(msg string, args ...interface{})) *MockLogger_Debugf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Debugf_Call) Return() *MockLogger_Debugf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Debugf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Debugf_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: err, msg
func (_m *MockLogger) Error(err error, msg string) {
	_m.Called(err, msg)
}

// MockLogger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type MockLogger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - err error
//   - msg string
func (_e *MockLogger_Expecter) Error(err interface{}, msg interface{}) *MockLogger_Error_Call {
	return &MockLogger_Error_Call{Call: _e.mock.On("Error", err, msg)}
}

func (_c *MockLogger_Error_Call) Run(run func(err error, msg string)) *MockLogger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error), args[1].(string))
	})
	return _c
}

func (_c *MockLogger_Error_Call) Return() *MockLogger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Error_Call) RunAndReturn(run func(error, string)) *MockLogger_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Errorf provides a mock function with given fields: err, msg, args
func (_m *MockLogger) Errorf(err error, msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Errorf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errorf'
type MockLogger_Errorf_Call struct {
	*mock.Call
}

// Errorf is a helper method to define mock.On call
//   - err error
//   - msg string
//   - args ...interface{}
func (_e *MockLogger_Expecter) Errorf(err interface{}, msg interface{}, args ...interface{}) *MockLogger_Errorf_Call {
	return &MockLogger_Errorf_Call{Call: _e.mock.On("Errorf",
		append([]interface{}{err, msg}, args...)...)}
}

func (_c *MockLogger_Errorf_Call) Run(run func(err error, msg string, args ...interface{})) *MockLogger_Errorf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(error), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Errorf_Call) Return() *MockLogger_Errorf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Errorf_Call) RunAndReturn(run func(error, string, ...interface{})) *MockLogger_Errorf_Call {
	_c.Call.Return(run)
	return _c
}

// Fatal provides a mock function with given fields: err, msg
func (_m *MockLogger) Fatal(err error, msg string) {
	_m.Called(err, msg)
}

// MockLogger_Fatal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatal'
type MockLogger_Fatal_Call struct {
	*mock.Call
}

// Fatal is a helper method to define mock.On call
//   - err error
//   - msg string
func (_e *MockLogger_Expecter) Fatal(err interface{}, msg interface{}) *MockLogger_Fatal_Call {
	return &MockLogger_Fatal_Call{Call: _e.mock.On("Fatal", err, msg)}
}

func (_c *MockLogger_Fatal_Call) Run(run func(err error, msg string)) *MockLogger_Fatal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error), args[1].(string))
	})
	return _c
}

func (_c *MockLogger_Fatal_Call) Return() *MockLogger_Fatal_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Fatal_Call) RunAndReturn(run func(error, string)) *MockLogger_Fatal_Call {
	_c.Call.Return(run)
	return _c
}

// Fatalf provides a mock function with given fields: err, msg, args
func (_m *MockLogger) Fatalf(err error, msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Fatalf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatalf'
type MockLogger_Fatalf_Call struct {
	*mock.Call
}

// Fatalf is a helper method to define mock.On call
//   - err error
//   - msg string
//   - args ...interface{}
func (_e *MockLogger_Expecter) Fatalf(err interface{}, msg interface{}, args ...interface{}) *MockLogger_Fatalf_Call {
	return &MockLogger_Fatalf_Call{Call: _e.mock.On("Fatalf",
		append([]interface{}{err, msg}, args...)...)}
}

func (_c *MockLogger_Fatalf_Call) Run(run func(err error, msg string, args ...interface{})) *MockLogger_Fatalf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(error), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Fatalf_Call) Return() *MockLogger_Fatalf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Fatalf_Call) RunAndReturn(run func(error, string, ...interface{})) *MockLogger_Fatalf_Call {
	_c.Call.Return(run)
	return _c
}

// Field provides a mock function with given fields: key, value
func (_m *MockLogger) Field(key string, value string) logger.Logger {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Field")
	}

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func(string, string) logger.Logger); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// MockLogger_Field_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Field'
type MockLogger_Field_Call struct {
	*mock.Call
}

// Field is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *MockLogger_Expecter) Field(key interface{}, value interface{}) *MockLogger_Field_Call {
	return &MockLogger_Field_Call{Call: _e.mock.On("Field", key, value)}
}

func (_c *MockLogger_Field_Call) Run(run func(key string, value string)) *MockLogger_Field_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockLogger_Field_Call) Return(_a0 logger.Logger) *MockLogger_Field_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_Field_Call) RunAndReturn(run func(string, string) logger.Logger) *MockLogger_Field_Call {
	_c.Call.Return(run)
	return _c
}

// Fields provides a mock function with given fields: data
func (_m *MockLogger) Fields(data logger.Fields) logger.Logger {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Fields")
	}

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func(logger.Fields) logger.Logger); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// MockLogger_Fields_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fields'
type MockLogger_Fields_Call struct {
	*mock.Call
}

// Fields is a helper method to define mock.On call
//   - data logger.Fields
func (_e *MockLogger_Expecter) Fields(data interface{}) *MockLogger_Fields_Call {
	return &MockLogger_Fields_Call{Call: _e.mock.On("Fields", data)}
}

func (_c *MockLogger_Fields_Call) Run(run func(data logger.Fields)) *MockLogger_Fields_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(logger.Fields))
	})
	return _c
}

func (_c *MockLogger_Fields_Call) Return(_a0 logger.Logger) *MockLogger_Fields_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLogger_Fields_Call) RunAndReturn(run func(logger.Fields) logger.Logger) *MockLogger_Fields_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: msg
func (_m *MockLogger) Info(msg string) {
	_m.Called(msg)
}

// MockLogger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type MockLogger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - msg string
func (_e *MockLogger_Expecter) Info(msg interface{}) *MockLogger_Info_Call {
	return &MockLogger_Info_Call{Call: _e.mock.On("Info", msg)}
}

func (_c *MockLogger_Info_Call) Run(run func(msg string)) *MockLogger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLogger_Info_Call) Return() *MockLogger_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Info_Call) RunAndReturn(run func(string)) *MockLogger_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Infof provides a mock function with given fields: msg, args
func (_m *MockLogger) Infof(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Infof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Infof'
type MockLogger_Infof_Call struct {
	*mock.Call
}

// Infof is a helper method to define mock.On call
//   - msg string
//   - args ...interface{}
func (_e *MockLogger_Expecter) Infof(msg interface{}, args ...interface{}) *MockLogger_Infof_Call {
	return &MockLogger_Infof_Call{Call: _e.mock.On("Infof",
		append([]interface{}{msg}, args...)...)}
}

func (_c *MockLogger_Infof_Call) Run(run func(msg string, args ...interface{})) *MockLogger_Infof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Infof_Call) Return() *MockLogger_Infof_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Infof_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Infof_Call {
	_c.Call.Return(run)
	return _c
}

// Warn provides a mock function with given fields: msg
func (_m *MockLogger) Warn(msg string) {
	_m.Called(msg)
}

// MockLogger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type MockLogger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//   - msg string
func (_e *MockLogger_Expecter) Warn(msg interface{}) *MockLogger_Warn_Call {
	return &MockLogger_Warn_Call{Call: _e.mock.On("Warn", msg)}
}

func (_c *MockLogger_Warn_Call) Run(run func(msg string)) *MockLogger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockLogger_Warn_Call) Return() *MockLogger_Warn_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Warn_Call) RunAndReturn(run func(string)) *MockLogger_Warn_Call {
	_c.Call.Return(run)
	return _c
}

// Warnf provides a mock function with given fields: msg, args
func (_m *MockLogger) Warnf(msg string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockLogger_Warnf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warnf'
type MockLogger_Warnf_Call struct {
	*mock.Call
}

// Warnf is a helper method to define mock.On call
//   - msg string
//   - args ...interface{}
func (_e *MockLogger_Expecter) Warnf(msg interface{}, args ...interface{}) *MockLogger_Warnf_Call {
	return &MockLogger_Warnf_Call{Call: _e.mock.On("Warnf",
		append([]interface{}{msg}, args...)...)}
}

func (_c *MockLogger_Warnf_Call) Run(run func(msg string, args ...interface{})) *MockLogger_Warnf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Warnf_Call) Return() *MockLogger_Warnf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Warnf_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Warnf_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLogger creates a new instance of MockLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLogger {
	mock := &MockLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
