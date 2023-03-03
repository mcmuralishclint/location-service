package service

import (
	third_party "github.com/mcmuralishclint/location-service/app/location/adapter/third-party"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLocationService(t *testing.T) {
	locationProvider := third_party.NewMockRepository()
	service := NewLocationService(locationProvider)
	assert.Equal(t, service, &locationService{locationRepository: locationProvider})
}

func TestLocationService_GetAddressByID(t *testing.T) {
	locationProvider := third_party.NewMockRepository()
	service := NewLocationService(locationProvider)

	// right address
	id := "GA123"
	address, err := service.GetAddressByID(id)
	assert.NoError(t, err)
	assert.Equal(t, address, &domain.Address{Type: "TestType", FormattedAddress: "123"})

	// invalid address
	id = "INVALID"
	address, err = service.GetAddressByID(id)
	assert.NoError(t, err)
	assert.Equal(t, address, (*domain.Address)(nil))
}

func TestLocationService_GetQueryAutoCompleteByText(t *testing.T) {
	locationProvider := third_party.NewMockRepository()
	service := NewLocationService(locationProvider)

	// valid suggestion
	id := "test_addr"
	addresses, err := service.GetQueryAutoCompleteByText(id)
	assert.NoError(t, err)
	assert.Equal(t, addresses, []domain.AutocompletePrediction([]domain.AutocompletePrediction{domain.AutocompletePrediction{Description: "Address 1", PlaceID: "GA123"}, domain.AutocompletePrediction{Description: "Address 2", PlaceID: "GA234"}}))

	// invalid suggestion
	id = "invalid"
	addresses, err = service.GetQueryAutoCompleteByText(id)
	assert.Equal(t, err.Error(), "No suggestions found")
	assert.Equal(t, addresses, []domain.AutocompletePrediction([]domain.AutocompletePrediction(nil)))
}
