#!/bin/bash
# this can be used to calculate the number of packets received from a specific IP address
sudo tcpdump -i enp0s5 host 80.249.99.148 -l | while read line; do
    echo "$line" | awk '{print strftime("%Y-%m-%d %H:%M:%S")}' >>  tcpdump_output.txt
done

sort tcpdump_output.txt | uniq -c

> tcpdump_output.txt