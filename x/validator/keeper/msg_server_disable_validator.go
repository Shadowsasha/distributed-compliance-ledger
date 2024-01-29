package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func (k msgServer) DisableValidator(goCtx context.Context, msg *types.MsgDisableValidator) (*types.MsgDisableValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid Address: (%s)", err)
	}

	// check if message creator has enough rights to propose disable validator
	if !k.dclauthKeeper.HasRole(ctx, sdk.AccAddress(creatorAddr), types.EnableDisableValidatorRole) {
		return nil, errors.Wrapf(sdkerrors.ErrUnauthorized,
			"Disable validator transaction should be signed by an account with the %s role",
			types.EnableDisableValidatorRole,
		)
	}

	// check if validator exists
	isFound := k.Keeper.IsValidatorPresent(ctx, creatorAddr)
	if !isFound {
		return nil, sdkstakingtypes.ErrNoValidatorFound
	}

	// check if disabled validator exists
	_, isFound = k.GetDisabledValidator(ctx, msg.Creator)
	if isFound {
		return nil, types.NewErrDisabledValidatorAlreadyExists(msg.Creator)
	}

	disabledValidator := types.DisabledValidator{
		Address:             creatorAddr.String(),
		Creator:             sdk.AccAddress(creatorAddr).String(),
		DisabledByNodeAdmin: true,
	}

	// Disable validator
	validator, _ := k.GetValidator(ctx, creatorAddr)
	k.Jail(ctx, validator, "disabled by a node admin")

	// store disabled validator
	k.SetDisabledValidator(ctx, disabledValidator)

	return &types.MsgDisableValidatorResponse{}, nil
}
