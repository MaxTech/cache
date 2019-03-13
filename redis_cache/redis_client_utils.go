package redis_cache

import (
    "github.com/go-redis/redis"
    "log"
    "time"
)

type redisClientUtils struct {
}

var RedisClientUtils *redisClientUtils

func (ru *redisClientUtils) InitRedisClient(address string, password string, dbNum int) *redis.Client {
    redisClient := redis.NewClient(&redis.Options{
        Addr:        address,
        Password:    password, // no password set
        DB:          dbNum,    // use default DB
        DialTimeout: time.Second * 2,
    })
    return redisClient
}

func (ru *redisClientUtils) CheckRedisClient(redisClient *redis.Client) bool {
    _, err := redisClient.Ping().Result()
    if err != nil {
        log.Println("redis connect error:", err)
        return false
    }
    return true
}
