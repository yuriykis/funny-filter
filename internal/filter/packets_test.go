package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: add test to handle error from linux package
func TestNewPacketsLimit(t *testing.T) {
	pl, err := NewPacketsLimitLinux("80.249.99.148", "10")
	assert.Nil(t, err)
	assert.NotNil(t, pl)
	assert.Equal(t, "80.249.99.148", pl.IP)
	assert.Equal(t, "10", pl.Limit)
}

func TestNewPacketsLimitWrongLimit(t *testing.T) {
	pl, err := NewPacketsLimitLinux("80.249.99.148", "10kbps")
	assert.NotNil(t, err)
	assert.Nil(t, pl)
}

func TestNewPacketsLimitEmptyLimit(t *testing.T) {
	pl, err := NewPacketsLimitLinux("80.249.99.148", "")
	assert.NotNil(t, err)
	assert.Nil(t, pl)
}

func TestNewPacketsLimitWrongIP(t *testing.T) {
	pl, err := NewPacketsLimitLinux("80.249.99148", "10")
	assert.NotNil(t, err)
	assert.Nil(t, pl)
}

func TestNewPacketsLimitEmptyIP(t *testing.T) {
	pl, err := NewPacketsLimitLinux("", "10")
	assert.NotNil(t, err)
	assert.Nil(t, pl)
}
func TestNewPacketsLimitWrongOS(t *testing.T) {
	pl, err := NewPacketsLimit("windows", "80.249.99148", "10")
	assert.NotNil(t, err)
	assert.Nil(t, pl)
}
