package filter

import (
	"fmt"
	"net"

	"github.com/yuriykis/funny-filter/log"
)

type PacketsLimit struct {
	IP    string
	Limit string
}

func NewPacketsLimit(ip string, limit string) (*PacketsLimit, error) {
	if err := validatePacketsParams(ip, limit); err != nil {
		return nil, err
	}
	return &PacketsLimit{
		IP:    ip,
		Limit: limit,
	}, nil
}

func (p *PacketsLimit) Apply() error {

	log.WithFields(log.Fields{
		"ip":    p.IP,
		"limit": p.Limit,
	}).Info("Applying packets limit")

	if err := setPacketsLimit(p.IP, p.Limit); err != nil {
		return err
	}
	if err := setDropPacketsLimit(p.IP); err != nil {
		return err
	}
	return nil
}

func (p *PacketsLimit) Unset() error {

	log.WithFields(log.Fields{
		"ip":    p.IP,
		"limit": p.Limit,
	}).Info("Unsetting packets limit")

	if err := unsetPacketsLimit(p.IP, p.Limit); err != nil {
		return err
	}
	if err := unsetDropPacketsLimit(p.IP); err != nil {
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
		return fmt.Errorf("Limit is empty")
	}
	if s := limit; s != "" {
		for _, r := range s {
			if r < '0' || r > '9' {
				return fmt.Errorf("Limit should contain only digits")
			}
		}
	}
	return nil
}

func validateIP(ip string) error {
	if ip == "" {
		return fmt.Errorf("IP is empty")
	}
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("IP is invalid")
	}
	return nil
}
