package db

import (
	"fmt"
	"time"

	"net"
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

var once sync.Once

func Init() {
	if redisClient == nil {

		once.Do(func() {
			ipSlice, err := net.LookupIP("127.0.0.1")
			if err != nil {
				fmt.Errorf("Invalid Database URL, Err: ", err)
				os.Exit(1)
				return
			}

			redisDB, _ := strconv.Atoi("1")
			redisClient = redis.NewClient(&redis.Options{
				Addr: fmt.Sprintf("%s:%d", ipSlice[0].String(), 6379),

				Password:     "",      // no password set
				DB:           redisDB, // use default DB
				ReadTimeout:  10 * time.Second,
				MinIdleConns: 10,
			})

			_, err = redisClient.Ping().Result()
			if err != nil {
				fmt.Print("Could not connect to redis", err)
				os.Exit(1)
				return
			}

		})
	}
}

type RedisConn struct {
	conn *redis.Client
}

func New() *RedisConn {
	Init()
	return &RedisConn{
		conn: redisClient,
	}
}

// This should return an interface!!!
func (r *RedisConn) GetQueryer() *redis.Client {
	return r.conn
}

// Initialize the Redis connection and assign the existing redis connection
func (r *RedisConn) Init() {
	r.conn = redisClient
}
