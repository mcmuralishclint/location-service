package google

import (
	"encoding/json"
	"fmt"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"googlemaps.github.io/maps"
)

type GoogleAddressFormatter struct{}

func NewGoogleAddressFormatter() *GoogleAddressFormatter {
	return &GoogleAddressFormatter{}
}

func (f *GoogleAddressFormatter) FormatAddress(data interface{}) (*domain.AddressComponents, error) {
	fmt.Println(data)
	var result []map[string]interface{}
	jsonData, _ := json.Marshal(data)
	json.Unmarshal([]byte(jsonData), &result)

	address := &domain.AddressComponents{}
	for _, component := range result {
		types := component["types"].([]interface{})
		switch types[0] {
		case "street_number":
			address.StreetNumber = component["long_name"].(string)
		case "route":
			address.StreetName = component["long_name"].(string)
		case "locality":
			address.City = component["long_name"].(string)
		case "country":
			address.Country = component["long_name"].(string)
		case "postal_code":
			address.PostalCode = component["long_name"].(string)
		}
	}
	return address, nil
}

func (f *GoogleAddressFormatter) FormatAddressSuggestion(data maps.AutocompleteResponse) ([]domain.AutocompletePrediction, error) {
	autoCompletePredictions := []domain.AutocompletePrediction{}
	for _, address := range data.Predictions {
		autoCompletePredictions = append(autoCompletePredictions, domain.AutocompletePrediction{PlaceID: address.PlaceID, Description: address.Description})
	}
	return autoCompletePredictions, nil
}
