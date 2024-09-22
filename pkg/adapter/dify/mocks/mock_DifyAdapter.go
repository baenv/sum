// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockDifyAdapter is an autogenerated mock type for the DifyAdapter type
type MockDifyAdapter struct {
	mock.Mock
}

type MockDifyAdapter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDifyAdapter) EXPECT() *MockDifyAdapter_Expecter {
	return &MockDifyAdapter_Expecter{mock: &_m.Mock}
}

// SummarizeArticle provides a mock function with given fields: url
func (_m *MockDifyAdapter) SummarizeArticle(url string) (string, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for SummarizeArticle")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDifyAdapter_SummarizeArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SummarizeArticle'
type MockDifyAdapter_SummarizeArticle_Call struct {
	*mock.Call
}

// SummarizeArticle is a helper method to define mock.On call
//   - url string
func (_e *MockDifyAdapter_Expecter) SummarizeArticle(url interface{}) *MockDifyAdapter_SummarizeArticle_Call {
	return &MockDifyAdapter_SummarizeArticle_Call{Call: _e.mock.On("SummarizeArticle", url)}
}

func (_c *MockDifyAdapter_SummarizeArticle_Call) Run(run func(url string)) *MockDifyAdapter_SummarizeArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDifyAdapter_SummarizeArticle_Call) Return(content string, err error) *MockDifyAdapter_SummarizeArticle_Call {
	_c.Call.Return(content, err)
	return _c
}

func (_c *MockDifyAdapter_SummarizeArticle_Call) RunAndReturn(run func(string) (string, error)) *MockDifyAdapter_SummarizeArticle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDifyAdapter creates a new instance of MockDifyAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDifyAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDifyAdapter {
	mock := &MockDifyAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
