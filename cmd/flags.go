package cmd

import "github.com/spf13/cobra"

func setBandwidthFlags(cmd *cobra.Command) {
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
	cmd.Flags().String(
		"limit",
		"",
		"Limit of bandwidth or packets",
	)
}

func setPacketsFlags(cmd *cobra.Command) {
	cmd.Flags().String(
		"ip",
		"",
		"IP address to apply filter",
	)
	cmd.Flags().String(
		"limit",
		"",
		"Limit of bandwidth or packets",
	)
}
