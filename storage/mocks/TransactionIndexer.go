// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	common "github.com/onflow/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	models "github.com/onflow/flow-evm-gateway/models"
)

// TransactionIndexer is an autogenerated mock type for the TransactionIndexer type
type TransactionIndexer struct {
	mock.Mock
}

// Get provides a mock function with given fields: ID
func (_m *TransactionIndexer) Get(ID common.Hash) (models.Transaction, error) {
	ret := _m.Called(ID)

	var r0 models.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Hash) (models.Transaction, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(common.Hash) models.Transaction); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: tx
func (_m *TransactionIndexer) Store(tx models.Transaction) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Transaction) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTransactionIndexer interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionIndexer creates a new instance of TransactionIndexer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionIndexer(t mockConstructorTestingTNewTransactionIndexer) *TransactionIndexer {
	mock := &TransactionIndexer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
