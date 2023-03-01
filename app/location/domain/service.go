package domain

type LocationService interface {
	GetAddressByID(id string) (*Address, error)
	GetQueryAutoCompleteByText(input string) ([]AutocompletePrediction, error)
}
