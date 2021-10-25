package keeper

import (
	"github.com/example/crypstarter/x/crypstarter/types"
)

var _ types.QueryServer = Keeper{}
