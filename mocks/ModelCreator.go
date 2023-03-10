// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ModelCreator is an autogenerated mock type for the ModelCreator type
type ModelCreator struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *ModelCreator) Execute() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type mockConstructorTestingTNewModelCreator interface {
	mock.TestingT
	Cleanup(func())
}

// NewModelCreator creates a new instance of ModelCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewModelCreator(t mockConstructorTestingTNewModelCreator) *ModelCreator {
	mock := &ModelCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
