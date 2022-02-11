/*================================================================
*
*  文件名称：redis.go
*  创 建 者: mongia
*  创建日期：2022年01月04日
*
================================================================*/

package db

import (
	"time"

	"github.com/gomodule/redigo/redis"
	redigo "github.com/gomodule/redigo/redis"
)

var (
	pool *redigo.Pool = nil
)

func RedisInit(server, password string, conns int) {
	pool = &redigo.Pool{
		MaxIdle:     conns / 2,
		IdleTimeout: 240 * time.Second,
		MaxActive:   conns,
		Dial: func() (redigo.Conn, error) {
			return redigo.Dial("tcp", server, redigo.DialReadTimeout(2*time.Second), redigo.DialPassword(password))
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxConnLifetime: 0, // connection not close when this server run
	}
}

func GetString(key string) (string, error) {
	c := pool.Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return "", c.Err()
	}

	return redis.String(c.Do("GET", key))
}

func Exists(key string) (bool, error) {
	c := pool.Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return false, c.Err()
	}

	return redis.Bool(c.Do("EXISTS", key))
}

func HMGet(ret interface{}, key string, field ...string) error {
	c := pool.Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return c.Err()
	}

	v, err := redis.Values(c.Do("HMGET", key, field))
	if nil != err {
		return err
	}

	err = redis.ScanStruct(v, ret)
	if nil != err {
		return err
	}
	return nil
}

func HGetALL(ret interface{}, key string) error {
	c := pool.Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return c.Err()
	}

	v, err := redis.Values(c.Do("HGETALL", key))
	if nil != err {
		return err
	}

	err = redis.ScanStruct(v, ret)
	if nil != err {
		return err
	}
	return nil
}

func HMSet(key string, val interface{}) error {
	c := pool.Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return c.Err()
	}

	_, err := c.Do("HMSET", redis.Args{}.Add(key).AddFlat(val))
	if nil != err {
		return err
	}

	return nil
}

func Remove(key []interface{}) (bool, error) {
	c := pool.Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return false, c.Err()
	}

	c.Send("MULTI")

	for _, v := range key {
		c.Send("DEL", v)
	}

	_, err := redis.Values(c.Do("EXEC"))
	if nil != err {
		return false, err
	}
	return true, nil
}

func GetRedisPool() *redis.Pool {
	return pool
}
