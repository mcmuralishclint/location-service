package adapter

import (
	"context"
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
	req := &maps.PlaceDetailsRequest{PlaceID: id}
	resp, err := r.client.PlaceDetails(context.Background(), req)
	if err != nil {
		return nil, err
	}
	address := &domain.Address{
		Type:             "google maps",
		FormattedAddress: resp.FormattedAddress,
	}
	return address, nil
}

func (r *GoogleMapsRepository) QueryAutoComplete(input string) ([]domain.AutocompletePrediction, error) {
	req := &maps.QueryAutocompleteRequest{Input: input}
	resp, err := r.client.QueryAutocomplete(context.Background(), req)
	if err != nil {
		return []domain.AutocompletePrediction{}, err
	}
	autoCompletePredictions := []domain.AutocompletePrediction{}
	for _, address := range resp.Predictions {
		autoCompletePredictions = append(autoCompletePredictions, domain.AutocompletePrediction{PlaceID: address.PlaceID, Description: address.Description})
	}
	return autoCompletePredictions, err
}
