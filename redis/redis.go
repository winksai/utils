package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	rdb  *redis.Client
	lock sync.Mutex
)

// InitRedis 初始化 Redis 客户端
// 如果之前已连接，将关闭旧连接并重新初始化
func InitRedis(host, pwd string, db int) *redis.Client {
	lock.Lock()
	defer lock.Unlock()

	if rdb != nil {
		_ = rdb.Close()
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     pwd,
		DB:           db,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Redis 已连接:", host)
	return rdb
}
