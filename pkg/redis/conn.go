package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewStorage(host, pass string) *Redis {
	return &Redis{client: redis.NewClient(&redis.Options{
		DB:       0,
		Password: pass,
		Addr:     host,
	})}
}

func (r Redis) Put(ctx context.Context, key string, link string) error {
	return r.client.Set(ctx, key, link, 0).Err()
}

func (r Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
