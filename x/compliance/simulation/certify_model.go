package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	dclcompltypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/compliance"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/keeper"
)

func SimulateMsgCertifyModel(
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &dclcompltypes.MsgCertifyModel{
			Signer: simAccount.Address.String(),
		}

		// TODO: Handling the CertifyModel simulation

		return simtypes.NoOpMsg(dclcompltypes.ModuleName, msg.Type(), "CertifyModel simulation not implemented"), nil, nil
	}
}
