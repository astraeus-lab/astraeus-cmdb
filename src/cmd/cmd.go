package cmd

import "flag"

const (
	ConfigPathFlag string = "c"
)

// ParseCMDParam parsing CMD parameters.
func ParseCMDParam() *Param {
	rest := &Param{}

	rest.ConfigPath = *flag.String(
		ConfigPathFlag,
		"./config.yaml",
		"specify the config file path, default: ./config.yaml",
	)

	return rest
}
