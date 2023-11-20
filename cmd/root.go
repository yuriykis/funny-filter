package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ff",
		Short: "funny-filter (ff) is a network filter CLI tool for Linux",
		Long:  "funny-filter (ff) is a network filter CLI tool for Linux",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(NewBandwidthCmd())
	cmd.AddCommand(NewPacketsCmd())
	return cmd
}

func Execute() error {
	if err := newRootCmd().Execute(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
