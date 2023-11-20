package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ff",
		Short: "funny-filter (ff) is network filter CLI tool for Linux",
		Long:  "funny-filter (ff) is network filter CLI tool for Linux",
		RunE:  parseCmd,
	}
	setFlags(cmd)
	return cmd
}

func Execute() error {
	if err := newRootCmd().Execute(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().String(
		"dev",
		"enp0s5",
		"Network interface to apply filter",
	)
	cmd.Flags().String(
		"ip",
		"",
		"IP address to apply filter",
	)
	// type of filter: bandwidth or packets
	cmd.Flags().String(
		"type",
		"bandwidth",
		"Type of filter: bandwidth or packets",
	)
	cmd.Flags().String(
		"limit",
		"",
		"Limit of bandwidth or packets",
	)
	cmd.Flags().String(
		"action",
		"set",
		"Action to apply: set or unset",
	)
}
