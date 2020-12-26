package packer

import (
	"github.com/GuoxiW/fractal/core/types"
	"github.com/GuoxiW/fractal/event"
)

type Packer interface {
	//
	InsertRemoteTxPackage(pkg *types.TxPackage) error

	// pack_service
	InsertTransactions(txs types.Transactions) []error
	StartPacking(packerIndex uint32)
	StopPacking()
	IsPacking() bool
	Subscribe(ch chan<- NewPackageEvent) event.Subscription
}

type NewPackageEvent struct {
	Pkgs []*types.TxPackage
}
