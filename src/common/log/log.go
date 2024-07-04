package log

import (
	"fmt"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/config"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/log/xlog"
	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"
)

var (
	defaultLogger *xlog.XLog

	Debug func(msg string, arg ...any)
	Info  func(msg string, arg ...any)
	Warn  func(msg string, arg ...any)
	Error func(msg string, arg ...any)
)

// InitLogger init different level of logger based on config.
// Parame level determines the log output level.
func InitLogger(c *config.Log) (err error) {
	if defaultLogger, err = InitCustomLogger(c.Path, c.Level, c.Stdout); err != nil {
		return err
	}
	initPackageLogger()

	return nil
}

func InitCustomLogger(path string, level string, isStdout bool) (*xlog.XLog, error) {
	if util.IsFileExistOrCreate(path) != nil || !util.IsFilePermission(path) {
		return nil, fmt.Errorf("log file(%s) may have creation or permission err", path)
	}

	var err error
	logger, err := xlog.NewXLog(path, level, isStdout)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func initPackageLogger() {
	Debug = defaultLogger.Debug
	Info = defaultLogger.Info
	Warn = defaultLogger.Warn
	Error = defaultLogger.Error
}
