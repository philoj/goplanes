package redisinfra

import (
	redis "github.com/redis/go-redis/v9"
)

func NewClient(uri string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: uri,
	})
}
