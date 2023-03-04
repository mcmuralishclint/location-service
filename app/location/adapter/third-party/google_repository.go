package third_party

import (
	"context"
	"encoding/json"
	"github.com/mcmuralishclint/location-service/app/location/domain"
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

	var result []map[string]interface{}
	jsonData, err := json.Marshal(resp.AddressComponents)
	json.Unmarshal([]byte(jsonData), &result)
	addressComponents := ParseAddressComponents(result)

	address = &domain.Address{
		Type:              "google maps",
		FormattedAddress:  resp.FormattedAddress,
		AddressComponents: addressComponents,
	}

	r.cacheRepo.SetAddress(id, address)
	return address, nil
}

func (r *GoogleMapsRepository) QueryAutoComplete(input string) ([]domain.AutocompletePrediction, error) {
	req := &maps.PlaceAutocompleteRequest{Input: input, Components: map[maps.Component][]string{maps.ComponentCountry: {r.country}}}

	resp, err := r.client.PlaceAutocomplete(context.Background(), req)
	if err != nil {
		return []domain.AutocompletePrediction{}, err
	}
	autoCompletePredictions := []domain.AutocompletePrediction{}
	for _, address := range resp.Predictions {
		autoCompletePredictions = append(autoCompletePredictions, domain.AutocompletePrediction{PlaceID: address.PlaceID, Description: address.Description})
	}
	return autoCompletePredictions, err
}

func ParseAddressComponents(components []map[string]interface{}) *domain.AddressComponents {
	address := &domain.AddressComponents{}
	for _, component := range components {
		types := component["types"].([]interface{})
		switch types[0] {
		case "subpremise":
			address.Subpremise = component["long_name"].(string)
		case "street_number":
			address.StreetNumber = component["long_name"].(string)
		case "route":
			address.Route = component["long_name"].(string)
		case "locality":
			address.Locality = component["long_name"].(string)
		case "administrative_area_level_3":
			address.AdministrativeArea3 = component["long_name"].(string)
		case "administrative_area_level_2":
			address.AdministrativeArea2 = component["short_name"].(string)
		case "administrative_area_level_1":
			address.AdministrativeArea1 = component["long_name"].(string)
		case "country":
			address.Country = component["long_name"].(string)
		case "postal_code":
			address.PostalCode = component["long_name"].(string)
		}
	}
	return address
}
