package log

import (
	"fmt"

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
func InitLogger(path string, level string, isStdout bool) error {
	if util.IsFileExistOrCreate(path) != nil || !util.IsFilePermission(path) {
		return fmt.Errorf("%s path log file may have error in exist, create or permission", path)
	}

	var err error
	defaultLogger, err = xlog.NewXLog(path, level, isStdout)
	if err != nil {
		return err
	}

	return nil
}
