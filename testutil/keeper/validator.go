package keeper

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	dclauthkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func ValidatorKeeper(tb testing.TB, dclauthK *dclauthkeeper.Keeper) (*keeper.Keeper, sdk.Context) {
	tb.Helper()
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	// TODO issue 99: might be not the best solution
	if dclauthK != nil {
		stateStore.MountStoreWithDB(dclauthK.StoreKey(), storetypes.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(dclauthK.MemKey(), storetypes.StoreTypeMemory, nil)
	}

	require.NoError(tb, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)

	k := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
		dclauthK,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}

func DefaultValidator() types.Validator {
	v, _ := types.NewValidator(
		sdk.ValAddress(testconstants.Address1),
		testconstants.PubKey1,
		types.Description{Moniker: testconstants.ProductName},
	)

	return v
}

type TestSetup struct {
	Ctx             sdk.Context
	ValidatorKeeper keeper.Keeper
	DclauthKeeper   dclauthkeeper.Keeper
}

func Setup(t *testing.T) TestSetup {
	t.Helper()
	dclauthK, _ := DclauthKeeper(t)
	k, ctx := ValidatorKeeper(t, dclauthK)

	setup := TestSetup{
		Ctx:             ctx,
		ValidatorKeeper: *k,
		DclauthKeeper:   *dclauthK,
	}

	return setup
}

func StoreTwoValidators(setup TestSetup) (types.Validator, types.Validator) {
	validator1, _ := types.NewValidator(
		sdk.ValAddress(testconstants.ValidatorAddress1),
		testconstants.ValidatorPubKey1,
		types.Description{Moniker: "Validator 1"},
	)
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator1)

	validator2, _ := types.NewValidator(
		sdk.ValAddress(testconstants.ValidatorAddress2),
		testconstants.ValidatorPubKey2,
		types.Description{Moniker: "Validator 2"},
	)
	setup.ValidatorKeeper.SetValidator(setup.Ctx, validator2)

	return validator1, validator2
}
