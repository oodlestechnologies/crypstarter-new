package keeper

import (
	"github.com/example/crypstarter_testnet/x/crypstartertestnet/types"
)

var _ types.QueryServer = Keeper{}
