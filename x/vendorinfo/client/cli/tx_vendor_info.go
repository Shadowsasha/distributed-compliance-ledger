package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/common"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func CmdCreateVendorInfo() *cobra.Command {
	var (
		vid                  int32
		vendorName           string
		companyLegalName     string
		companyPreferredName string
		vendorLandingPageURL string
		schemaVersion        uint32
	)

	cmd := &cobra.Command{
		Use:   "add-vendor",
		Short: "Add a new VendorInfo",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVendorInfo(
				clientCtx.GetFromAddress().String(),
				vid,
				vendorName,
				companyLegalName,
				companyPreferredName,
				vendorLandingPageURL,
				schemaVersion,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0,
		"Vendor ID")
	cmd.Flags().StringVarP(&vendorName, FlagVendorName, FlagVendorNameShortcut,
		"", "Vendor Name")
	cmd.Flags().StringVarP(&companyLegalName, FlagCompanyLegalName, FlagCompanyLegalNameShortcut,
		"", "Company Legal Name")
	cmd.Flags().StringVarP(&companyPreferredName, FlagCompanyPreferredName, FlagCompanyPreferredNameShortcut,
		"", "Company Preferred Name")
	cmd.Flags().StringVarP(&vendorLandingPageURL, FlagVendorLandingPageURL, FlagVendorLandingPageURLShortcut,
		"", "Landing Page URL for the Vendor")
	cli.AddTxFlagsToCmd(cmd)
	cmd.Flags().Uint32Var(&schemaVersion, common.FlagSchemaVersion, 0, "Schema version")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagVendorName)
	_ = cmd.MarkFlagRequired(FlagCompanyLegalName)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func CmdUpdateVendorInfo() *cobra.Command {
	var (
		vid                  int32
		vendorName           string
		companyLegalName     string
		companyPreferredName string
		vendorLandingPageURL string
		schemaVersion        uint32
	)

	cmd := &cobra.Command{
		Use:   "update-vendor",
		Short: "Update a VendorInfo",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVendorInfo(
				clientCtx.GetFromAddress().String(),
				vid,
				vendorName,
				companyLegalName,
				companyPreferredName,
				vendorLandingPageURL,
				schemaVersion,
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}

			return err
		},
	}

	cmd.Flags().Int32Var(&vid, FlagVID, 0,
		"Vendor ID")
	cmd.Flags().StringVarP(&vendorName, FlagVendorName, FlagVendorNameShortcut,
		"", "Vendor Name")
	cmd.Flags().StringVarP(&companyLegalName, FlagCompanyLegalName, FlagCompanyLegalNameShortcut,
		"", "Company Legal Name")
	cmd.Flags().StringVarP(&companyPreferredName, FlagCompanyPreferredName, FlagCompanyPreferredNameShortcut,
		"", "Company Preferred Name")
	cmd.Flags().StringVarP(&vendorLandingPageURL, FlagVendorLandingPageURL, FlagVendorLandingPageURLShortcut,
		"", "Landing Page URL for the Vendor")
	cmd.Flags().Uint32Var(&schemaVersion, common.FlagSchemaVersion, 0, "Schema version")

	cli.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
