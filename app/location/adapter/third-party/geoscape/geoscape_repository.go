package geoscape

import (
	"errors"
	"fmt"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/mcmuralishclint/location-service/util"
	"io"
	"net/http"
)

type GeoscapeRepository struct {
	country   string
	cacheRepo domain.CacheRepository
	ApiKey    string
}

func NewGeoscapeRepository(apiKey string, country string, cacheRepo domain.CacheRepository) (*GeoscapeRepository, error) {
	return &GeoscapeRepository{ApiKey: apiKey, country: country, cacheRepo: cacheRepo}, nil
}

func (r *GeoscapeRepository) GetByID(id string) (*domain.Address, error) {
	address, _ := r.cacheRepo.GetAddress(id)
	if address != nil {
		return nil, errors.New("Address not found")
	}
	// call API

	endpoint := fmt.Sprintf("https://api.psma.com.au/v1/predictive/address/%s", id)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.New("Address not found")
	}
	req.Header.Set("Authorization", r.ApiKey)
	req.Header.Set("Accept", "")
	req.Header.Set("Accept-Crs", "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("Address not found")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	addressFormatter := NewGeoscapeAddressFormatter()
	address, err = addressFormatter.FormatAddress(body)

	if address.Type == "" {
		return nil, errors.New("Address not found")
	}
	util.PopulateAddressCountry(address, r.country)

	// end call API
	r.cacheRepo.SetAddress(id, address)
	return address, nil
}

func (r *GeoscapeRepository) QueryAutoComplete(input string) ([]domain.AutocompletePrediction, error) {
	endpoint := fmt.Sprintf("https://api.psma.com.au/v1/predictive/address?query=%s", input)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.New("Keyword not found")
	}
	req.Header.Set("Authorization", r.ApiKey)
	req.Header.Set("Accept", "")
	req.Header.Set("Accept-Crs", "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.New("Keyword not found")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	addressFormatter := NewGeoscapeAddressFormatter()
	predictions, err := addressFormatter.FormatAddressSuggestion(body)

	return predictions, err
}
