package operations

import (
	"github.com/gomodule/redigo/redis"
)

// GetKeysInt64 returns keys matching regex
func (ctx *Context) GetKeysInt64(regex string) map[string]int64 {
	values := map[string]int64{}
	keys, err := redis.Strings(
		ctx.Do("KEYS", regex),
	)
	ctx.Log.Error(err)
	for _, key := range keys {
		if exists, value := ctx.GetInt64(key); exists {
			values[key] = value
		}
	}
	return values
}

// GetKeysFloat64 returns keys matching regex
func (ctx *Context) GetKeysFloat64(regex string) map[string]float64 {
	values := map[string]float64{}
	keys, err := redis.Strings(
		ctx.Do("KEYS", regex),
	)
	ctx.Log.Error(err)
	for _, key := range keys {
		if exists, value := ctx.GetFloat64(key); exists {
			values[key] = value
		}
	}
	return values
}

// GetKeysString returns keys matching regex
func (ctx *Context) GetKeysString(regex string) map[string]string {
	values := map[string]string{}
	keys, err := redis.Strings(
		ctx.Do("KEYS", regex),
	)
	ctx.Log.Error(err)
	for _, key := range keys {
		if exists, value := ctx.GetString(key); exists {
			values[key] = value
		}
	}
	return values
}
