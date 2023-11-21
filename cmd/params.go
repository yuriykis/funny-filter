package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// params is a struct for validating CLI arguments
// dev is a network interface and is used only for bandwidth limit
type params struct {
	dev   string
	ip    string
	limit string
}

func validateBandwidthParams(cmd *cobra.Command) (*params, error) {
	dev, err := cmd.Flags().GetString("dev")
	if err != nil {
		return nil, err
	}
	ip, limit, err := validateParams(cmd)
	if err != nil {
		return nil, err
	}
	return &params{
		dev:   dev,
		ip:    ip,
		limit: limit,
	}, nil
}

func validatePacketsParams(cmd *cobra.Command) (*params, error) {
	ip, limit, err := validateParams(cmd)
	if err != nil {
		return nil, err
	}
	return &params{
		ip:    ip,
		limit: limit,
	}, nil
}

func validateParams(cmd *cobra.Command) (string, string, error) {
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return "", "", err
	}
	limit, err := cmd.Flags().GetString("limit")
	if err != nil {
		return "", "", err
	}
	if ip == "" || limit == "" {
		return "", "", fmt.Errorf("Not all arguments provided")
	}
	return ip, limit, nil
}
