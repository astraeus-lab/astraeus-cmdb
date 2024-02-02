package xlog

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	DebugLevel string = "debug"
	InfoLevel  string = "info"
	WarnLevel  string = "warn"
	ErrorLevel string = "error"
)

// XLog multiple level of logger.
type XLog struct {
	logger *slog.Logger
}

// NewXLog create a logger based on config.
// Based on log/slog, level param affects the output level of the log.
func NewXLog(dest string, level string, isStdout bool) (*XLog, error) {
	file, err := os.OpenFile(dest, os.O_APPEND, os.ModeAppend)
	if err != nil {
		return nil, fmt.Errorf("open %s log file err: %v", dest, err)
	}
	writer := io.MultiWriter(file)
	if isStdout {
		writer = io.MultiWriter(file, os.Stdout)
	}

	return &XLog{
		logger: slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: converLogLevel(level)})),
	}, nil
}

func (x *XLog) Debug(mgs string, arg ...any) {
	x.logger.Debug(mgs, arg)
}

func (x *XLog) Info(mgs string, arg ...any) {
	x.logger.Info(mgs, arg)
}

func (x *XLog) Warn(mgs string, arg ...any) {
	x.logger.Warn(mgs, arg)
}

func (x *XLog) Error(mgs string, arg ...any) {
	x.logger.Error(mgs, arg)
}

// converLogLevel convert custom log levels to log/slog format.
// If the log level is not recognized, it defaults to slog.LevelInfo.
func converLogLevel(level string) slog.Level {
	switch level {
	case DebugLevel:
		return slog.LevelDebug
	case InfoLevel:
		return slog.LevelInfo
	case WarnLevel:
		return slog.LevelWarn
	case ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
