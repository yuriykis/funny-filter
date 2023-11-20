package filter

import "fmt"

type BandwidthLimit struct {
	Dev   string
	IP    string
	Limit string
}

func NewBandwidthLimit(dev string, ip string, limit string) (*BandwidthLimit, error) {
	if err := validateBandwidthParams(dev, ip, limit); err != nil {
		return nil, err
	}
	return &BandwidthLimit{
		Dev:   dev,
		IP:    ip,
		Limit: limit,
	}, nil
}

func (b *BandwidthLimit) Apply() error {
	if err := checkIfbModule(); err != nil {
		return err
	}
	if err := createIfb(); err != nil {
		return err
	}
	if err := setUpIfb(); err != nil {
		return err
	}
	if err := setIngressQdisc(b.Dev); err != nil {
		return err
	}
	if err := setIngressFilter(b.Dev); err != nil {
		return err
	}
	if err := setIfbQdisc(); err != nil {
		return err
	}
	if err := setIfbClass(b.Limit, b.Limit); err != nil {
		return err
	}
	if err := setIfbFilter(b.IP); err != nil {
		return err
	}
	return nil
}

func (b *BandwidthLimit) Unset() error {
	if err := unsetIngressQdisc(b.Dev); err != nil {
		return err
	}
	if err := tearDownIfb(); err != nil {
		return err
	}
	if err := unsetIfbQdisc(); err != nil {
		return err
	}
	if err := deleteIfb(); err != nil {
		return err
	}
	return nil
}

func validateBandwidthParams(dev string, ip string, limit string) error {
	if dev == "" || ip == "" || limit == "" {
		return fmt.Errorf("dev, ip, limit must be specified")
	}
	return nil
}
