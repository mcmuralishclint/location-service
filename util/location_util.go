package util

import (
	"encoding/json"
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

func PopulateAddressCountry(address *domain.Address, country string) *domain.Address {
	addressComponents := domain.AddressComponents{}
	addressComponentsMarshal, _ := json.Marshal(address.AddressComponents)
	json.Unmarshal(addressComponentsMarshal, &addressComponents)
	addressComponents.Country = country
	address.AddressComponents = addressComponents
	return address
}
