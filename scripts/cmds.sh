#!/bin/bash

# set the bandwidth for ingress traffic

# sudo modprobe ifb
# sudo ip link add name ifb0 type ifb
# sudo ip link set dev ifb0 up

# sudo tc qdisc add dev enp0s5 handle ffff: ingress
# sudo tc filter add dev enp0s5 parent ffff: protocol ip u32 match u32 0 0 action mirred egress redirect dev ifb0

# sudo tc qdisc add dev ifb0 root handle 1: htb default 11
# sudo tc class add dev ifb0 parent 1: classid 1:1 htb rate 100kbps ceil 100kbps

# sudo tc filter add dev ifb0 protocol ip parent 1: prio 1 u32 match ip src 80.249.99.148 flowid 1:1


# remove the bandwidth limit
# sudo tc qdisc del dev enp0s5 ingress
# sudo ip link set dev ifb0 down
# sudo tc qdisc del dev ifb0 root
# sudo ip link delete ifb0 type ifb

# limit packets per second
# sudo iptables -A INPUT -p tcp -s 80.249.99.148 -m hashlimit --hashlimit 10/sec --hashlimit-mode srcip --hashlimit-name limit -j ACCEPT
# sudo iptables -A INPUT -p tcp -s 80.249.99.148 -j DROP

# remove the limit
# sudo iptables -D INPUT -p tcp -s 80.249.99.148 -m hashlimit --hashlimit 10/sec --hashlimit-mode srcip --hashlimit-name limit -j ACCEPT
# sudo iptables -D INPUT -p tcp -s 80.249.99.148 -j DROP
