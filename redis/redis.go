package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"utils/global"
)

func InitRedis() {
	config := global.Config.Redis
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       0,
	})
	_, err := global.Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接")
}
