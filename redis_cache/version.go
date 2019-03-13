package redis_cache

import "fmt"

var Version string

func init() {
    Version = fmt.Sprintf(
        "|- %s module:\t\t\t%s",
        "redis cache",
        "0.0.1")
}
