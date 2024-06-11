package log

import (
	"fmt"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/config"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/log/xlog"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
)

var (
	defaultLogger *xlog.XLog

	Debug = defaultLogger.Debug
	Info  = defaultLogger.Info
	Warn  = defaultLogger.Warn
	Error = defaultLogger.Error
)

// InitLogger init different level of logger based on config.
// Parame level determines the log output level.
func InitLogger(c *config.Log) error {
	// path string, level string, isStdout bool
	if util.IsFileExistOrCreate(c.Path) != nil || !util.IsFilePermission(c.Path) {
		return fmt.Errorf("%s path log file may have error in exist, create or permission", c.Path)
	}

	var err error
	defaultLogger, err = xlog.NewXLog(c.Path, c.Level, c.Stdout)
	if err != nil {
		return err
	}

	return nil
}
