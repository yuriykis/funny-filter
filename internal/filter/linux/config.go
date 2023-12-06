package linux

const (
	configTypeBash    = "bash"
	configTypeNetlink = "netlink"
)

type LinuxConfig struct {
	Type string
}

func NewLinuxConfig(configType string) (*LinuxConfig, error) {
	if err := validateLinuxConfigParams(configType); err != nil {
		return nil, err
	}
	return &LinuxConfig{
		Type: configType,
	}, nil
}

func validateLinuxConfigParams(configType string) error {
	if configType != configTypeBash && configType != configTypeNetlink {
		return ErrInvalidLinuxConfigType()
	}
	return nil
}
