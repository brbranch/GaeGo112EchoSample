package client

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"strings"
)

type redisClient struct {
	pool *redis.Pool
}

var client *redisClient = nil

func GetRedisClient() *redisClient {
	if client == nil {
		panic("client is not initialized")
	}
	return client
}

func InitRedis(endpoint string, password string) {
	if client != nil {
		return
	}
	client = &redisClient{}
	client.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", endpoint)
			if password == "" {
				return conn, err
			}
			if err != nil {
				return nil, err
			}
			// パスワードが設定されてる場合のみ認証する(ローカルだと不要なため)
			if password != "" {
				if _, err := conn.Do("AUTH", password); err != nil {
					conn.Close()
					return nil, err
				}
			} else if (strings.HasPrefix(endpoint, "redis")) {
				// EndpointがRedusCloudなのにパスワードがないのはおかしい
				conn.Close()
				return nil, errors.New("invalid password.")
			}
			return conn, nil
		},
	}
}

func (c *redisClient) GetConnection() redis.Conn {
	return c.pool.Get()
}

func (c *redisClient) PutString(key string, value string) error {
	con := c.GetConnection()
	defer con.Close()
	_, err := con.Do("SET", key, value)
	return err
}

func (c *redisClient) GetString(key string) (string, error) {
	con := c.GetConnection()
	defer con.Close()

	return redis.String(con.Do("GET", key))
}
