package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetIfbClassBuildCmd(t *testing.T) {
	rate := "100kbps"
	ceil := "200kbps"
	params := []string{"sudo tc class add dev ifb0 parent 1: classid 1:1 htb rate", rate, "ceil", ceil}
	cmd := build(params...)
	assert.Equal(t, "sudo tc class add dev ifb0 parent 1: classid 1:1 htb rate 100kbps ceil 200kbps", cmd)
}

func TestSetIngressFilterBuildCmd(t *testing.T) {
	dev := "enp0s5"
	params := []string{"sudo tc filter add dev", dev, "parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0"}
	cmd := build(params...)
	assert.Equal(t, "sudo tc filter add dev enp0s5 parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0", cmd)
}
