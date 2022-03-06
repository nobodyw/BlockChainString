package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmonaut/MyBlockchain/x/myblockchain/types"
    "github.com/cosmonaut/MyBlockchain/x/myblockchain/keeper"
    keepertest "github.com/cosmonaut/MyBlockchain/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MyblockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
