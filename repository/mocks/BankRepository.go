// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "BE-Golang/model"

	mock "github.com/stretchr/testify/mock"
)

// BankRepository is an autogenerated mock type for the BankRepository type
type BankRepository struct {
	mock.Mock
}

// CreateBankRepository provides a mock function with given fields: bank
func (_m *BankRepository) CreateBankRepository(bank *model.Bank) (*model.Bank, error) {
	ret := _m.Called(bank)

	var r0 *model.Bank
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Bank) (*model.Bank, error)); ok {
		return rf(bank)
	}
	if rf, ok := ret.Get(0).(func(*model.Bank) *model.Bank); ok {
		r0 = rf(bank)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Bank)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Bank) error); ok {
		r1 = rf(bank)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBankByIdRepository provides a mock function with given fields: bankId
func (_m *BankRepository) DeleteBankByIdRepository(bankId string) error {
	ret := _m.Called(bankId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bankId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBanksRepository provides a mock function with given fields: page, limit
func (_m *BankRepository) GetAllBanksRepository(page int, limit int) ([]*model.Bank, error) {
	ret := _m.Called(page, limit)

	var r0 []*model.Bank
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*model.Bank, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*model.Bank); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Bank)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBankByIdRepository provides a mock function with given fields: bankId
func (_m *BankRepository) GetBankByIdRepository(bankId string) (*model.Bank, error) {
	ret := _m.Called(bankId)

	var r0 *model.Bank
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Bank, error)); ok {
		return rf(bankId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Bank); ok {
		r0 = rf(bankId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Bank)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bankId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBankByIdRepository provides a mock function with given fields: bankId, bank
func (_m *BankRepository) UpdateBankByIdRepository(bankId string, bank *model.Bank) (*model.Bank, error) {
	ret := _m.Called(bankId, bank)

	var r0 *model.Bank
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.Bank) (*model.Bank, error)); ok {
		return rf(bankId, bank)
	}
	if rf, ok := ret.Get(0).(func(string, *model.Bank) *model.Bank); ok {
		r0 = rf(bankId, bank)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Bank)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.Bank) error); ok {
		r1 = rf(bankId, bank)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBankRepository creates a new instance of BankRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBankRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BankRepository {
	mock := &BankRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
