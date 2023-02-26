package adapter

import (
	"context"
	"fmt"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"googlemaps.github.io/maps"
)

type GoogleMapsRepository struct {
	client *maps.Client
}

func NewGoogleMapsRepository(apiKey string) (*GoogleMapsRepository, error) {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &GoogleMapsRepository{client: c}, nil
}

func (r *GoogleMapsRepository) GetByID(id string) (*domain.Address, error) {
	req := &maps.GeocodingRequest{PlaceID: id}
	resp, err := r.client.Geocode(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, fmt.Errorf("address not found")
	}
	address := &domain.Address{
		Type:             "google maps",
		FormattedAddress: resp[0].FormattedAddress,
	}
	return address, nil
}
