package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GabrielEValenzuela/RefuCare/internal/core/domain"
	"github.com/redis/go-redis/v9"
)

type RedisVitalsRepository struct {
	client *redis.Client
}

var redisKeyPrefix = "vitals:"

func NewRedisVitalsRepository() *RedisVitalsRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // or from config
		DB:   0,
	})
	return &RedisVitalsRepository{client: rdb}
}

func (r *RedisVitalsRepository) SaveVitals(v domain.Vitals) error {
	ctx := context.Background()

	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to serialize vitals: %w", err)
	}

	key := redisKeyPrefix + v.ID
	return r.client.Set(ctx, key, data, 24*time.Hour).Err()
}

func (r *RedisVitalsRepository) GetVitalsByID(id string) (*domain.Vitals, error) {
	ctx := context.Background()

	data, err := r.client.Get(ctx, redisKeyPrefix+id).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("vitals not found")
	} else if err != nil {
		return nil, err
	}

	var v domain.Vitals
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		return nil, fmt.Errorf("failed to parse vitals: %w", err)
	}

	return &v, nil
}
