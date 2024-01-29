package storage

import (
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/onflow/flow-go/fvm/evm/types"
)

type BlockIndexer interface {
	Store(block *types.Block) error
	Get(height uint64) (*types.Block, error)
	LatestHeight() (uint64, error)
	FirstHeight() (uint64, error)
}

type ReceiptIndexer interface {
	Store(receipt *gethTypes.ReceiptForStorage) error
	Get(txID common.Hash) *gethTypes.ReceiptForStorage
}

type TransactionIndexer interface {
	Store(tx *gethTypes.Transaction) error
	Get(txID common.Hash) *gethTypes.Transaction
}
