// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	common "github.com/onflow/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	models "github.com/onflow/flow-evm-gateway/models"

	pebble "github.com/cockroachdb/pebble"
)

// ReceiptIndexer is an autogenerated mock type for the ReceiptIndexer type
type ReceiptIndexer struct {
	mock.Mock
}

// BloomsForBlockRange provides a mock function with given fields: start, end
func (_m *ReceiptIndexer) BloomsForBlockRange(start uint64, end uint64) ([]*models.BloomsHeight, error) {
	ret := _m.Called(start, end)

	if len(ret) == 0 {
		panic("no return value specified for BloomsForBlockRange")
	}

	var r0 []*models.BloomsHeight
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) ([]*models.BloomsHeight, error)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(uint64, uint64) []*models.BloomsHeight); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.BloomsHeight)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByBlockHeight provides a mock function with given fields: height
func (_m *ReceiptIndexer) GetByBlockHeight(height uint64) ([]*models.Receipt, error) {
	ret := _m.Called(height)

	if len(ret) == 0 {
		panic("no return value specified for GetByBlockHeight")
	}

	var r0 []*models.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]*models.Receipt, error)); ok {
		return rf(height)
	}
	if rf, ok := ret.Get(0).(func(uint64) []*models.Receipt); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTransactionID provides a mock function with given fields: ID
func (_m *ReceiptIndexer) GetByTransactionID(ID common.Hash) (*models.Receipt, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for GetByTransactionID")
	}

	var r0 *models.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Hash) (*models.Receipt, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(common.Hash) *models.Receipt); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: receipts, batch
func (_m *ReceiptIndexer) Store(receipts []*models.Receipt, batch *pebble.Batch) error {
	ret := _m.Called(receipts, batch)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]*models.Receipt, *pebble.Batch) error); ok {
		r0 = rf(receipts, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewReceiptIndexer creates a new instance of ReceiptIndexer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReceiptIndexer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReceiptIndexer {
	mock := &ReceiptIndexer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
