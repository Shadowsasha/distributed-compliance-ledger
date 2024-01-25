package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func SimulateMsgApproveAddAccount() simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgApproveAddAccount{
			Signer: simAccount.Address.String(),
		}

		// TODO: Handling the ApproveAddAccount simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ApproveAddAccount simulation not implemented"), nil, nil
	}
}
