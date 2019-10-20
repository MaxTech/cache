package redis_cache

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "os"
    "time"
)

type redisPoolUtils struct {
}

var RedisPoolUtils *redisPoolUtils

type RedisPool interface {
    CheckRedisPool() bool
}

type pool redis.Pool

func (rpu *redisPoolUtils) InitRedisPool(_address string, _password string, _dbNum int) RedisPool {
    temp := pool(redis.Pool{
        MaxIdle:     500,
        MaxActive:   50,
        IdleTimeout: 60 * time.Second,
        Wait:        true,
        Dial: func() (redis.Conn, error) {
            con, err := redis.Dial("tcp", _address,
                redis.DialPassword(_password),
                redis.DialDatabase(_dbNum),
                redis.DialConnectTimeout(2*time.Second),
                redis.DialReadTimeout(2*time.Second),
                redis.DialWriteTimeout(2*time.Second))
            if err != nil {
                return nil, err
            }
            return con, nil
        },
    })
    return &temp
}

func (rpu *redisPoolUtils) InitRedisPoolByConfig(_redisConfig RedisConfigFormat) RedisPool {
    temp := pool(redis.Pool{
        MaxIdle:     500,
        MaxActive:   50,
        IdleTimeout: 60 * time.Second,
        Wait:        true,
        Dial: func() (redis.Conn, error) {
            con, err := redis.Dial("tcp", _redisConfig.Address,
                redis.DialPassword(_redisConfig.Password),
                redis.DialDatabase(_redisConfig.DBNum),
                redis.DialConnectTimeout(2*time.Second),
                redis.DialReadTimeout(2*time.Second),
                redis.DialWriteTimeout(2*time.Second))
            if err != nil {
                return nil, err
            }
            return con, nil
        },
    })
    return &temp
}

func (r *pool) CheckRedisPool() bool {
    conn, err := r.Dial()
    if err != nil {
        return false
    }
    defer conn.Close()
    if conn.Err() != nil {
        _, _ = fmt.Fprintln(os.Stderr, time.Now().Format(time.RFC3339Nano), "[Error]", "redis pool connect error:", err)
        return false
    }
    _, err = redis.String(conn.Do("ping"))
    if err != nil {
        _, _ = fmt.Fprintln(os.Stderr, time.Now().Format(time.RFC3339Nano), "[Error]", "redis pool connect ping error:", err)
        return false
    }
    return true
}
