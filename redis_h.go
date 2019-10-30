package myredis

import (
	"strings"
	//
	"github.com/gomodule/redigo/redis"
)

// HKeys retrieves a slice of keys from the map
func (ctx *Context) HKeys(hash string) (bool, []string) {
	value, err := redis.Strings(
		ctx.Do("HKEYS", hash),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, nil
	}
	return true, value
}

// HDel removes a key from the map
func (ctx *Context) HDel(hash, key string) bool {
	_, err := ctx.Do("HDEL", hash, key)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false
	}
	return true
}
