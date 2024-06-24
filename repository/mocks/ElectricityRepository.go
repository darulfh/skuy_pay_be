// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "github.com/darulfh/skuy_pay_be/model"

	mock "github.com/stretchr/testify/mock"
)

// ElectricityRepository is an autogenerated mock type for the ElectricityRepository type
type ElectricityRepository struct {
	mock.Mock
}

// CreateElectricityRepository provides a mock function with given fields: electricity
func (_m *ElectricityRepository) CreateElectricityRepository(electricity *model.Electricity) (*model.Electricity, error) {
	ret := _m.Called(electricity)

	var r0 *model.Electricity
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Electricity) (*model.Electricity, error)); ok {
		return rf(electricity)
	}
	if rf, ok := ret.Get(0).(func(*model.Electricity) *model.Electricity); ok {
		r0 = rf(electricity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Electricity)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Electricity) error); ok {
		r1 = rf(electricity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteElectricityByIdRepository provides a mock function with given fields: electricityId
func (_m *ElectricityRepository) DeleteElectricityByIdRepository(electricityId string) error {
	ret := _m.Called(electricityId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(electricityId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllElectricityRepository provides a mock function with given fields: page, limit
func (_m *ElectricityRepository) GetAllElectricityRepository(page int, limit int) ([]*model.Electricity, error) {
	ret := _m.Called(page, limit)

	var r0 []*model.Electricity
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*model.Electricity, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*model.Electricity); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Electricity)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetElectricityByIdRepository provides a mock function with given fields: electricityId
func (_m *ElectricityRepository) GetElectricityByIdRepository(electricityId string) (*model.Electricity, error) {
	ret := _m.Called(electricityId)

	var r0 *model.Electricity
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Electricity, error)); ok {
		return rf(electricityId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Electricity); ok {
		r0 = rf(electricityId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Electricity)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(electricityId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetElectricityByPeriodRepository provides a mock function with given fields: customerID, period
func (_m *ElectricityRepository) GetElectricityByPeriodRepository(customerID string, period string) (*model.Electricity, error) {
	ret := _m.Called(customerID, period)

	var r0 *model.Electricity
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.Electricity, error)); ok {
		return rf(customerID, period)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.Electricity); ok {
		r0 = rf(customerID, period)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Electricity)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(customerID, period)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateElectricityByIdRepository provides a mock function with given fields: electricityId, electricity
func (_m *ElectricityRepository) UpdateElectricityByIdRepository(electricityId string, electricity *model.Electricity) (*model.Electricity, error) {
	ret := _m.Called(electricityId, electricity)

	var r0 *model.Electricity
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.Electricity) (*model.Electricity, error)); ok {
		return rf(electricityId, electricity)
	}
	if rf, ok := ret.Get(0).(func(string, *model.Electricity) *model.Electricity); ok {
		r0 = rf(electricityId, electricity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Electricity)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.Electricity) error); ok {
		r1 = rf(electricityId, electricity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewElectricityRepository creates a new instance of ElectricityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewElectricityRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ElectricityRepository {
	mock := &ElectricityRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
