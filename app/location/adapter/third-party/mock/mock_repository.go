package mock

import (
	"errors"
	"sync"

	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type mockRepository struct {
	sync.Mutex
	addresses   map[string]*domain.Address
	suggestions map[string][]domain.AutocompletePrediction
}

func NewMockRepository() domain.AddressRepository {
	return &mockRepository{
		addresses: map[string]*domain.Address{
			"GA123": {Type: "TestType", FormattedAddress: "123"},
		},
		suggestions: map[string][]domain.AutocompletePrediction{
			"test_addr": {domain.AutocompletePrediction{Description: "Address 1", PlaceID: "GA123"},
				domain.AutocompletePrediction{Description: "Address 2", PlaceID: "GA234"}},
		},
	}
}

func (r *mockRepository) QueryAutoComplete(input string) ([]domain.AutocompletePrediction, error) {
	addresses, ok := r.suggestions[input]
	if !ok {
		return nil, errors.New("No suggestions found")
	}
	if addresses != nil {
		return addresses, nil
	}
	return []domain.AutocompletePrediction{}, nil
}

func (r *mockRepository) GetByID(id string) (*domain.Address, error) {
	address, ok := r.addresses[id]
	if !ok {
		return nil, errors.New("address not found")
	}
	return address, nil
}
