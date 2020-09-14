package redis

import (
	"fmt"
	"strconv"

	"github.com/chnejohnson/shortener/domain"
	"github.com/go-redis/redis"
)

// RedirectRepository ...
type redisRedirectRepository struct {
	client *redis.Client
}

// NewRedisRedirectRepository ...
func NewRedisRedirectRepository(client *redis.Client) domain.RedirectRepository {
	return &redisRedirectRepository{client}
}

// Find ...
func (r *redisRedirectRepository) Find(code string) (*domain.Redirect, error) {
	var rdrt domain.Redirect

	data, err := r.client.HGetAll(code).Result()
	if err != nil {
		return nil, err
	}

	createdAt, err := strconv.ParseInt(data["CreatedAt"], 10, 64)
	if err != nil {
		return nil, err
	}

	rdrt.URL = data["URL"]
	rdrt.Code = data["Code"]
	rdrt.CreatedAt = createdAt

	return &rdrt, nil
}

// Store ...
func (r *redisRedirectRepository) Store(rdrt *domain.Redirect) error {
	data := map[string]interface{}{
		"URL":       rdrt.URL,
		"Code":      rdrt.Code,
		"CreatedAt": rdrt.CreatedAt,
	}

	_, err := r.client.HMSet(rdrt.Code, data).Result()
	if err != nil {
		return err
	}

	fmt.Println("成功存入Redis")
	return nil
}
