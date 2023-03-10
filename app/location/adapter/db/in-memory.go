package db

import (
	"encoding/json"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"sync"
	"time"
)

type inMemRepo struct {
	value   sync.Map
	expires int64
}

func NewInMemRepo() domain.CacheRepository {
	return &inMemRepo{}
}

func (r *inMemRepo) GetAddress(key string) (*domain.Address, error) {
	if r.Expired(time.Now().UnixNano()) {
		return nil, nil
	}
	value, ok := r.value.Load(key)
	if !ok {
		return nil, nil
	}
	var address *domain.Address
	bytes, _ := json.Marshal(value)
	err := json.Unmarshal(bytes, &address)
	if err != nil {
		return nil, nil
	}
	return address, nil
}

func (r *inMemRepo) SetAddress(key string, address *domain.Address) error {
	r.value.Store(key, address)
	r.expires = time.Now().UnixNano() + int64(5*time.Minute)
	return nil
}

func (r *inMemRepo) Expired(time int64) bool {
	if r.expires == 0 {
		return false
	}
	return time > r.expires
}
