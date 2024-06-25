package requester

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/onflow/flow-go-sdk"
	gethTypes "github.com/onflow/go-ethereum/core/types"
	"github.com/rs/zerolog"

	errs "github.com/onflow/flow-evm-gateway/api/errors"
)

const evmErrorRegex = `evm_error=(\d+)`

// todo this is a simple implementation of the transaction pool that is mostly used
// to track the status of submitted transaction, but transactions will always be submitted
// right away, future improvements can make it so the transactions are collected in the pool
// and after submitted based on different strategies.

type TxPool struct {
	logger zerolog.Logger
	client *CrossSporkClient
	pool   *sync.Map
	// todo add a broadcaster for pending transaction streaming
	// todo add methods to inspect transaction pool state
}

func NewTxPool(client *CrossSporkClient, logger zerolog.Logger) *TxPool {
	return &TxPool{
		logger: logger.With().Str("component", "tx-pool").Logger(),
		client: client,
		pool:   &sync.Map{},
	}
}

// Send flow transaction that executes EVM run function which takes in the encoded EVM transaction.
// The flow transaction status is awaited and an error is returned in case of a failure in submission,
// or an EVM validation error.
// Until the flow transaction is sealed the transaction will stay in the transaction pool marked as pending.
func (t *TxPool) Send(
	ctx context.Context,
	flowTx *flow.Transaction,
	evmTx *gethTypes.Transaction,
) error {
	if err := t.client.SendTransaction(ctx, *flowTx); err != nil {
		return err
	}

	// add to pool
	t.pool.Store(evmTx.Hash(), evmTx)

	const fetchInterval = time.Millisecond * 500
	const fetchTimeout = time.Minute * 3
	ticker := time.NewTicker(fetchInterval)
	timeout := time.NewTimer(fetchTimeout)

	defer func() {
		t.pool.Delete(evmTx.Hash())
		timeout.Stop()
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			res, err := t.client.GetTransactionResult(ctx, flowTx.ID())
			if err != nil {
				return fmt.Errorf("failed to retrieve flow transaction result %s: %w", flowTx.ID(), err)
			}
			// wait until transaction is sealed
			if res.Status < flow.TransactionStatusSealed {
				continue
			}

			if res.Error != nil {
				t.logger.Error().Err(res.Error).
					Str("flow-id", flowTx.ID().String()).
					Str("evm-id", evmTx.Hash().Hex()).
					Msg("flow transaction error")

				if err, ok := parseInvalidError(res.Error); ok {
					return err
				}

				// hide specific cause since it's an implementation issue
				return fmt.Errorf("failed to submit flow evm transaction %s", evmTx.Hash())
			}

			return nil
		case <-timeout.C:
			err := fmt.Errorf("failed to retrieve evm transaction result: %s, flow id: %s", evmTx.Hash(), flowTx.ID())
			t.logger.Error().Err(err).Msg("failed to get result")
			return err
		}
	}
}

// this will extract the evm specific error from the Flow transaction error message
// the run.cdc script panics with the evm specific error as the message which we
// extract and return to the client. Any error returned that is evm specific
// is a validation error due to assert statement in the run.cdc script.
func parseInvalidError(err error) (error, bool) {
	r := regexp.MustCompile(evmErrorRegex)
	matches := r.FindStringSubmatch(err.Error())
	if len(matches) != 2 {
		return nil, false
	}

	return fmt.Errorf("%w: %s", errs.ErrInvalid, matches[1]), true
}
