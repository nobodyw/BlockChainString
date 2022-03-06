package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/MyBlockchain/testutil/keeper"
	"github.com/cosmonaut/MyBlockchain/x/myblockchain/keeper"
	"github.com/cosmonaut/MyBlockchain/x/myblockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MyblockchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
