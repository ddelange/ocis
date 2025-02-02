// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	store "go-micro.dev/v4/store"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

type Store_Expecter struct {
	mock *mock.Mock
}

func (_m *Store) EXPECT() *Store_Expecter {
	return &Store_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with no fields
func (_m *Store) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Store_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Store_Expecter) Close() *Store_Close_Call {
	return &Store_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Store_Close_Call) Run(run func()) *Store_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Store_Close_Call) Return(_a0 error) *Store_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Close_Call) RunAndReturn(run func() error) *Store_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: key, opts
func (_m *Store) Delete(key string, opts ...store.DeleteOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...store.DeleteOption) error); ok {
		r0 = rf(key, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Store_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - key string
//   - opts ...store.DeleteOption
func (_e *Store_Expecter) Delete(key interface{}, opts ...interface{}) *Store_Delete_Call {
	return &Store_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{key}, opts...)...)}
}

func (_c *Store_Delete_Call) Run(run func(key string, opts ...store.DeleteOption)) *Store_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]store.DeleteOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(store.DeleteOption)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Store_Delete_Call) Return(_a0 error) *Store_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Delete_Call) RunAndReturn(run func(string, ...store.DeleteOption) error) *Store_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: _a0
func (_m *Store) Init(_a0 ...store.Option) error {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...store.Option) error); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type Store_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - _a0 ...store.Option
func (_e *Store_Expecter) Init(_a0 ...interface{}) *Store_Init_Call {
	return &Store_Init_Call{Call: _e.mock.On("Init",
		append([]interface{}{}, _a0...)...)}
}

func (_c *Store_Init_Call) Run(run func(_a0 ...store.Option)) *Store_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]store.Option, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(store.Option)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Store_Init_Call) Return(_a0 error) *Store_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Init_Call) RunAndReturn(run func(...store.Option) error) *Store_Init_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: opts
func (_m *Store) List(opts ...store.ListOption) ([]string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(...store.ListOption) ([]string, error)); ok {
		return rf(opts...)
	}
	if rf, ok := ret.Get(0).(func(...store.ListOption) []string); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(...store.ListOption) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type Store_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - opts ...store.ListOption
func (_e *Store_Expecter) List(opts ...interface{}) *Store_List_Call {
	return &Store_List_Call{Call: _e.mock.On("List",
		append([]interface{}{}, opts...)...)}
}

func (_c *Store_List_Call) Run(run func(opts ...store.ListOption)) *Store_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]store.ListOption, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(store.ListOption)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Store_List_Call) Return(_a0 []string, _a1 error) *Store_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_List_Call) RunAndReturn(run func(...store.ListOption) ([]string, error)) *Store_List_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with no fields
func (_m *Store) Options() store.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 store.Options
	if rf, ok := ret.Get(0).(func() store.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(store.Options)
	}

	return r0
}

// Store_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type Store_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
func (_e *Store_Expecter) Options() *Store_Options_Call {
	return &Store_Options_Call{Call: _e.mock.On("Options")}
}

func (_c *Store_Options_Call) Run(run func()) *Store_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Store_Options_Call) Return(_a0 store.Options) *Store_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Options_Call) RunAndReturn(run func() store.Options) *Store_Options_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: key, opts
func (_m *Store) Read(key string, opts ...store.ReadOption) ([]*store.Record, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 []*store.Record
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...store.ReadOption) ([]*store.Record, error)); ok {
		return rf(key, opts...)
	}
	if rf, ok := ret.Get(0).(func(string, ...store.ReadOption) []*store.Record); ok {
		r0 = rf(key, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*store.Record)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...store.ReadOption) error); ok {
		r1 = rf(key, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type Store_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - key string
//   - opts ...store.ReadOption
func (_e *Store_Expecter) Read(key interface{}, opts ...interface{}) *Store_Read_Call {
	return &Store_Read_Call{Call: _e.mock.On("Read",
		append([]interface{}{key}, opts...)...)}
}

func (_c *Store_Read_Call) Run(run func(key string, opts ...store.ReadOption)) *Store_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]store.ReadOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(store.ReadOption)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Store_Read_Call) Return(_a0 []*store.Record, _a1 error) *Store_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_Read_Call) RunAndReturn(run func(string, ...store.ReadOption) ([]*store.Record, error)) *Store_Read_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with no fields
func (_m *Store) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Store_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type Store_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *Store_Expecter) String() *Store_String_Call {
	return &Store_String_Call{Call: _e.mock.On("String")}
}

func (_c *Store_String_Call) Run(run func()) *Store_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Store_String_Call) Return(_a0 string) *Store_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_String_Call) RunAndReturn(run func() string) *Store_String_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: r, opts
func (_m *Store) Write(r *store.Record, opts ...store.WriteOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, r)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*store.Record, ...store.WriteOption) error); ok {
		r0 = rf(r, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type Store_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - r *store.Record
//   - opts ...store.WriteOption
func (_e *Store_Expecter) Write(r interface{}, opts ...interface{}) *Store_Write_Call {
	return &Store_Write_Call{Call: _e.mock.On("Write",
		append([]interface{}{r}, opts...)...)}
}

func (_c *Store_Write_Call) Run(run func(r *store.Record, opts ...store.WriteOption)) *Store_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]store.WriteOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(store.WriteOption)
			}
		}
		run(args[0].(*store.Record), variadicArgs...)
	})
	return _c
}

func (_c *Store_Write_Call) Return(_a0 error) *Store_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Write_Call) RunAndReturn(run func(*store.Record, ...store.WriteOption) error) *Store_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
