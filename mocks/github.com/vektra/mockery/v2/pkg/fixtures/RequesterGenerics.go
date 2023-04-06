// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	io "io"

	constraints "github.com/vektra/mockery/v2/pkg/fixtures/constraints"

	mock "github.com/stretchr/testify/mock"

	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// RequesterGenerics is an autogenerated mock type for the RequesterGenerics type
type RequesterGenerics[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	mock.Mock
}

type RequesterGenerics_Expecter[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	mock *mock.Mock
}

func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) EXPECT() *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{mock: &_m.Mock}
}

// GenericAnonymousStructs provides a mock function with given fields: _a0
func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericAnonymousStructs(_a0 struct{ Type1 TExternalIntf }) struct {
	Type2 test.GenericType[string, test.EmbeddedGet[int]]
} {
	ret := _m.Called(_a0)

	var r0 struct {
		Type2 test.GenericType[string, test.EmbeddedGet[int]]
	}
	if rf, ok := ret.Get(0).(func(struct{ Type1 TExternalIntf }) struct {
		Type2 test.GenericType[string, test.EmbeddedGet[int]]
	}); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(struct {
			Type2 test.GenericType[string, test.EmbeddedGet[int]]
		})
	}

	return r0
}

// RequesterGenerics_GenericAnonymousStructs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenericAnonymousStructs'
type RequesterGenerics_GenericAnonymousStructs_Call[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	*mock.Call
}

// GenericAnonymousStructs is a helper method to define mock.On call
//   - _a0 struct{Type1 TExternalIntf}
func (_e *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericAnonymousStructs(_a0 interface{}) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{Call: _e.mock.On("GenericAnonymousStructs", _a0)}
}

func (_c *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Run(run func(_a0 struct{ Type1 TExternalIntf })) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(struct{ Type1 TExternalIntf }))
	})
	return _c
}

func (_c *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Return(_a0 struct {
	Type2 test.GenericType[string, test.EmbeddedGet[int]]
}) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) RunAndReturn(run func(struct{ Type1 TExternalIntf }) struct {
	Type2 test.GenericType[string, test.EmbeddedGet[int]]
}) *RequesterGenerics_GenericAnonymousStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(run)
	return _c
}

// GenericArguments provides a mock function with given fields: _a0, _a1
func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericArguments(_a0 TAny, _a1 TComparable) (TSigned, TIntf) {
	ret := _m.Called(_a0, _a1)

	var r0 TSigned
	var r1 TIntf
	if rf, ok := ret.Get(0).(func(TAny, TComparable) (TSigned, TIntf)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(TAny, TComparable) TSigned); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(TSigned)
	}

	if rf, ok := ret.Get(1).(func(TAny, TComparable) TIntf); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(TIntf)
	}

	return r0, r1
}

// RequesterGenerics_GenericArguments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenericArguments'
type RequesterGenerics_GenericArguments_Call[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	*mock.Call
}

// GenericArguments is a helper method to define mock.On call
//   - _a0 TAny
//   - _a1 TComparable
func (_e *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericArguments(_a0 interface{}, _a1 interface{}) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{Call: _e.mock.On("GenericArguments", _a0, _a1)}
}

func (_c *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Run(run func(_a0 TAny, _a1 TComparable)) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(TAny), args[1].(TComparable))
	})
	return _c
}

func (_c *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Return(_a0 TSigned, _a1 TIntf) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) RunAndReturn(run func(TAny, TComparable) (TSigned, TIntf)) *RequesterGenerics_GenericArguments_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(run)
	return _c
}

// GenericStructs provides a mock function with given fields: _a0
func (_m *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericStructs(_a0 test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf] {
	ret := _m.Called(_a0)

	var r0 test.GenericType[TSigned, TIntf]
	if rf, ok := ret.Get(0).(func(test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf]); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(test.GenericType[TSigned, TIntf])
	}

	return r0
}

// RequesterGenerics_GenericStructs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenericStructs'
type RequesterGenerics_GenericStructs_Call[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}] struct {
	*mock.Call
}

// GenericStructs is a helper method to define mock.On call
//   - _a0 test.GenericType[TAny,TIntf]
func (_e *RequesterGenerics_Expecter[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) GenericStructs(_a0 interface{}) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	return &RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{Call: _e.mock.On("GenericStructs", _a0)}
}

func (_c *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Run(run func(_a0 test.GenericType[TAny, TIntf])) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(test.GenericType[TAny, TIntf]))
	})
	return _c
}

func (_c *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) Return(_a0 test.GenericType[TSigned, TIntf]) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]) RunAndReturn(run func(test.GenericType[TAny, TIntf]) test.GenericType[TSigned, TIntf]) *RequesterGenerics_GenericStructs_Call[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRequesterGenerics interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequesterGenerics creates a new instance of RequesterGenerics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterGenerics[TAny interface{}, TComparable comparable, TSigned constraints.Signed, TIntf test.GetInt, TExternalIntf io.Writer, TGenIntf test.GetGeneric[TSigned], TInlineType interface{ ~int | ~uint }, TInlineTypeGeneric interface {
	~int | test.GenericType[int, test.GetInt]
	comparable
}](t mockConstructorTestingTNewRequesterGenerics) *RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric] {
	mock := &RequesterGenerics[TAny, TComparable, TSigned, TIntf, TExternalIntf, TGenIntf, TInlineType, TInlineTypeGeneric]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
