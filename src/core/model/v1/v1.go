package v1

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// TODO: After enabling caching, all data updates should be cascaded to the cache.

type NewOptions struct {
	DBClient    *gorm.DB
	CacheClient redis.UniversalClient
}
