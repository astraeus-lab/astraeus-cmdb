package xlog

type LogLevel string

const (
	DebugLogLevel LogLevel = "DEBUG"
	InfoLogLevel  LogLevel = "INFO"
	WarnLogLevel  LogLevel = "WARN"
	ErrorLogLevel LogLevel = "ERROR"
)
