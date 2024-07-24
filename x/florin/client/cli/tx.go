package cli

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noble-assets/florin/x/florin/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Transactions commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetBlacklistTxCmd())

	cmd.AddCommand(TxAcceptOwnership())
	cmd.AddCommand(TxAddAdminAccount())
	cmd.AddCommand(TxAddSystemAccount())
	cmd.AddCommand(TxAllowDenom())
	cmd.AddCommand(TxBurn())
	cmd.AddCommand(TxMint())
	cmd.AddCommand(TxRemoveAdminAccount())
	cmd.AddCommand(TxRemoveSystemAccount())
	cmd.AddCommand(TxSetMaxMintAllowance())
	cmd.AddCommand(TxSetMintAllowance())
	cmd.AddCommand(TxTransferOwnership())

	return cmd
}

func TxAcceptOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-ownership [denom]",
		Short: "Accept ownership of a specific denom",
		Long:  "Accept ownership of a specific denom, assuming there is an pending ownership transfer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgAcceptOwnership{
				Denom:  args[0],
				Signer: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxAddAdminAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-admin-account [denom] [account]",
		Short: "Add an admin account for a specific denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgAddAdminAccount{
				Denom:   args[0],
				Signer:  clientCtx.GetFromAddress().String(),
				Account: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxAddSystemAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-system-account [denom] [account]",
		Short: "Add a system account for a specific denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgAddSystemAccount{
				Denom:   args[0],
				Signer:  clientCtx.GetFromAddress().String(),
				Account: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxAllowDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allow-denom [denom] [owner]",
		Short: "Allow a new denom with an initial owner",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgAllowDenom{
				Signer: clientCtx.GetFromAddress().String(),
				Denom:  args[0],
				Owner:  args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [denom] [from] [amount]",
		Short: "Transaction that burns a specific denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgBurn{
				Denom:  args[0],
				Signer: clientCtx.GetFromAddress().String(),
				From:   args[1],
				Amount: amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [denom] [to] [amount]",
		Short: "Transaction that mints a specific denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgMint{
				Denom:  args[0],
				Signer: clientCtx.GetFromAddress().String(),
				To:     args[1],
				Amount: amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxRemoveAdminAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-admin-account [denom] [account]",
		Short: "Remove an admin account for a specific denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveAdminAccount{
				Denom:   args[0],
				Signer:  clientCtx.GetFromAddress().String(),
				Account: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxRemoveSystemAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-system-account [denom] [account]",
		Short: "Remove a system account for a specific denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveSystemAccount{
				Denom:   args[0],
				Signer:  clientCtx.GetFromAddress().String(),
				Account: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxSetMaxMintAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-max-mint-allowance [denom] [amount]",
		Short: "Set the max mint allowance for a specific denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgSetMaxMintAllowance{
				Denom:  args[0],
				Signer: clientCtx.GetFromAddress().String(),
				Amount: amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxSetMintAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-mint-allowance [denom] [account] [amount]",
		Short: "Set the mint allowance of a system account for a specific denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return errors.New("invalid amount")
			}

			msg := &types.MsgSetMintAllowance{
				Denom:   args[0],
				Signer:  clientCtx.GetFromAddress().String(),
				Account: args[1],
				Amount:  amount,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func TxTransferOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-ownership [denom] [new-owner]",
		Short: "Transfer ownership of a specific denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgTransferOwnership{
				Denom:    args[0],
				Signer:   clientCtx.GetFromAddress().String(),
				NewOwner: args[1],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
