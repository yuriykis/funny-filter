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

func CreateIfb() error {
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
	if err := netlink.LinkAdd(ifb); err != nil {
		return err
	}
	return nil
}

func SetUpIfb() error {
	// _, err := Run("sudo ip link set dev ifb0 up")
	ifb, err := netlink.LinkByName(ifbName)
	if err != nil {
		return err
	}
	if err := netlink.LinkSetUp(ifb); err != nil {
		return err
	}
	return nil
}

func TearDownIfb() error {
	// _, err := Run("sudo ip link set dev ifb0 down")
	ifb, err := netlink.LinkByName(ifbName)
	if err != nil {
		return err
	}
	if err := netlink.LinkSetDown(ifb); err != nil {
		return err
	}
	return nil
}

func DeleteIfb() error {
	// _, err := Run("sudo ip link delete ifb0 type ifb")
	ifb, err := netlink.LinkByName(ifbName)
	if err != nil {
		return err
	}
	if err := netlink.LinkDel(ifb); err != nil {
		return err
	}
	return err
}
