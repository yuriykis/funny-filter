package filter

func setIngressQdisc(dev string) error {
	_, err := run(build("sudo tc qdisc add dev", dev, "handle ffff: ingress"))
	return err
}

func setIngressFilter(dev string) error {
	_, err := run(build("sudo tc filter add dev", dev, "parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0"))
	return err
}

func setIfbQdisc() error {
	_, err := run("sudo tc qdisc add dev ifb0 root handle 1: htb default 11")
	return err
}

func setIfbClass(rate string, ceil string) error {
	_, err := run(build("sudo tc class add dev ifb0 parent 1: classid 1:1 htb rate", rate, "ceil", ceil))
	return err
}

func setIfbFilter(ip string) error {
	_, err := run(build("sudo tc filter add dev ifb0 protocol ip parent 1: prio 1 u32 match ip src", ip, "flowid 1:1"))
	return err
}

func unsetIngressQdisc(dev string) error {
	_, err := run(build("sudo tc qdisc del dev", dev, "ingress"))
	return err
}

func unsetIfbQdisc() error {
	_, err := run("sudo tc qdisc del dev ifb0 root")
	return err
}
