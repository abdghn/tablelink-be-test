package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"tablelink-be-test/internal/domain/entity"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{client: rdb}
}

func (r *RedisCache) SetUserSession(ctx context.Context, token string, userSession *entity.UserSession) error {
	data, err := json.Marshal(userSession)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, "session:"+token, data, 24*time.Hour).Err()
}

func (r *RedisCache) GetUserSession(ctx context.Context, token string) (*entity.UserSession, error) {
	data, err := r.client.Get(ctx, "session:"+token).Bytes()
	if err != nil {
		return nil, err
	}

	var userSession entity.UserSession
	if err := json.Unmarshal(data, &userSession); err != nil {
		return nil, err
	}
	return &userSession, nil
}

func (r *RedisCache) DeleteUserSession(ctx context.Context, token string) error {
	return r.client.Del(ctx, "session:"+token).Err()
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}
