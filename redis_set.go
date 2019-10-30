package myredis

import (
	//"github.com/gomodule/redigo/redis"
)

// SetAndExpire uses the SET command to store an interface{} against the supplied key.
// Once the key is set, the expire time is modified using the EXPIRE command.
func (ctx *Context) SetAndExpire(key string, value interface{}, expireSeconds int) error {

	_, err := ctx.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = ctx.Do("EXPIRE", key, expireSeconds)
	if err != nil {
		return err
	}

	return nil
}

// Set uses the SET command to store an interface{} against the supplied key.
func (ctx *Context) Set(key string, value interface{}) error {

	_, err := ctx.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// HSet uses the HSET command to store an interface{} against the supplied key, in the supplied map.
func (ctx *Context) HSet(obj, key string, value interface{}) error {

	_, err := ctx.Do("HSET", obj, key, value)
	if err != nil {
		return err
	}

	return nil
}
