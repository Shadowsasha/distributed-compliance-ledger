package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

func CmdListRevokedAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-revoked-accounts",
		Short: "list all RevokedAccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRevokedAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RevokedAccountAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowRevokedAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoked-account",
		Short: "shows a RevokedAccount",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			argAddress, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
			if err != nil {
				return err
			}

			var res types.RevokedAccount

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.RevokedAccountKeyPrefix,
				types.RevokedAccountKey(argAddress),
				&res,
			)
		},
	}

	cmd.Flags().String(FlagAddress, "", "Bech32 encoded account address")
	flags.AddQueryFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
