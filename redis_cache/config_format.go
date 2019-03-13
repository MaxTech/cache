package redis_cache

type RedisConfigFormat struct {
    Address  string `yaml:"address"`
    Password string `yaml:"password"`
    DBNum    int    `yaml:"db_num"`
}
