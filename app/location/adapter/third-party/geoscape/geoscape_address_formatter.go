package geoscape

import (
	"encoding/json"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/mcmuralishclint/location-service/app/location/formatter"
	"log"
)

type GeoscapeAddressFormatter struct{}

func NewGeoscapeAddressFormatter() formatter.AddressFormatter {
	return &GeoscapeAddressFormatter{}
}

func (f *GeoscapeAddressFormatter) FormatAddress(data interface{}) (*domain.Address, error) {
	var response map[string]interface{}
	if err := json.Unmarshal(data.([]byte), &response); err != nil {
		log.Fatal(err)
	}
	address := domain.Address{}

	if response["address"] == nil {
		return &address, nil
	}
	propertiesResponse := response["address"].(map[string]interface{})["properties"]

	addressComponent := &domain.AddressComponents{}
	for key, component := range propertiesResponse.(map[string]interface{}) {
		switch key {
		case "street_number_1":
			addressComponent.StreetNumber = component.(string)
		case "street_name":
			addressComponent.StreetName = component.(string)
		case "locality_name":
			addressComponent.City = component.(string)
		case "postcode":
			addressComponent.PostalCode = component.(string)
		case "formatted_address":
			address.FormattedAddress = component.(string)
		}
	}
	address.AddressComponents = addressComponent
	address.Type = "Geoscape"
	return &address, nil
}

func (f *GeoscapeAddressFormatter) FormatAddressSuggestion(data interface{}) ([]domain.AutocompletePrediction, error) {
	type LocationSuggestions struct {
		Suggest []struct {
			Address string `json:"address"`
			ID      string `json:"id"`
		} `json:"suggest"`
	}
	var locationSugestions LocationSuggestions

	json.Unmarshal(data.([]byte), &locationSugestions)

	var autocompletePredictions []domain.AutocompletePrediction
	for _, prediction := range locationSugestions.Suggest {
		autocompletePredictions = append(autocompletePredictions, domain.AutocompletePrediction{
			Description: prediction.Address,
			PlaceID:     prediction.ID,
		})
	}
	return autocompletePredictions, nil
}
