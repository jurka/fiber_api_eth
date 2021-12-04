package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type Group struct {
	Name    string
	Indexes []*big.Int
}

type Index struct {
	Id                *big.Int
	Name              string
	EthPriceInWei     *big.Int
	UsdPriceInCents   *big.Int
	UsdCapitalization *big.Int
	PercentageChange  *big.Int
}

type BlockInfo types.Block
