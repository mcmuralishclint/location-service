package service

import (
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type locationService struct {
	locationRepository domain.AddressRepository
}

func NewLocationService(locationRepository domain.AddressRepository) domain.LocationService {
	return &locationService{locationRepository: locationRepository}
}

func (s *locationService) GetAddressByID(id string) (*domain.Address, error) {
	address, _ := s.locationRepository.GetByID(id)
	return address, nil
}

func (s *locationService) GetQueryAutoCompleteByText(input string) ([]domain.AutocompletePrediction, error) {
	return s.locationRepository.QueryAutoComplete(input)
}
