// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// JwtUtilityInterface is an autogenerated mock type for the JwtUtilityInterface type
type JwtUtilityInterface struct {
	mock.Mock
}

// DecodToken provides a mock function with given fields: token
func (_m *JwtUtilityInterface) DecodToken(token *jwt.Token) float64 {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for DecodToken")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func(*jwt.Token) float64); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// DecodTokenV2 provides a mock function with given fields: c
func (_m *JwtUtilityInterface) DecodTokenV2(c echo.Context) (uint, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for DecodTokenV2")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context) (uint, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(echo.Context) uint); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(echo.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateJwt provides a mock function with given fields: id
func (_m *JwtUtilityInterface) GenerateJwt(id uint) (string, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GenerateJwt")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJwtUtilityInterface creates a new instance of JwtUtilityInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtUtilityInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtUtilityInterface {
	mock := &JwtUtilityInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
