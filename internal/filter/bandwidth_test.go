package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: add test to handle error from linux package
func TestNewBandwidthLimit(t *testing.T) {
	bw, err := NewBandwidthLimitLinux("enp0s5", "80.249.99.148", "100kbps")
	assert.Nil(t, err)
	assert.NotNil(t, bw)
	assert.Equal(t, "enp0s5", bw.Dev)
	assert.Equal(t, "80.249.99.148", bw.IP)
	assert.Equal(t, "100kbps", bw.Limit)
}

func TestNewBandwidthLimitWrongLimit(t *testing.T) {
	bw, err := NewBandwidthLimitLinux("enp0s5", "80.249.99.148", "100")
	assert.NotNil(t, err)
	assert.Nil(t, bw)
}

func TestNewBandwidthLimitEmptyLimit(t *testing.T) {
	bw, err := NewBandwidthLimitLinux("enp0s5", "80.249.99.148", "")
	assert.NotNil(t, err)
	assert.Nil(t, bw)
}

func TestNewBandwidthLimitWrongIP(t *testing.T) {
	bw, err := NewBandwidthLimitLinux("enp0s5", "80.249.99148", "100kbps")
	assert.NotNil(t, err)
	assert.Nil(t, bw)
}

func TestNewBandwidthLimitEmptyIP(t *testing.T) {
	bw, err := NewBandwidthLimitLinux("enp0s5", "", "100kbps")
	assert.NotNil(t, err)
	assert.Nil(t, bw)
}

func TestNewBandwidthLimitEmptyDev(t *testing.T) {
	bw, err := NewBandwidthLimitLinux("", "80.249.99148", "100kbps")
	assert.NotNil(t, err)
	assert.Nil(t, bw)
}

func TestNewBandwidthLimitWrongOS(t *testing.T) {
	bw, err := NewBandwidthLimit("windows", "enp0s5", "80.249.99148", "100kbps")
	assert.NotNil(t, err)
	assert.Nil(t, bw)
}
