package service

import (
	"fmt"
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type locationService struct {
	locationRepository domain.AddressRepository
	cacheRepository    domain.CacheRepository
}

func NewLocationService(locationRepository domain.AddressRepository, cacheRepository domain.CacheRepository) domain.LocationService {
	return &locationService{locationRepository: locationRepository, cacheRepository: cacheRepository}
}

func (s *locationService) GetAddressByID(id string) (*domain.Address, error) {
	address, _ := s.cacheRepository.GetAddress(id)
	if address != nil {
		fmt.Println("Receive the address from cache for id: ", id)
		return address, nil
	}
	address, _ = s.locationRepository.GetByID(id)
	err := s.cacheRepository.SetAddress(id, address)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (s *locationService) GetQueryAutoCompleteByText(input string) ([]domain.AutocompletePrediction, error) {
	return s.locationRepository.QueryAutoComplete(input)
}
