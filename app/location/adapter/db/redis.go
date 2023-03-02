package db

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"time"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) domain.RedisRepository {
	return &redisRepo{client: client}
}

func (r redisRepo) GetAddress(key string) (*domain.Address, error) {
	data, err := r.client.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	var address *domain.Address
	err = json.Unmarshal(data, &address)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (r redisRepo) SetAddress(key string, address *domain.Address) error {
	data, err := json.Marshal(address)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, data, 10*time.Minute).Err()
}
