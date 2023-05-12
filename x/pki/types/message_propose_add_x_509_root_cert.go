package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	pkitypes "github.com/zigbee-alliance/distributed-compliance-ledger/types/pki"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

const TypeMsgProposeAddX509RootCert = "propose_add_x_509_root_cert"

var _ sdk.Msg = &MsgProposeAddX509RootCert{}

func NewMsgProposeAddX509RootCert(signer string, cert string, info string) *MsgProposeAddX509RootCert {
	return &MsgProposeAddX509RootCert{
		Signer: signer,
		Cert:   cert,
		Info:   info,
		Time:   time.Now().Unix(),
	}
}

func (msg *MsgProposeAddX509RootCert) Route() string {
	return pkitypes.RouterKey
}

func (msg *MsgProposeAddX509RootCert) Type() string {
	return TypeMsgProposeAddX509RootCert
}

func (msg *MsgProposeAddX509RootCert) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

func (msg *MsgProposeAddX509RootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeAddX509RootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	err = validator.Validate(msg)
	if err != nil {
		return err
	}

	return nil
}
