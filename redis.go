package operations

import (
	"fmt"
	//
	"github.com/gomodule/redigo/redis"
	"gitlab.com/TheDarkula/jsonrouter/logging/testing"
)

const (
	// redis
	REDIS_MAX_IDLE_CONN = 10
	REDIS_PORT = "6379"
)

type RedisClient struct {
	*redis.Pool
	Log *logs.Logger
}

func (self *RedisClient) NewConn() *Context {
	return &Context{
		self.Get(),
		self.Log,
	}
}

type Context struct {
	redis.Conn
	Log *logs.Logger
}

func (self *Context) Close() {
	self.Conn.Close()
}

func NewClient(redisHost, redisPassword string) *RedisClient {

	client := &RedisClient{
		Pool: redis.NewPool(
			func() (redis.Conn, error) {
				return redis.Dial(
					"tcp",
					fmt.Sprintf("%s:%s", redisHost, REDIS_PORT),
					redis.DialPassword(redisPassword),
				)
			},
			REDIS_MAX_IDLE_CONN,
		),
		Log: logs.NewClient().NewLogger(""),
	}

	return client
}
