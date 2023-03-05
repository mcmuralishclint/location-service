package google

import (
	"encoding/json"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/mcmuralishclint/location-service/app/location/formatter"
	"googlemaps.github.io/maps"
)

type GoogleAddressFormatter struct{}

func NewGoogleAddressFormatter() formatter.AddressFormatter {
	return &GoogleAddressFormatter{}
}

func (f *GoogleAddressFormatter) FormatAddress(data interface{}) (*domain.Address, error) {
	var result []map[string]interface{}
	jsonData, _ := json.Marshal(data)
	json.Unmarshal([]byte(jsonData), &result)

	addressComponent := &domain.AddressComponents{}
	for _, component := range result {
		types := component["types"].([]interface{})
		switch types[0] {
		case "street_number":
			addressComponent.StreetNumber = component["long_name"].(string)
		case "route":
			addressComponent.StreetName = component["long_name"].(string)
		case "locality":
			addressComponent.City = component["long_name"].(string)
		case "country":
			addressComponent.Country = component["long_name"].(string)
		case "postal_code":
			addressComponent.PostalCode = component["long_name"].(string)
		}
	}
	address := &domain.Address{
		Type:              "google maps",
		AddressComponents: addressComponent,
	}
	return address, nil
}

func (f *GoogleAddressFormatter) FormatAddressSuggestion(data interface{}) ([]domain.AutocompletePrediction, error) {
	predictionData := data.(maps.AutocompleteResponse)
	autoCompletePredictions := []domain.AutocompletePrediction{}
	for _, address := range predictionData.Predictions {
		autoCompletePredictions = append(autoCompletePredictions, domain.AutocompletePrediction{PlaceID: address.PlaceID, Description: address.Description})
	}
	return autoCompletePredictions, nil
}
