package config

const (
	DefaultConfigLogPath  = "/var/log/astraeus-cmdb/cmdb.log"
	DefaultConfigLogLevel = "info"

	DefaultConfigDBMaxOpenConns       = 50
	DefaultConfigDBMaxIdleConns       = 15
	DefaultConfigDBConnMaxIdleTimeMin = 5

	DefaultConfigRedisMaxOpenConns       = 30
	DefaultConfigRedisMaxIdleConns       = 10
	DefaultConfigRedisConnMaxIdleTimeMin = 3
)
