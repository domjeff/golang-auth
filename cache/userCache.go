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
	return NewRedisCache("localhost:6379", 0, time.Duration(10))
}

func (cache *RedisCache) CheckUserToken(user models.User) error {
	key := fmt.Sprintf("user%d.tokens", user.Id)
	tokens, err := cache.LRange(key)
	if err != nil {
		return err
	}

	fmt.Println(*tokens)
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

func (cache *RedisCache) UpdateUserTokens(
	user models.User,
	cookieToken string,
	generateFunction func(user models.User) (*string, error),
) error {
	tokens, err := cache.getUserTokens(user)
	if err != nil {
		return err
	}
	for _, token := range *tokens {
		if token == fmt.Sprintf("\"%s\"", cookieToken) {
			fmt.Println("Successs to here")
			key := fmt.Sprintf("user%d.token", user.Id)
			newToken, err := generateFunction(user)
			if err != nil {
				return err
			}
			fmt.Println(token)
			fmt.Println(*newToken)
			err = cache.LSet(key, 1, *newToken)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}
