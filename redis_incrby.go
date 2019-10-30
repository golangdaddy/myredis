package operations

import (
	"strings"
)

// Incrby increments the key's value with an integer
func (ctx *Context) Incrby(key string, n int) error {
	_, err := ctx.Do("INCRBY", key, n)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return err
	}
	return nil
}

// Incrbyfloat increments the key's value by a float value
func (ctx *Context) Incrbyfloat(key string, n float64) error {
	_, err := ctx.Do("INCRBYFLOAT", key, n)
	if err != nil {
		if !strings.Contains(err.Error(), "nil returned") {
			ctx.Log.Error(err)
		}
		return err
	}
	return nil
}
