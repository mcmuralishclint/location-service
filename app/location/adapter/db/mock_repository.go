package db

import (
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type mockRepo struct {
}

func NewMockRepo(host string, password string, db int) domain.CacheRepository {
	return &mockRepo{}
}

func (r mockRepo) GetAddress(key string) (*domain.Address, error) {
	return nil, nil
}

func (r mockRepo) SetAddress(key string, address *domain.Address) error {
	return nil
}
