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

func CmdListProposedDisableValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-disable-nodes",
		Short: "Query the list of all proposed disable validators",
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

			params := &types.QueryAllProposedDisableValidatorRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposedDisableValidatorAll(context.Background(), params)
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

func CmdShowProposedDisableValidator() *cobra.Command {
	var address string

	cmd := &cobra.Command{
		Use:   "proposed-disable-node --address [address]",
		Short: "Query proposed disable validator by address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			var res types.ProposedDisableValidator

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
				types.ProposedDisableValidatorKeyPrefix,
				types.ProposedDisableValidatorKey(addr.String()),
				&res,
			)
		},
	}

	cmd.Flags().StringVar(&address, FlagAddress, "", "Bech32 encoded validator address or owner account")

	_ = cmd.MarkFlagRequired(FlagAddress)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
