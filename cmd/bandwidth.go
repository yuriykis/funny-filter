package cmd

import (
	"runtime"

	"github.com/spf13/cobra"
	"github.com/yuriykis/funny-filter/internal/filter"
)

func NewBandwidthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bandwidth",
		Short: "Bandwidth limit",
		Long:  "Bandwidth limit",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(NewSetBandwidthCmd())
	cmd.AddCommand(NewUnsetBandwidthCmd())
	return cmd
}

func NewSetBandwidthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set bandwidth limit",
		Long:  "Set bandwidth limit",
		RunE:  parseSetBandwidthCmd,
	}
	setBandwidthFlags(cmd)
	return cmd
}

func NewUnsetBandwidthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unset",
		Short: "Unset bandwidth limit",
		Long:  "Unset bandwidth limit",
		RunE:  parseUnsetBandwidthCmd,
	}
	setBandwidthFlags(cmd)
	return cmd
}

func parseSetBandwidthCmd(cmd *cobra.Command, args []string) error {
	ps, err := validateBandwidthParams(cmd)
	if err != nil {
		return err
	}
	b, err := filter.NewBandwidthLimit(runtime.GOOS, ps.dev, ps.ip, ps.limit)
	if err != nil {
		return err
	}
	if err := b.Set(); err != nil {
		return err
	}
	return nil
}

func parseUnsetBandwidthCmd(cmd *cobra.Command, args []string) error {
	ps, err := validateBandwidthParams(cmd)
	if err != nil {
		return err
	}
	b, err := filter.NewBandwidthLimit(runtime.GOOS, ps.dev, ps.ip, ps.limit)
	if err != nil {
		return err
	}
	if err := b.Unset(); err != nil {
		return err
	}
	return nil
}
