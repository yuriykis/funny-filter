package filter

import "fmt"

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
	if err := setPacketsLimit(p.IP, p.Limit); err != nil {
		return err
	}
	if err := setDropPacketsLimit(p.IP); err != nil {
		return err
	}
	return nil
}

func (p *PacketsLimit) Unset() error {
	if err := unsetPacketsLimit(p.IP, p.Limit); err != nil {
		return err
	}
	if err := unsetDropPacketsLimit(p.IP); err != nil {
		return err
	}
	return nil
}

func validatePacketsParams(ip string, limit string) error {
	if ip == "" {
		return fmt.Errorf("IP is empty")
	}
	if limit == "" {
		return fmt.Errorf("Limit is empty")
	}
	return nil
}
