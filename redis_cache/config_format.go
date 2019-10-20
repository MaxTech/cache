package redis_cache

type RedisConfigFormat struct {
    Address  string `json:"address" yaml:"address"`
    Password string `json:"password" yaml:"password"`
    DBNum    int    `json:"db_num" yaml:"db_num"`
}
