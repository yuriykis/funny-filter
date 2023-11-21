package filter

import (
	"net"

	"github.com/yuriykis/funny-filter/internal/filter/linux"
	"github.com/yuriykis/funny-filter/log"
)

const (
	PacketsLimitTypeLinux = "linux"
)

// PacketsLimit is generic interface for packets limit
// We can implement this interface for other OSes
type PacketsLimit interface {
	Set() error
	Unset() error
}

func NewPacketsLimit(osType string, ip string, limit string) (PacketsLimit, error) {
	switch osType {
	case PacketsLimitTypeLinux:
		return NewPacketsLimitLinux(ip, limit)
	default:
		return nil, ErrWrongOS(osType)
	}
}

type PacketsLimitLinux struct {
	IP    string
	Limit string
}

func NewPacketsLimitLinux(ip string, limit string) (*PacketsLimitLinux, error) {
	if err := validatePacketsParams(ip, limit); err != nil {
		return nil, err
	}
	return &PacketsLimitLinux{
		IP:    ip,
		Limit: limit,
	}, nil
}

func (p *PacketsLimitLinux) Set() error {

	log.WithFields(log.Fields{
		"ip":    p.IP,
		"limit": p.Limit,
	}).Info("Applying packets limit")

	if err := linux.SetPacketsLimit(p.IP, p.Limit); err != nil {
		return err
	}
	if err := linux.SetDropPacketsLimit(p.IP); err != nil {
		return err
	}
	return nil
}

func (p *PacketsLimitLinux) Unset() error {

	log.WithFields(log.Fields{
		"ip":    p.IP,
		"limit": p.Limit,
	}).Info("Unsetting packets limit")

	if err := linux.UnsetPacketsLimit(p.IP, p.Limit); err != nil {
		return err
	}
	if err := linux.UnsetDropPacketsLimit(p.IP); err != nil {
		return err
	}
	return nil
}

func validatePacketsParams(ip string, limit string) error {
	if err := validatePacketsLimit(limit); err != nil {
		return err
	}
	if err := validateIP(ip); err != nil {
		return err
	}
	return nil
}

func validatePacketsLimit(limit string) error {
	if limit == "" {
		return ErrInvalidPacketsLimit("Limit is empty", limit)
	}
	// check if limit contains only digits
	if s := limit; s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return ErrInvalidBandwidthLimit("Limit should contain only digits", limit)
			}
		}
	}
	return nil
}
func validateIP(ip string) error {
	if ip == "" {
		return ErrInvalidIP("IP is empty", ip)
	}
	if net.ParseIP(ip) == nil {
		return ErrInvalidIP("IP is invalid", ip)
	}
	return nil
}
