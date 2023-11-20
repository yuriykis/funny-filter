package filter

func checkIfbModule() error {
	_, err := run("sudo modprobe ifb")
	return err
}

func createIfb() error {
	_, err := run("sudo ip link add name ifb0 type ifb")
	return err
}

func setUpIfb() error {
	_, err := run("sudo ip link set dev ifb0 up")
	return err
}

func tearDownIfb() error {
	_, err := run("sudo ip link set dev ifb0 down")
	return err
}

func deleteIfb() error {
	_, err := run("sudo ip link delete ifb0 type ifb")
	return err
}
