package myredis

import (
	"github.com/gomodule/redigo/redis"
)

// Delete uses the DEL command to remove any keys featuring in the supplied slice.
func (ctx *Context) Delete(keys ...interface{}) error {

	_, err := ctx.Do("DEL", keys...)
	if err != nil {
		return err
	}
	return nil
}


// Delete uses the DEL command to remove any keys featuring in the supplied slice.
func (ctx *Context) DeleteOrFail(keys ...interface{}) bool {

	n, err := redis.Int(
		ctx.Do("DEL", keys...),
	)
	if ctx.Log.Error(err) {
		return false
	}
	return n == len(keys)
}
