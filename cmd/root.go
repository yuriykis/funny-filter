package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ff",
	Short: "funny-filter (ff) is network filter CLI tool for Linux",
	Long:  "funny-filter (ff) is network filter CLI tool for Linux",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
		fmt.Println("funny-filter")
	},
}

func init() {
	rootCmd.Flags().String(
		"dev",
		"enp0s5",
		"Network interface to apply filter",
	)
	rootCmd.Flags().String(
		"ip",
		"",
		"IP address to apply filter",
	)
	// type of filter: bandwidth or packets
	rootCmd.Flags().String(
		"type",
		"bandwidth",
		"Type of filter: bandwidth or packets",
	)
	rootCmd.Flags().String(
		"limit",
		"",
		"Limit of bandwidth or packets",
	)
	rootCmd.Flags().String(
		"action",
		"set",
		"Action to apply: set or unset",
	)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
