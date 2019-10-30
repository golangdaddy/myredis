package operations

import (
	"strings"
)

// Incrby increments the map's key's value with an integer
func (ctx *Context) HIncrby(hashtable, key string, n int) error {
	_, err := ctx.Do("HINCRBY", hashtable, key, n)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return err
	}
	return nil
}

// Incrby increments the map's key's value with an float
func (ctx *Context) HIncrbyfloat(hashtable, key string, n float64) error {
	_, err := ctx.Do("HINCRBYFLOAT", hashtable, key, n)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return err
	}
	return nil
}
