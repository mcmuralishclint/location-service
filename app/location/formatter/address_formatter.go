package formatter

import (
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type AddressFormatter interface {
	FormatAddress(interface{}) (*domain.Address, error)
	FormatAddressSuggestion(interface{}) ([]domain.AutocompletePrediction, error)
}
