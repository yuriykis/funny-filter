package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type params struct {
	dev   string
	ip    string
	limit string
}

func validate(cmd *cobra.Command) (*params, error) {
	dev, err := cmd.Flags().GetString("dev")
	if err != nil {
		return nil, err
	}
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return nil, err
	}
	limit, err := cmd.Flags().GetString("limit")
	if err != nil {
		return nil, err
	}
	if dev == "" || ip == "" || limit == "" {
		return nil, fmt.Errorf("Not all arguments provided")
	}
	return &params{
		dev:   dev,
		ip:    ip,
		limit: limit,
	}, nil
}
