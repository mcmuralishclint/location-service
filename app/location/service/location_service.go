package service

import "github.com/mcmuralishclint/location-service/app/location/domain"

type locationService struct {
	repository domain.AddressRepository
}

func NewLocationService(repository domain.AddressRepository) domain.LocationService {
	return &locationService{repository: repository}
}

func (s *locationService) GetAddressByID(id string) (*domain.Address, error) {
	return s.repository.GetByID(id)
}
