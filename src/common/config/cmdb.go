package config

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// CMDBConfig represents the config.
type CMDBConfig struct {
	Config          CoreConfig      `json:"config" yaml:"config"`
	APIServerConfig APIServerConfig `json:"apiServer,omitempty" yaml:"apiServer,omitempty"`
}

func NewCMDBConfig(source []byte) (*CMDBConfig, error) {
	conf, err := unmarshalConfig(source)
	if err != nil {
		return nil, err
	}

	conf.Config.completeConfig()

	return conf, nil
}

// unmarshalConfig unmarshal configuration content to CMDBConfig.
// It will try JSON unmarshal first, then try yaml unmarshal,
// can also be used to verify configuration content.
func unmarshalConfig(source []byte) (*CMDBConfig, error) {
	conf := &CMDBConfig{}

	if json.Valid(source) {
		if err := json.Unmarshal(source, conf); err != nil {
			return nil, err
		}

		return conf, nil
	}

	if err := yaml.Unmarshal(source, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
