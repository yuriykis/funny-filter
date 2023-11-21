package linux

// we probably need to extend this to handle specific error messages as warnings
func SetIngressQdisc(dev string) error {
	_, err := Run(Build("sudo tc qdisc add dev", dev, "handle ffff: ingress"))
	return err
}

func SetIngressFilter(dev string) error {
	_, err := Run(Build("sudo tc filter add dev", dev, "parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0"))
	return err
}

func SetIfbQdisc() error {
	_, err := Run("sudo tc qdisc add dev ifb0 root handle 1: htb default 11")
	return err
}

func SetIfbClass(rate string, ceil string) error {
	_, err := Run(Build("sudo tc class add dev ifb0 parent 1: classid 1:1 htb rate", rate, "ceil", ceil))
	return err
}

func SetIfbFilter(ip string) error {
	_, err := Run(Build("sudo tc filter add dev ifb0 protocol ip parent 1: prio 1 u32 match ip src", ip, "flowid 1:1"))
	return err
}

func UnsetIngressQdisc(dev string) error {
	_, err := Run(Build("sudo tc qdisc del dev", dev, "ingress"))
	return err
}

func UnsetIfbQdisc() error {
	_, err := Run("sudo tc qdisc del dev ifb0 root")
	return err
}
