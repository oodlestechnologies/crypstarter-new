package keeper

import (
	"github.com/example/crypstarter/x/wasm/types"
)

var _ types.QueryServer = Keeper{}
