package myblockchain_test

import (
	"testing"

	keepertest "github.com/cosmonaut/MyBlockchain/testutil/keeper"
	"github.com/cosmonaut/MyBlockchain/testutil/nullify"
	"github.com/cosmonaut/MyBlockchain/x/myblockchain"
	"github.com/cosmonaut/MyBlockchain/x/myblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyblockchainKeeper(t)
	myblockchain.InitGenesis(ctx, *k, genesisState)
	got := myblockchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
