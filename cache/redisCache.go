package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *RedisCache) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(key string, entity interface{}) error {
	client := cache.GetClient()

	json, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	if err = client.Set(context.Background(), key, json, cache.expires).Err(); err != nil {
		return err
	}
	return nil
}

func (cache *RedisCache) Get(key string) (*interface{}, error) {
	client := cache.GetClient()
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	var entity interface{}
	json.Unmarshal([]byte(val), &entity)

	return &entity, nil
}

func (cache *RedisCache) RPush(key string, entity interface{}) error {
	client := cache.GetClient()

	json, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	err = client.RPush(context.Background(), key, json).Err()
	return err
}

func (cache *RedisCache) LRange(key string) (*[]string, error) {
	client := cache.GetClient()
	values, err := client.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return &values, nil
}
