package linux

import (
	"github.com/vishvananda/netlink"
)

// we probably need to extend this to handle specific error messages as warnings
func CheckIfbModule() error {
	_, err := Run("sudo modprobe ifb")
	return err
}

func CreateIfb() error {
	// _, err := Run("sudo ip link add name ifb0 type ifb")
	// if err != nil && err.Error() != "exit status 2" {
	// 	return err
	// }
	// return nil
	ifbName := "ifb0"
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
	_, err := Run("sudo ip link set dev ifb0 up")
	return err
}

func TearDownIfb() error {
	_, err := Run("sudo ip link set dev ifb0 down")
	return err
}

func DeleteIfb() error {
	_, err := Run("sudo ip link delete ifb0 type ifb")
	return err
}
