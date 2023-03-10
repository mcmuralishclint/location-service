package service

import (
	"errors"
	"fmt"
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type locationService struct {
	locationRepository domain.AddressRepository
}

func NewLocationService(locationRepository domain.AddressRepository) domain.LocationService {
	return &locationService{locationRepository: locationRepository}
}

func (s *locationService) GetAddressByID(id string) (*domain.Address, error) {
	address, err := s.locationRepository.GetByID(id)
	return address, err
}

func (s *locationService) GetQueryAutoCompleteByText(input string) ([]domain.AutocompletePrediction, error) {
	addresses, err := s.locationRepository.QueryAutoComplete(input)
	if len(addresses) == 0 {
		err = errors.New(fmt.Sprintf("Unable to find addresses for input %s", input))
	}
	return addresses, err
}
