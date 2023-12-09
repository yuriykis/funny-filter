package linux

import (
	"github.com/vishvananda/netlink"
)

const (
	ifbName = "ifb0"
)

// we probably need to extend this to handle specific error messages as warnings
func CheckIfbModule() error {
	_, err := Run("modprobe ifb")
	return err
}

func CreateIfb(netlinker Netlinker) error {
	// _, err := Run("sudo ip link add name ifb0 type ifb")
	// if err != nil && err.Error() != "exit status 2" {
	// 	return err
	// }
	// return nil
	ifb := &netlink.Ifb{
		LinkAttrs: netlink.LinkAttrs{
			Name: ifbName,
		},
	}
	return netlinker.LinkAdd(ifb)
}

func SetUpIfb(netlinker Netlinker) error {
	// _, err := Run("sudo ip link set dev ifb0 up")
	ifb, err := netlinker.LinkByName(ifbName)
	if err != nil {
		return err
	}
	return netlinker.LinkSetUp(ifb)
}

func TearDownIfb(netlinker Netlinker) error {
	// _, err := Run("sudo ip link set dev ifb0 down")
	ifb, err := netlinker.LinkByName(ifbName)
	if err != nil {
		return err
	}
	return netlinker.LinkSetDown(ifb)
}

func DeleteIfb(netlinker Netlinker) error {
	// _, err := Run("sudo ip link delete ifb0 type ifb")
	ifb, err := netlinker.LinkByName(ifbName)
	if err != nil {
		return err
	}
	return netlinker.LinkDel(ifb)
}
