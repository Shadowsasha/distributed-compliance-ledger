package dclauth

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgProposeAddAccount:
			res, err := msgServer.ProposeAddAccount(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgApproveAddAccount:
			res, err := msgServer.ApproveAddAccount(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgProposeRevokeAccount:
			res, err := msgServer.ProposeRevokeAccount(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgApproveRevokeAccount:
			res, err := msgServer.ApproveRevokeAccount(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRejectAddAccount:
			res, err := msgServer.RejectAddAccount(sdk.WrapSDKContext(ctx), msg)

			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)

			return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
