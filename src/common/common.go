package common

import (
	"fmt"

	"github.com/astraeus-lab/astraeus-cmdb/src/cmd"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/cache"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/config"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/constant"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/db"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/log"
)

// TODO: Hot update init

// InitCommonDepend init required common dependencies.
func InitCommonDepend() (*config.CMDBConfig, error) {
	cmdParam := cmd.ParseCMDParam()

	cmdbConfig, err := config.InitConfig(cmdParam.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("init config err: %v", err)
	}

	// common dependencies should avoid cross-referencing except for util
	if err = log.InitLogger(
		cmdbConfig.Config.Log.Path,
		cmdbConfig.Config.Log.Level,
		cmdbConfig.Config.Log.Stdout,
	); err != nil {
		return nil, fmt.Errorf("init logger err: %v", err.Error())
	}

	if err = db.InitDBConnectPool(
		cmdbConfig.Config.DB.Type,
		cmdbConfig.Config.DB.Host,
		cmdbConfig.Config.DB.User,
		cmdbConfig.Config.DB.Passwd,
		cmdbConfig.Config.DB.DBName,

		cmdbConfig.Config.DB.Option.MaxOpenConns,
		cmdbConfig.Config.DB.Option.MaxIdleConns,
		cmdbConfig.Config.DB.Option.ConnMaxIdleTimeMin,
	); err != nil {
		return nil, fmt.Errorf("init db connect err: %v", err.Error())
	}

	if cmdbConfig.Config.Redis.Enable {
		if err = cache.InitRedisClient(
			cmdbConfig.Config.Redis.Endpoint,
			cmdbConfig.Config.Redis.User,
			cmdbConfig.Config.Redis.Passwd,
			constant.AstraeusCMDBLower,

			cmdbConfig.Config.Redis.Option.MaxOpenConns,
			cmdbConfig.Config.Redis.Option.MaxIdleConns,
			cmdbConfig.Config.Redis.Option.ConnMaxIdleTimeMin,
		); err != nil {
			return nil, fmt.Errorf("init redis connect err: %v", err.Error())
		}
	}

	return cmdbConfig, nil
}
