package api

import (
	"context"
	"github.com/GuoxiW/fractal/chain"
	"github.com/GuoxiW/fractal/common"
	"github.com/GuoxiW/fractal/core/config"
	"github.com/GuoxiW/fractal/core/pool"
	"github.com/GuoxiW/fractal/core/types"
	"github.com/GuoxiW/fractal/dbwrapper"
	"github.com/GuoxiW/fractal/ftl/sync"
	"github.com/GuoxiW/fractal/keys"
	"github.com/GuoxiW/fractal/logbloom/bloomquery"
	"github.com/GuoxiW/fractal/packer"
	"math/big"
)

type fractal interface {
	IsMining() bool
	StartMining() error
	StopMining()
	MiningKeyManager() *keys.MiningKeyManager
	Coinbase() common.Address

	Config() *config.Config
	Packer() packer.Packer
	BlockChain() *chain.BlockChain
	TxPool() pool.Pool
	Signer() types.Signer
	GasPrice() *big.Int
	GetPoolTransactions() types.Transactions

	FtlVersion() int

	Synchronizer() *sync.Synchronizer

	ChainDb() dbwrapper.Database
	GetBlock(ctx context.Context, fullHash common.Hash) *types.Block
	GetBlockStr(blockStr string) *types.Block
	GetReceipts(ctx context.Context, blockHash common.Hash) types.Receipts
	GetLogs(ctx context.Context, blockHash common.Hash) [][]*types.Log

	GetMainBranchBlock(height uint64) (*types.BlockHeader, error)
	BloomRequestsReceiver() chan chan *bloomquery.Retrieval
}
