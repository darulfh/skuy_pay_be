// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "BE-Golang/model"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// CreateTransactionByUserIdRepository provides a mock function with given fields: transaction
func (_m *TransactionRepository) CreateTransactionByUserIdRepository(transaction *model.Transaction) (*model.Transaction, error) {
	ret := _m.Called(transaction)

	var r0 *model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Transaction) (*model.Transaction, error)); ok {
		return rf(transaction)
	}
	if rf, ok := ret.Get(0).(func(*model.Transaction) *model.Transaction); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTransactionsRepository provides a mock function with given fields: page, limit
func (_m *TransactionRepository) GetAllTransactionsRepository(page int, limit int) ([]*model.Transaction, error) {
	ret := _m.Called(page, limit)

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*model.Transaction, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*model.Transaction); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductDetailsByPeriodAndCustomerID provides a mock function with given fields: payload
func (_m *TransactionRepository) GetProductDetailsByPeriodAndCustomerID(payload model.GetProductDetail) (*model.Transaction, error) {
	ret := _m.Called(payload)

	var r0 *model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(model.GetProductDetail) (*model.Transaction, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(model.GetProductDetail) *model.Transaction); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(model.GetProductDetail) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByIdRepository provides a mock function with given fields: transactionID
func (_m *TransactionRepository) GetTransactionByIdRepository(transactionID string) (*model.Transaction, error) {
	ret := _m.Called(transactionID)

	var r0 *model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Transaction, error)); ok {
		return rf(transactionID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Transaction); ok {
		r0 = rf(transactionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByUserIdRepository provides a mock function with given fields: userID, productType, page, limit
func (_m *TransactionRepository) GetTransactionByUserIdRepository(userID string, productType string, page int, limit int) ([]*model.Transaction, error) {
	ret := _m.Called(userID, productType, page, limit)

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]*model.Transaction, error)); ok {
		return rf(userID, productType, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*model.Transaction); ok {
		r0 = rf(userID, productType, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(userID, productType, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByMonthRepository provides a mock function with given fields: month, year
func (_m *TransactionRepository) GetTransactionsByMonthRepository(month time.Month, year int) ([]*model.Transaction, error) {
	ret := _m.Called(month, year)

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Month, int) ([]*model.Transaction, error)); ok {
		return rf(month, year)
	}
	if rf, ok := ret.Get(0).(func(time.Month, int) []*model.Transaction); ok {
		r0 = rf(month, year)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Month, int) error); ok {
		r1 = rf(month, year)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByQueryRepository provides a mock function with given fields: query, page, limit
func (_m *TransactionRepository) GetTransactionsByQueryRepository(query string, page int, limit int) ([]*model.Transaction, error) {
	ret := _m.Called(query, page, limit)

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]*model.Transaction, error)); ok {
		return rf(query, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.Transaction); ok {
		r0 = rf(query, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(query, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByStatusQueryRepository provides a mock function with given fields: query, status, page, limit
func (_m *TransactionRepository) GetTransactionsByStatusQueryRepository(query string, status string, page int, limit int) ([]*model.Transaction, error) {
	ret := _m.Called(query, status, page, limit)

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]*model.Transaction, error)); ok {
		return rf(query, status, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*model.Transaction); ok {
		r0 = rf(query, status, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(query, status, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsPriceCountRepository provides a mock function with given fields:
func (_m *TransactionRepository) GetTransactionsPriceCountRepository() ([]*model.Transaction, error) {
	ret := _m.Called()

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Transaction, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsProductTypeRepository provides a mock function with given fields: productType, status, page, limit
func (_m *TransactionRepository) GetTransactionsProductTypeRepository(productType string, status string, page int, limit int) ([]*model.Transaction, error) {
	ret := _m.Called(productType, status, page, limit)

	var r0 []*model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]*model.Transaction, error)); ok {
		return rf(productType, status, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []*model.Transaction); ok {
		r0 = rf(productType, status, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(productType, status, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransactionByIdRepository provides a mock function with given fields: userId, transaction
func (_m *TransactionRepository) UpdateTransactionByIdRepository(userId string, transaction *model.Transaction) (*model.Transaction, error) {
	ret := _m.Called(userId, transaction)

	var r0 *model.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.Transaction) (*model.Transaction, error)); ok {
		return rf(userId, transaction)
	}
	if rf, ok := ret.Get(0).(func(string, *model.Transaction) *model.Transaction); ok {
		r0 = rf(userId, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.Transaction) error); ok {
		r1 = rf(userId, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
