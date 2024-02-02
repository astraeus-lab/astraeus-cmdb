package cache

import (
	"strings"
	"time"

	"github.com/astraeus-lab/astraeus-cmdb/src/common/util"

	"github.com/redis/go-redis/v9"
)

var defaultRedisConnect redis.UniversalClient

// InitRedisClient initialize redis connection based on config.
//
// Enabling cache should trigger an update to the Redis cache for
// each resource operation, updating the DB first and then Redis.
//
// All Key of Redis have a default expiration time of 1 hour
// and are maintained and updated by goroutine.
func InitRedisClient(endpoint []string, user, passwd, clientPrefix string, maxConn, maxIdelConn, coonMaxIdel int) error {
	defaultRedisConnect = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      endpoint,
		Username:   user,
		Password:   passwd,
		DB:         0,
		ClientName: strings.Join([]string{clientPrefix, util.RandStr(5)}, "-"),

		PoolFIFO:        true,
		PoolSize:        maxConn,
		MaxIdleConns:    maxIdelConn,
		MinIdleConns:    maxIdelConn / 2,
		ConnMaxIdleTime: time.Duration(coonMaxIdel) * time.Minute,
	})

	return nil
}

// GetCacheConnect get redis connection.
// Created by initializing configuration parameters.
//
// To ensure high availability, data read and write have timeout limits.
//
// Query errors may not necessarily be caused by the queries
// (e.g. no connections in the connection pool).
func GetCacheConnect() redis.UniversalClient {

	return defaultRedisConnect
}
