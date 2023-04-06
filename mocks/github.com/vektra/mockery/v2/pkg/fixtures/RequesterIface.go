// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// RequesterIface is an autogenerated mock type for the RequesterIface type
type RequesterIface struct {
	mock.Mock
}

type RequesterIface_Expecter struct {
	mock *mock.Mock
}

func (_m *RequesterIface) EXPECT() *RequesterIface_Expecter {
	return &RequesterIface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields:
func (_m *RequesterIface) Get() io.Reader {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	return r0
}

// RequesterIface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RequesterIface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *RequesterIface_Expecter) Get() *RequesterIface_Get_Call {
	return &RequesterIface_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *RequesterIface_Get_Call) Run(run func()) *RequesterIface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequesterIface_Get_Call) Return(_a0 io.Reader) *RequesterIface_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequesterIface_Get_Call) RunAndReturn(run func() io.Reader) *RequesterIface_Get_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRequesterIface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequesterIface creates a new instance of RequesterIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterIface(t mockConstructorTestingTNewRequesterIface) *RequesterIface {
	mock := &RequesterIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
