package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func CmdListDisabledValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-disabled-nodes",
		Short: "Query the list of all disabled validators",
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

			params := &types.QueryAllDisabledValidatorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DisabledValidatorAll(context.Background(), params)
			if cli.IsKeyNotFoundRPCError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForListQueries)
			}
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

func CmdShowDisabledValidator() *cobra.Command {
	var address string

	cmd := &cobra.Command{
		Use:   "disabled-node --address [address]",
		Short: "Query disabled validator by address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var res types.DisabledValidator

			addr, err := sdk.ValAddressFromBech32(address)
			if err != nil {
				owner, err2 := sdk.AccAddressFromBech32(address)
				if err2 != nil {
					return err2
				}
				addr = sdk.ValAddress(owner)
			}

			return cli.QueryWithProof(
				clientCtx,
				types.StoreKey,
				types.DisabledValidatorKeyPrefix,
				types.DisabledValidatorKey(addr.String()),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&address, FlagAddress, "", "Bech32 encoded validator address or owner account")

	_ = cmd.MarkFlagRequired(FlagAddress)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
