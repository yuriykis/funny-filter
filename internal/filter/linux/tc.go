package linux

import (
	"encoding/binary"
	"net"
	"strconv"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

// we probably need to extend this to handle specific error messages as warnings
func SetIngressQdisc(dev string) error {
	// _, err := Run(Build("sudo tc qdisc add dev", dev, "handle ffff: ingress"))
	// return err

	link, err := netlink.LinkByName(dev)
	if err != nil {
		return err
	}

	qdisc := &netlink.Ingress{
		QdiscAttrs: netlink.QdiscAttrs{
			LinkIndex: link.Attrs().Index,
			Handle:    netlink.MakeHandle(0xffff, 0),
			Parent:    netlink.HANDLE_INGRESS,
		},
	}

	return netlink.QdiscAdd(qdisc)
}

func SetIngressFilter(dev string) error {
	// _, err := Run(Build("sudo tc filter add dev", dev, "parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0"))
	// return err

	link, err := netlink.LinkByName(dev)
	if err != nil {
		return err
	}

	ifbLink, err := netlink.LinkByName("ifb0")
	if err != nil {
		return err
	}

	filter := &netlink.U32{
		FilterAttrs: netlink.FilterAttrs{
			LinkIndex: link.Attrs().Index,
			Parent:    netlink.MakeHandle(0xffff, 0),
			Protocol:  unix.ETH_P_IP,
			Priority:  1,
		},
		Actions: []netlink.Action{
			&netlink.MirredAction{
				ActionAttrs:  netlink.ActionAttrs{},
				Ifindex:      ifbLink.Attrs().Index,
				MirredAction: netlink.TCA_EGRESS_REDIR,
			},
		},
	}
	return netlink.FilterAdd(filter)
}

func SetIfbQdisc() error {
	// _, err := Run("sudo tc qdisc add dev ifb0 root handle 1: htb default 11")
	// return err

	link, err := netlink.LinkByName("ifb0")
	if err != nil {
		return err
	}

	qdisc := &netlink.Htb{
		QdiscAttrs: netlink.QdiscAttrs{
			LinkIndex: link.Attrs().Index,
			Handle:    netlink.MakeHandle(1, 0),
			Parent:    netlink.HANDLE_ROOT,
		},
		Defcls: 11,
	}

	return netlink.QdiscAdd(qdisc)
}

func SetIfbClass(rate string, ceil string) error {
	// _, err := Run(Build("sudo tc class add dev ifb0 parent 1: classid 1:1 htb rate", rate, "ceil", ceil))
	// return err

	link, err := netlink.LinkByName("ifb0")
	if err != nil {
		return err
	}

	rateUint64, err := strconv.ParseUint(rate, 10, 32)
	if err != nil {
		return err
	}
	ceilUint64, err := strconv.ParseUint(ceil, 10, 32)
	if err != nil {
		return err
	}
	class := &netlink.HtbClass{
		ClassAttrs: netlink.ClassAttrs{
			LinkIndex: link.Attrs().Index,
			Handle:    netlink.MakeHandle(1, 1),
			Parent:    netlink.MakeHandle(1, 0),
		},
		Rate:    rateUint64,
		Ceil:    ceilUint64,
		Quantum: 1500,
	}

	return netlink.ClassAdd(class)
}

func SetIfbFilter(ip string) error {
	// _, err := Run(Build("sudo tc filter add dev ifb0 protocol ip parent 1: prio 1 u32 match ip src", ip, "flowid 1:1"))
	// return err

	link, err := netlink.LinkByName("ifb0")
	if err != nil {
		return err
	}

	srcIP := net.ParseIP(ip)

	filter := &netlink.U32{
		FilterAttrs: netlink.FilterAttrs{
			LinkIndex: link.Attrs().Index,
			Parent:    netlink.MakeHandle(1, 0),
			Protocol:  unix.ETH_P_IP,
			Priority:  1,
		},
		Sel: &netlink.TcU32Sel{
			Keys: []netlink.TcU32Key{
				{Mask: 0xffffffff, Val: binary.BigEndian.Uint32(srcIP), Off: 12},
			},
			Flags: netlink.TC_U32_TERMINAL,
		},
		ClassId: netlink.MakeHandle(1, 1),
	}

	return netlink.FilterAdd(filter)
}

func UnsetIngressQdisc(dev string) error {
	// _, err := Run(Build("sudo tc qdisc del dev", dev, "ingress"))
	// return err

	link, err := netlink.LinkByName(dev)
	if err != nil {
		return err
	}

	qdisc := &netlink.Ingress{
		QdiscAttrs: netlink.QdiscAttrs{
			LinkIndex: link.Attrs().Index,
			Handle:    netlink.MakeHandle(0xffff, 0),
			Parent:    netlink.HANDLE_INGRESS,
		},
	}

	return netlink.QdiscDel(qdisc)
}

func UnsetIfbQdisc() error {
	// _, err := Run("sudo tc qdisc del dev ifb0 root")
	// return err

	link, err := netlink.LinkByName("ifb0")
	if err != nil {
		return err
	}

	qdisc := &netlink.Htb{
		QdiscAttrs: netlink.QdiscAttrs{
			LinkIndex: link.Attrs().Index,
			Handle:    netlink.MakeHandle(1, 0),
			Parent:    netlink.HANDLE_ROOT,
		},
	}

	return netlink.QdiscDel(qdisc)
}
