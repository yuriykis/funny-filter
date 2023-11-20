package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yuriykis/funny-filter/internal/filter"
)

const (
	limitTypeBandwidth = "bandwidth"
	limitTypePackets   = "packets"
)

type params struct {
	t      string // type
	dev    string
	ip     string
	limit  string
	action string
}

func parseCmd(cmd *cobra.Command, args []string) error {
	ps, err := validate(cmd)
	if err != nil {
		return err
	}
	switch ps.t {
	case limitTypeBandwidth:
		if err := makeBandwidthLimit(ps.ip, ps.dev, ps.action, ps.limit); err != nil {
			return cmd.Help()
		}
	case limitTypePackets:
		if err := makePacketsLimit(ps.ip, ps.action, ps.limit); err != nil {
			return cmd.Help()
		}
	default:
		return fmt.Errorf("Unknown type: %s", ps.t)
	}

	return nil
}

func makeBandwidthLimit(ip string, dev string, action string, limit string) error {
	b, err := filter.NewBandwidthLimit(dev, ip, limit)
	if err != nil {
		return err
	}
	switch action {
	case "set":
		if err := b.Apply(); err != nil {
			return err
		}
	case "unset":
		if err := b.Unset(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown action: %s", action)
	}
	return nil
}

func makePacketsLimit(ip string, action string, limit string) error {
	p, err := filter.NewPacketsLimit(ip, limit)
	if err != nil {
		return err
	}
	switch action {
	case "set":
		if err := p.Apply(); err != nil {
			return err
		}
	case "unset":
		if err := p.Unset(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown action: %s", action)
	}
	return nil
}

func validate(cmd *cobra.Command) (*params, error) {
	if len(cmd.Flags().Args()) <= 0 {
		return nil, cmd.Help()
	}
	dev, err := cmd.Flags().GetString("dev")
	if err != nil {
		return nil, cmd.Help()
	}
	t, err := cmd.Flags().GetString("type")
	if err != nil {
		return nil, cmd.Help()
	}
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return nil, cmd.Help()
	}
	limit, err := cmd.Flags().GetString("limit")
	if err != nil {
		return nil, cmd.Help()
	}
	action, err := cmd.Flags().GetString("action")
	if err != nil {
		return nil, cmd.Help()
	}
	if dev == "" || ip == "" || limit == "" || action == "" || t == "" {
		return nil, cmd.Help()
	}
	return &params{
		dev:    dev,
		ip:     ip,
		limit:  limit,
		action: action,
	}, nil
}
