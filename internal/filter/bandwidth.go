package filter

import (
	"github.com/yuriykis/funny-filter/internal/filter/linux"
	"github.com/yuriykis/funny-filter/log"
)

const (
	BandwidthLimitTypeLinux = "linux"
)

type BandwidthLimit interface {
	Set() error
	Unset() error
}

func NewBandwidthLimit(osType string, dev string, ip string, limit string) (BandwidthLimit, error) {
	switch osType {
	case BandwidthLimitTypeLinux:
		return NewBandwidthLimitLinux(dev, ip, limit)
	default:
		return nil, ErrWrongOS(osType)
	}
}

type BandwidthLimitLinux struct {
	Dev   string
	IP    string
	Limit string
}

func NewBandwidthLimitLinux(dev string, ip string, limit string) (*BandwidthLimitLinux, error) {
	if err := validateBandwidthParams(dev, ip, limit); err != nil {
		return nil, err
	}
	return &BandwidthLimitLinux{
		Dev:   dev,
		IP:    ip,
		Limit: limit,
	}, nil
}

func (b *BandwidthLimitLinux) Set() error {

	log.WithFields(log.Fields{
		"dev":   b.Dev,
		"ip":    b.IP,
		"limit": b.Limit,
	}).Info("Applying bandwidth limit")

	// TODO: handle errors, some errors can be ignored and treated as warnings
	// we can handle specific error messages and decide if we should return error or not
	// there is no error handling linux bash commands at the moment
	if err := linux.CheckIfbModule(); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.CreateIfb(); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.SetUpIfb(); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.SetIngressQdisc(b.Dev); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.SetIngressFilter(b.Dev); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.SetIfbQdisc(); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.SetIfbClass(b.Limit, b.Limit); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.SetIfbFilter(b.IP); err != nil {
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

func (b *BandwidthLimitLinux) Unset() error {

	log.WithFields(log.Fields{
		"dev":   b.Dev,
		"ip":    b.IP,
		"limit": b.Limit,
	}).Info("Unsetting bandwidth limit")

	if err := linux.UnsetIngressQdisc(b.Dev); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.TearDownIfb(); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.UnsetIfbQdisc(); err != nil {
		log.Error(err)
		// return err
	}
	if err := linux.DeleteIfb(); err != nil {
		log.Error(err)
		// return err
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
		return ErrInvalidDev("dev is empty", dev)
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
		return ErrInvalidBandwidthLimit("limit is too short", limit)
	}
	if limit[len(limit)-4:] != "kbps" && limit[len(limit)-4:] != "mbps" && limit[len(limit)-4:] != "gbps" {
		return ErrInvalidBandwidthLimit("limit is not in kbps, mbps or gbps format", limit)
	}
	return nil
}
