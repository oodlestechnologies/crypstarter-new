package wasm_test

import (
	"testing"

	keepertest "github.com/example/crypstarter/testutil/keeper"
	"github.com/example/crypstarter/testutil/nullify"
	"github.com/example/crypstarter/x/wasm"
	"github.com/example/crypstarter/x/wasm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WasmKeeper(t)
	wasm.InitGenesis(ctx, *k, genesisState)
	got := wasm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
