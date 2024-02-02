package config

import (
	"fmt"
	"os"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
)

// InitConfig init global config.
func InitConfig(configPath string) (*CMDBConfig, error) {
	if !util.IsFileExist(configPath) || !util.IsFilePermission(configPath) {
		return nil, fmt.Errorf("%s path config file does not exist or has insufficient permission", configPath)
	}

	configContent, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	return NewCMDBConfig(configContent)
}
