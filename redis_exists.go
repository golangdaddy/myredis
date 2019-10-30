package operations

import (
	"github.com/gomodule/redigo/redis"
)

// Exists returns whether the specified key exists
func (ctx *Context) Exists(key string) (bool, error) {
	return redis.Bool(ctx.Do("EXISTS", key))
}
