package formatter

import (
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type AddressFormatter interface {
	FormatAddress(interface{}) (domain.AddressComponents, error)
	FormatAddressSuggestion(interface{}) ([]domain.AutocompletePrediction, error)
}
