package google

import (
	"context"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/mcmuralishclint/location-service/util"
	"googlemaps.github.io/maps"
)

type GoogleMapsRepository struct {
	client    *maps.Client
	country   string
	cacheRepo domain.CacheRepository
}

func NewGoogleMapsRepository(apiKey string, country string, cacheRepo domain.CacheRepository) (*GoogleMapsRepository, error) {
	c, err := maps.NewClient(
		maps.WithAPIKey(apiKey),
	)
	if err != nil {
		return nil, err
	}
	return &GoogleMapsRepository{client: c, country: country, cacheRepo: cacheRepo}, nil
}

func (r *GoogleMapsRepository) GetByID(id string) (*domain.Address, error) {
	address, _ := r.cacheRepo.GetAddress(id)
	if address != nil {
		return address, nil
	}

	req := &maps.PlaceDetailsRequest{PlaceID: id}
	resp, err := r.client.PlaceDetails(context.Background(), req)
	if err != nil {
		return nil, err
	}

	addressFormatter := NewGoogleAddressFormatter()
	address, err = addressFormatter.FormatAddress(resp.AddressComponents)
	address.FormattedAddress = resp.FormattedAddress
	util.PopulateAddressCountry(address, r.country)

	r.cacheRepo.SetAddress(id, address)
	return address, nil
}

func (r *GoogleMapsRepository) QueryAutoComplete(input string) ([]domain.AutocompletePrediction, error) {
	req := &maps.PlaceAutocompleteRequest{Input: input, Components: map[maps.Component][]string{maps.ComponentCountry: {r.country}}}

	resp, err := r.client.PlaceAutocomplete(context.Background(), req)
	if err != nil {
		return []domain.AutocompletePrediction{}, err
	}

	addressFormatter := NewGoogleAddressFormatter()
	predictions, err := addressFormatter.FormatAddressSuggestion(resp)

	return predictions, err
}
