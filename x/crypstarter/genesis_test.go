package crypstarter_test

import (
	"testing"

	keepertest "github.com/example/crypstarter/testutil/keeper"
	"github.com/example/crypstarter/x/crypstarter"
	"github.com/example/crypstarter/x/crypstarter/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrypstarterKeeper(t)
	crypstarter.InitGenesis(ctx, *k, genesisState)
	got := crypstarter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
