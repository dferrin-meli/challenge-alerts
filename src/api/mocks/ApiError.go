// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	common "challenge/alerts/src/api/application/common"

	mock "github.com/stretchr/testify/mock"
)

// ApiError is an autogenerated mock type for the ApiError type
type ApiError struct {
	mock.Mock
}

// Cause provides a mock function with given fields:
func (_m *ApiError) Cause() common.CauseList {
	ret := _m.Called()

	var r0 common.CauseList
	if rf, ok := ret.Get(0).(func() common.CauseList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.CauseList)
		}
	}

	return r0
}

// Code provides a mock function with given fields:
func (_m *ApiError) Code() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *ApiError) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *ApiError) Message() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *ApiError) Status() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewApiError interface {
	mock.TestingT
	Cleanup(func())
}

// NewApiError creates a new instance of ApiError. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApiError(t mockConstructorTestingTNewApiError) *ApiError {
	mock := &ApiError{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
