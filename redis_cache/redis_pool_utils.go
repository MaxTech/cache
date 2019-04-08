package redis_cache

import (
    "github.com/gomodule/redigo/redis"
    "log"
    "time"
)

type redisPoolUtils struct {
}

var RedisPoolUtils *redisPoolUtils

func (rpu *redisPoolUtils) InitRedisPool(address string, password string, dbNum int) *redis.Pool {
    return &redis.Pool{
        MaxIdle:     500,
        MaxActive:   50,
        IdleTimeout: 60 * time.Second,
        Wait:        true,
        Dial: func() (redis.Conn, error) {
            con, err := redis.Dial("tcp", address,
                redis.DialPassword(password),
                redis.DialDatabase(dbNum),
                redis.DialConnectTimeout(2*time.Second),
                redis.DialReadTimeout(2*time.Second),
                redis.DialWriteTimeout(2*time.Second))
            if err != nil {
                return nil, err
            }
            return con, nil
        },
    }
}

func (rpu *redisPoolUtils) InitRedisPoolByConfig(redisConfig RedisConfigFormat) *redis.Pool {
    return &redis.Pool{
        MaxIdle:     500,
        MaxActive:   50,
        IdleTimeout: 60 * time.Second,
        Wait:        true,
        Dial: func() (redis.Conn, error) {
            con, err := redis.Dial("tcp", redisConfig.Address,
                redis.DialPassword(redisConfig.Password),
                redis.DialDatabase(redisConfig.DBNum),
                redis.DialConnectTimeout(2*time.Second),
                redis.DialReadTimeout(2*time.Second),
                redis.DialWriteTimeout(2*time.Second))
            if err != nil {
                return nil, err
            }
            return con, nil
        },
    }
}

func (rpu *redisPoolUtils) CheckRedisPool(redisPool *redis.Pool) bool {
    conn := redisPool.Get()
    defer conn.Close()
    if conn.Err() != nil {
        log.Println("redis connect error:", conn.Err().Error())
        return false
    }
    _, err := redis.String(conn.Do("ping"))
    if err != nil {
        log.Println("redis connect ping error:", err.Error())
        return false
    }
    return true
}
