package cache

import (
	"context"
	"encoding/json"
	"time"
)

func (cache *RedisCache) Set(key string, entity interface{}) error {
	client := cache.GetClient()

	json, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
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
