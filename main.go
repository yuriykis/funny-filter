package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/yuriykis/funny-filter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func setBandwidthLimit() {
	executeCommand("sudo modprobe ifb")
	executeCommand("sudo ip link add name ifb0 type ifb")
	executeCommand("sudo ip link set dev ifb0 up")
	executeCommand("sudo tc qdisc add dev enp0s3 handle ffff: ingress")
	executeCommand("sudo tc filter add dev enp0s3 parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0")
	executeCommand("sudo tc qdisc add dev ifb0 root handle 1: htb default 11")
	executeCommand("sudo tc class add dev ifb0 parent 1: classid 1:1 htb rate 100kbps ceil 100kbps")
	executeCommand("sudo tc filter add dev ifb0 protocol ip parent 1: prio 1 u32 match ip src 80.249.99.148 flowid 1:1")
}

func unsetBandwidthLimit() {
	executeCommand("sudo tc qdisc del dev enp0s3 ingress")
	executeCommand("sudo ip link set dev ifb0 down")
	executeCommand("sudo tc qdisc del dev ifb0 root")
	executeCommand("sudo ip link delete ifb0 type ifb")
}
func setPacketsLimit() {
	executeCommand("sudo iptables -A INPUT -p tcp -s 80.249.99.148 -m hashlimit --hashlimit 10/sec --hashlimit-mode srcip --hashlimit-name limit -j ACCEPT")
	executeCommand("sudo iptables -A INPUT -p tcp -s 80.249.99.148 -j DROP")
}

func unsetPacketsLimit() {
	executeCommand("sudo iptables -D INPUT -p tcp -s 80.249.99.148 -m hashlimit --hashlimit 10/sec --hashlimit-mode srcip --hashlimit-name limit -j ACCEPT")
	executeCommand("sudo iptables -D INPUT -p tcp -s 80.249.99.148 -j DROP")
}

func executeCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	} else {
		fmt.Println(string(output))
	}
}
