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
	addressMap := map[string]domain.Address{
		"CACHED_ADDRESS": {Type: "TestType", FormattedAddress: "TestFormattedAddr"},
	}
	address, ok := addressMap[key]
	if !ok {
		return nil, nil
	}
	return &address, nil
}

func (r mockRepo) SetAddress(key string, address *domain.Address) error {
	return nil
}
