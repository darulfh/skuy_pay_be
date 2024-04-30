// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "BE-Golang/model"

	mock "github.com/stretchr/testify/mock"
)

// BillerOyApiRepository is an autogenerated mock type for the BillerOyApiRepository type
type BillerOyApiRepository struct {
	mock.Mock
}

// BillInquryRepository provides a mock function with given fields: payload
func (_m *BillerOyApiRepository) BillInquryRepository(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error) {
	ret := _m.Called(payload)

	var r0 *model.OyBillerApiResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.OyBillerApi) (*model.OyBillerApiResponse, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*model.OyBillerApi) *model.OyBillerApiResponse); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OyBillerApiResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.OyBillerApi) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BillPaymentStatusRepository provides a mock function with given fields: payload
func (_m *BillerOyApiRepository) BillPaymentStatusRepository(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error) {
	ret := _m.Called(payload)

	var r0 *model.OyBillerApiResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.OyBillerApi) (*model.OyBillerApiResponse, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*model.OyBillerApi) *model.OyBillerApiResponse); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OyBillerApiResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.OyBillerApi) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PayBillRepository provides a mock function with given fields: payload
func (_m *BillerOyApiRepository) PayBillRepository(payload *model.OyBillerApi) (*model.OyBillerApiResponse, error) {
	ret := _m.Called(payload)

	var r0 *model.OyBillerApiResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.OyBillerApi) (*model.OyBillerApiResponse, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*model.OyBillerApi) *model.OyBillerApiResponse); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OyBillerApiResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.OyBillerApi) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBillerOyApiRepository creates a new instance of BillerOyApiRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBillerOyApiRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BillerOyApiRepository {
	mock := &BillerOyApiRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
