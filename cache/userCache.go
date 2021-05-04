package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/domjeff/golang-auth/models"
)

type UserCache interface {
	Set(key string, entity interface{})
	Get(key string) interface{}
}

func SetupUserCache() *RedisCache {
	return NewRedisCache("localhost:6379", 0, time.Duration(time.Hour*24))
}

func (cache *RedisCache) CheckUserToken(user models.User) error {
	key := fmt.Sprintf("user%d.tokens", user.Id)
	tokens, err := cache.LRange(key)
	if err != nil {
		return err
	}

	if len(*tokens) > 0 {
		return errors.New("Number of session reached limit already")
	}
	return nil
}

func (cache *RedisCache) SetUserToken(user models.User, token string) error {
	key := fmt.Sprintf("user%d.tokens", user.Id)
	return cache.RPush(key, token)
}

func (cache *RedisCache) getUserTokens(user models.User) (*[]string, error) {
	key := fmt.Sprintf("user%d.tokens", user.Id)
	tokens, err := cache.LRange(key)
	if err != nil {
		return nil, err
	}

	if len(*tokens) == 0 {
		return nil, errors.New("Number of session reached limit already")
	}
	return tokens, nil
}

func (cache *RedisCache) UpdateUserTokens(user models.User, cookieToken string) error {
	tokens, err := cache.getUserTokens(user)
	if err != nil {
		return err
	}
	// updatedToken := []string{}
	for _, token := range *tokens {
		if token == cookieToken {
			// updatedToken = append(updatedToken, cookieToken)

			break
		}
	}
	return nil
}
