package crypstartertestnet_test

import (
	"testing"

	keepertest "github.com/example/crypstarter_testnet/testutil/keeper"
	"github.com/example/crypstarter_testnet/x/crypstartertestnet"
	"github.com/example/crypstarter_testnet/x/crypstartertestnet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrypstartertestnetKeeper(t)
	crypstartertestnet.InitGenesis(ctx, *k, genesisState)
	got := crypstartertestnet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
