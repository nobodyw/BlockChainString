package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/MyBlockchain/testutil/keeper"
	"github.com/cosmonaut/MyBlockchain/x/myblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MyblockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
