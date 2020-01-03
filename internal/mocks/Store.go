// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddCounter provides a mock function with given fields: id
func (_m *Store) AddCounter(id string) {
	_m.Called(id)
}

// AddCounterWithTTL provides a mock function with given fields: id, second
func (_m *Store) AddCounterWithTTL(id string, second int) {
	_m.Called(id, second)
}

// ClearCounter provides a mock function with given fields: id
func (_m *Store) ClearCounter(id string) {
	_m.Called(id)
}

// GetCounter provides a mock function with given fields: id
func (_m *Store) GetCounter(id string) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
