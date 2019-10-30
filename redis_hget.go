package myredis

import (
	"strings"
	//
	"github.com/gomodule/redigo/redis"
)

// HGetAll retrieves all values from the map
func (ctx *Context) HGetAll(hash string) (map[string]string, error) {
	return redis.StringMap(
		ctx.Do("HGETALL", hash),
	)
}

// HGetString retrieves the map's key's value as a bool
func (ctx *Context) HGetBool(m, key string) (bool, bool) {
	value, err := redis.Bool(
		ctx.Do("HGET", m, key),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, false
	}
	return true, value
}

// HGetString retrieves the map's key's value as a string
func (ctx *Context) HGetString(m, key string) (bool, string) {
	value, err := redis.String(
		ctx.Do("HGET", m, key),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, ""
	}
	return true, value
}
