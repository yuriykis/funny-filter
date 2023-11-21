package linux

// we probably need to extend this to handle specific error messages as warnings
func SetPacketsLimit(ip string, limit string) error {
	_, err := Run(Build("sudo iptables -A INPUT -p tcp -s", ip, "-m hashlimit --hashlimit", limit, "--hashlimit-mode srcip --hashlimit-name limit -j ACCEPT"))
	return err
}

func SetDropPacketsLimit(ip string) error {
	_, err := Run(Build("sudo iptables -A INPUT -p tcp -s", ip, "-j DROP"))
	return err
}
func UnsetPacketsLimit(ip string, limit string) error {
	_, err := Run(Build("sudo iptables -D INPUT -p tcp -s", ip, "-m hashlimit --hashlimit", limit, "--hashlimit-mode srcip --hashlimit-name limit -j ACCEPT"))
	return err
}

func UnsetDropPacketsLimit(ip string) error {
	_, err := Run(Build("sudo iptables -D INPUT -p tcp -s", ip, "-j DROP"))
	return err
}
