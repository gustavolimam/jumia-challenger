// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	model "github.com/jumia-challenger/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ListCustomers provides a mock function with given fields:
func (_m *Repository) ListCustomers() ([]model.Customer, error) {
	ret := _m.Called()

	var r0 []model.Customer
	if rf, ok := ret.Get(0).(func() []model.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
