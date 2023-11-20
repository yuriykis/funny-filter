package filter

func setPacketsLimit(ip string, limit string) error {
	_, err := run(build("sudo iptables -A INPUT -p tcp -s", ip, "-m hashlimit --hashlimit", limit, "--hashlimit-mode srcip --hashlimit-name limit -j ACCEPT"))
	return err
}

func setDropPacketsLimit(ip string) error {
	_, err := run(build("sudo iptables -A INPUT -p tcp -s", ip, "-j DROP"))
	return err
}
func unsetPacketsLimit(ip string, limit string) error {
	_, err := run(build("sudo iptables -D INPUT -p tcp -s", ip, "-m hashlimit --hashlimit", limit, "--hashlimit-mode srcip --hashlimit-name limit -j ACCEPT"))
	return err
}

func unsetDropPacketsLimit(ip string) error {
	_, err := run(build("sudo iptables -D INPUT -p tcp -s", ip, "-j DROP"))
	return err
}
