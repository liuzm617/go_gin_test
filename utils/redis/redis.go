package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"go_gin_example/utils/conf"
	"time"
)

var RedisConn *redis.Pool

func Init() {
	RedisConn = &redis.Pool{
		MaxIdle:     conf.RedisConf.MaxIdle,
		MaxActive:   conf.RedisConf.MaxActive,
		IdleTimeout: conf.RedisConf.IdleTimeout,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", conf.RedisConf.Host)
			if err != nil {
				return nil, err
			}
			if conf.RedisConf.Password != "" {
				if _, err := conn.Do("AUTH", conf.RedisConf.Password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// set key/value
func Set(key string, value interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, v)
	if time != 0{
		_, err = conn.Do("EXPIRE", key, time)
	}

	if err != nil {
		return err
	}
	return nil
}

// get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

//delete a key
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// check a key is exists
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

// incr a key
func Incr(key string) error{
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("INCR", key)
	if err != nil{
		return err
	}
	return nil

}


// incr key by num
func IncrBy(key string, num int) error{
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("INCRBY", key, num)
	if err != nil{
		return err
	}
	return nil
}

// decr a key
func Decr(key string)error{
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("DECR", key)
	if err != nil{
		return err
	}
	return nil
}

func DecrBy(key string)error{
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("DECRBY", key)
	if err != nil{
		return err
	}
	return nil
}
