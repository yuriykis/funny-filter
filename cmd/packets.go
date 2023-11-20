package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuriykis/funny-filter/internal/filter"
)

func NewPacketsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "packets",
		Short: "Packets limit",
		Long:  "Packets limit",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(NewSetPacketsCmd())
	cmd.AddCommand(NewUnsetPacketsCmd())
	return cmd
}

func NewSetPacketsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set packets limit",
		Long:  "Set packets limit",
		RunE:  parseSetPacketsCmd,
	}
	setPacketsFlags(cmd)
	return cmd
}

func NewUnsetPacketsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unset",
		Short: "Unset packets limit",
		Long:  "Unset packets limit",
		RunE:  parseUnsetPacketsCmd,
	}
	setPacketsFlags(cmd)
	return cmd
}

func parseSetPacketsCmd(cmd *cobra.Command, args []string) error {
	ps, err := validate(cmd)
	if err != nil {
		return err
	}
	p, err := filter.NewPacketsLimit(ps.ip, ps.limit)
	if err != nil {
		return err
	}
	if err := p.Apply(); err != nil {
		return err
	}
	return nil
}

func parseUnsetPacketsCmd(cmd *cobra.Command, args []string) error {
	ps, err := validate(cmd)
	if err != nil {
		return err
	}
	p, err := filter.NewPacketsLimit(ps.ip, ps.limit)
	if err != nil {
		return err
	}
	if err := p.Unset(); err != nil {
		return err
	}
	return nil
}
