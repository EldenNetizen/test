package database

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_addr     = "127.0.0.1:6379"
	_password = ""
	_db       = 1
)

type RedisManager struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisManager() *RedisManager {
	return &RedisManager{}
}

func (rm *RedisManager) Connect() {
	rm.client = redis.NewClient(&redis.Options{
		Addr:     _addr,
		Password: _password,
		DB:       _db,
	})
	rm.ctx = context.Background()
}

func (rm *RedisManager) Close() error {
	return rm.client.Close()
}

func (rm *RedisManager) Set(key string, value string, expiration time.Duration) error {
	return rm.client.Set(rm.ctx, key, value, expiration).Err()
}

func (rm *RedisManager) Get(key string) (string, error) {
	return rm.client.Get(rm.ctx, key).Result()
}

func (rm *RedisManager) Delete(key string) (int64, error) {
	return rm.client.Del(rm.ctx, key).Result()
}
