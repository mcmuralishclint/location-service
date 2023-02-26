package adapter

import (
	"errors"
	"sync"

	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type mockRepository struct {
	sync.Mutex
	addresses map[string]*domain.Address
}

func NewMockRepository() domain.AddressRepository {
	return &mockRepository{
		addresses: map[string]*domain.Address{
			"GA123": {Type: "TestType", FormattedAddress: "123"},
		},
	}
}

func (r *mockRepository) GetByID(id string) (*domain.Address, error) {
	r.Lock()
	defer r.Unlock()

	address, ok := r.addresses[id]
	if !ok {
		return nil, errors.New("address not found")
	}
	return address, nil
}
