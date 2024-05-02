package validator

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateValidator:
			res, err := msgServer.CreateValidator(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgProposeDisableValidator:
			res, err := msgServer.ProposeDisableValidator(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgApproveDisableValidator:
			res, err := msgServer.ApproveDisableValidator(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDisableValidator:
			res, err := msgServer.DisableValidator(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgEnableValidator:
			res, err := msgServer.EnableValidator(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRejectDisableValidator:
			res, err := msgServer.RejectDisableValidator(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)

			return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
