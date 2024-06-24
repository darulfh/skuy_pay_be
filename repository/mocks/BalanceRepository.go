// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "github.com/darulfh/skuy_pay_be/model"

	mock "github.com/stretchr/testify/mock"
)

// BalanceRepository is an autogenerated mock type for the BalanceRepository type
type BalanceRepository struct {
	mock.Mock
}

// GetVaIdRepository provides a mock function with given fields: patnerId
func (_m *BalanceRepository) GetVaIdRepository(patnerId string) (*model.VaNumber, error) {
	ret := _m.Called(patnerId)

	var r0 *model.VaNumber
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.VaNumber, error)); ok {
		return rf(patnerId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.VaNumber); ok {
		r0 = rf(patnerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VaNumber)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(patnerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVaNumber provides a mock function with given fields: vaNumber
func (_m *BalanceRepository) GetVaNumber(vaNumber string) (*model.VaNumber, error) {
	ret := _m.Called(vaNumber)

	var r0 *model.VaNumber
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.VaNumber, error)); ok {
		return rf(vaNumber)
	}
	if rf, ok := ret.Get(0).(func(string) *model.VaNumber); ok {
		r0 = rf(vaNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VaNumber)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(vaNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVaNumberByUserId provides a mock function with given fields: userId
func (_m *BalanceRepository) GetVaNumberByUserId(userId string) (*model.VaNumber, error) {
	ret := _m.Called(userId)

	var r0 *model.VaNumber
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.VaNumber, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.VaNumber); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VaNumber)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertVaRepository provides a mock function with given fields: payload
func (_m *BalanceRepository) InsertVaRepository(payload *model.VaNumber) (*model.VaNumber, error) {
	ret := _m.Called(payload)

	var r0 *model.VaNumber
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.VaNumber) (*model.VaNumber, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*model.VaNumber) *model.VaNumber); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.VaNumber)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.VaNumber) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBalanceRepository creates a new instance of BalanceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBalanceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BalanceRepository {
	mock := &BalanceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
