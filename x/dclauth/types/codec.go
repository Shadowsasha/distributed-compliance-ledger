package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProposeAddAccount{}, "dclauth/ProposeAddAccount", nil)
	cdc.RegisterConcrete(&MsgApproveAddAccount{}, "dclauth/ApproveAddAccount", nil)
	cdc.RegisterConcrete(&MsgProposeRevokeAccount{}, "dclauth/ProposeRevokeAccount", nil)
	cdc.RegisterConcrete(&MsgApproveRevokeAccount{}, "dclauth/ApproveRevokeAccount", nil)
	cdc.RegisterConcrete(&MsgRejectAddAccount{}, "dclauth/RejectAddAccount", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeAddAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveAddAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeRevokeAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveRevokeAccount{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRejectAddAccount{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc) //nolint:nosnakecase
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()

	cryptocodec.RegisterInterfaces(ModuleCdc.InterfaceRegistry())
}
