package keeper

import (
	"github.com/cosmonaut/MyBlockchain/x/myblockchain/types"
)

var _ types.QueryServer = Keeper{}
