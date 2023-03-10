// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto " hery-ciaputra/assignment-05-golang-backend/dto"
	mock "github.com/stretchr/testify/mock"
)

// RegisterService is an autogenerated mock type for the RegisterService type
type RegisterService struct {
	mock.Mock
}

// Register provides a mock function with given fields: req
func (_m *RegisterService) Register(req *dto.RegisterReq) (*dto.RegisterRes, error) {
	ret := _m.Called(req)

	var r0 *dto.RegisterRes
	if rf, ok := ret.Get(0).(func(*dto.RegisterReq) *dto.RegisterRes); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.RegisterRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.RegisterReq) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRegisterService interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegisterService creates a new instance of RegisterService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegisterService(t mockConstructorTestingTNewRegisterService) *RegisterService {
	mock := &RegisterService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
