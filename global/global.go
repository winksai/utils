package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"utils/config"
)

var (
	Config *config.T
	DB     *gorm.DB
	Rdb    *redis.Client
)
