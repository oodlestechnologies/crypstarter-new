package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/example/crypstarter_testnet/testutil/keeper"
	"github.com/example/crypstarter_testnet/x/crypstartertestnet/keeper"
	"github.com/example/crypstarter_testnet/x/crypstartertestnet/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CrypstartertestnetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
