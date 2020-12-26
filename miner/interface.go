// Copyright 2018 The go-fractal Authors
// This file is part of the go-fractal library.

// Package miner contains implementations for block mining strategy.
package miner

import (
	"github.com/GuoxiW/fractal/common"
	"github.com/GuoxiW/fractal/core/types"
	"github.com/GuoxiW/fractal/event"
)

type Miner interface {
	// Start begins the miner loop.
	Start()

	// Stop terminates the miner loop.
	Stop()

	Close()

	GetCoinbase() common.Address

	SetCoinbase(coinbase common.Address)

	IsMining() bool

	SubscribeNewMinedBlockEvent(ch chan<- types.NewMinedBlockEvent) event.Subscription
}
