package cmd

import "flag"

const (
	ConfigPathFlag string = "c"
)

// ParseCMDParam parsing CMD parameters.
func ParseCMDParam() *Param {
	rest := &Param{}

	flag.StringVar(
		&rest.ConfigPath,
		ConfigPathFlag,
		"./config.yaml",
		"Specify the config file path, can be yaml or json",
	)

	return rest
}
