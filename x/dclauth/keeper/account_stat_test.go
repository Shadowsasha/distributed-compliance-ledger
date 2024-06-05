package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/nullify"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func createTestAccountStat(keeper *keeper.Keeper, ctx sdk.Context) types.AccountStat { //nolint:unparam
	item := types.AccountStat{}
	keeper.SetAccountStat(ctx, item)

	return item
}

func TestAccountStatGet(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	item := createTestAccountStat(keeper, ctx)
	rst, found := keeper.GetAccountStat(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestGetNextAccountNumber(t *testing.T) {
	keeper, ctx := keepertest.DclauthKeeper(t)
	require.Equal(t, uint64(0), keeper.GetNextAccountNumber(ctx))
	require.Equal(t, uint64(1), keeper.GetNextAccountNumber(ctx))
	require.Equal(t, uint64(2), keeper.GetNextAccountNumber(ctx))
}
