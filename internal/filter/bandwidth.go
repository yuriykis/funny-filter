package filter

import (
	"fmt"

	"github.com/yuriykis/funny-filter/log"
)

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

	log.WithFields(log.Fields{
		"dev":   b.Dev,
		"ip":    b.IP,
		"limit": b.Limit,
	}).Info("Applying bandwidth limit")

	if err := checkIfbModule(); err != nil {
		return err
	}
	if err := createIfb(); err != nil {
		log.Error(err)
		// return err
	}
	if err := setUpIfb(); err != nil {
		log.Error(err)
		// return err
	}
	if err := setIngressQdisc(b.Dev); err != nil {
		log.Error(err)
		// return err
	}
	if err := setIngressFilter(b.Dev); err != nil {
		log.Error(err)
		// return err
	}
	if err := setIfbQdisc(); err != nil {
		log.Error(err)
		// return err
	}
	if err := setIfbClass(b.Limit, b.Limit); err != nil {
		log.Error(err)
		// return err
	}
	if err := setIfbFilter(b.IP); err != nil {
		// return err
		log.Error(err)
	}

	log.WithFields(log.Fields{
		"dev":   b.Dev,
		"ip":    b.IP,
		"limit": b.Limit,
	}).Info("Bandwidth limit applied")

	return nil
}

func (b *BandwidthLimit) Unset() error {

	log.WithFields(log.Fields{
		"dev":   b.Dev,
		"ip":    b.IP,
		"limit": b.Limit,
	}).Info("Unsetting bandwidth limit")

	if err := unsetIngressQdisc(b.Dev); err != nil {
		log.Error(err)
	}
	if err := tearDownIfb(); err != nil {
		log.Error(err)
	}
	if err := unsetIfbQdisc(); err != nil {
		log.Error(err)
	}
	if err := deleteIfb(); err != nil {
		log.Error(err)
	}

	log.WithFields(log.Fields{
		"dev":   b.Dev,
		"ip":    b.IP,
		"limit": b.Limit,
	}).Info("Bandwidth limit unset")

	return nil
}

func validateBandwidthParams(dev string, ip string, limit string) error {
	if dev == "" {
		return fmt.Errorf("dev is empty")
	}
	if err := validateBandwidthLimit(limit); err != nil {
		return err
	}
	if err := validateIP(ip); err != nil {
		return err
	}
	return nil
}

func validateBandwidthLimit(limit string) error {
	if len(limit) < 4 {
		return fmt.Errorf("limit is too short")
	}
	if limit[len(limit)-4:] != "kbps" && limit[len(limit)-4:] != "mbps" && limit[len(limit)-4:] != "gbps" {
		return fmt.Errorf("limit must be in kbps, mbps or gbps")
	}
	return nil
}
