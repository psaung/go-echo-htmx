package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	redisClient redis.UniversalClient
}

func NewRedisCache(redisClient redis.UniversalClient) *RedisCache {
	return &RedisCache{redisClient: redisClient}
}

func (r RedisCache) Get(ctx context.Context, key string) (*string, error) {
	result, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (r RedisCache) Set(ctx context.Context, key, value string) error {
	ttl := 24 * time.Hour
	return r.redisClient.Set(ctx, key, value, ttl).Err()
}

func (r RedisCache) Delete(ctx context.Context, key string) error {
	return r.redisClient.Del(ctx, key).Err()
}
