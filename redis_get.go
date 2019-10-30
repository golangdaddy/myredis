package operations

import (
	"strings"
	//
	"github.com/gomodule/redigo/redis"
)

// GetFloat64 returns a key's value as float64
func (ctx *Context) GetFloat64(key string) (bool, float64) {
	value, err := redis.Float64(
		ctx.Do("GET", key),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, 0.0
	}
	return true, value
}

// GetInt64 returns a key's value as int64
func (ctx *Context) GetInt64(key string) (bool, int64) {
	value, err := redis.Int64(
		ctx.Do("GET", key),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, 0.0
	}
	return true, value
}

// GetString returns a key's value as string
func (ctx *Context) GetString(key string) (bool, string) {
	value, err := redis.String(
		ctx.Do("GET", key),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, ""
	}
	return true, value
}

// GetBytes returns a key's value as byte slice
func (ctx *Context) GetBytes(key string) (bool, []byte) {
	value, err := redis.Bytes(
		ctx.Do("GET", key),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, nil
	}
	return true, value
}

// GetBool returns a key's value as boolean
func (ctx *Context) GetBool(key string) (bool, bool) {
	value, err := redis.Bool(
		ctx.Do("GET", key),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return false, false
	}
	return true, value
}
