package geoscape

import (
	"encoding/json"
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type GeoscapeAddressFormatter struct{}

func NewGeoscapeAddressFormatter() *GeoscapeAddressFormatter {
	return &GeoscapeAddressFormatter{}
}

func (f *GeoscapeAddressFormatter) FormatAddress(data interface{}) (*domain.AddressComponents, error) {
	return nil, nil
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
