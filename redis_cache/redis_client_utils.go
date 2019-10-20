package redis_cache

import (
    "fmt"
    "github.com/go-redis/redis"
    "os"
    "time"
)

type redisClientUtils struct {
}

var RedisClientUtils *redisClientUtils

func (rcu *redisClientUtils) InitRedisClient(_address string, _password string, _dbNum int) *redis.Client {
    redisClient := redis.NewClient(&redis.Options{
        Addr:        _address,
        Password:    _password, // no password set
        DB:          _dbNum,    // use default DB
        DialTimeout: time.Second * 2,
    })
    return redisClient
}

func (rcu *redisClientUtils) InitRedisClientByConfig(_redisConfig RedisConfigFormat) *redis.Client {
    redisClient := redis.NewClient(&redis.Options{
        Addr:        _redisConfig.Address,
        Password:    _redisConfig.Password, // no password set
        DB:          _redisConfig.DBNum,    // use default DB
        DialTimeout: time.Second * 2,
    })
    return redisClient
}

func (rcu *redisClientUtils) CheckRedisClient(_redisClient *redis.Client) bool {
    _, err := _redisClient.Ping().Result()
    if err != nil {
        _, _ = fmt.Fprintln(os.Stderr, time.Now().Format(time.RFC3339Nano), "[Error]", "redis connect error:", err)
        return false
    }
    return true
}
