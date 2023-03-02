package domain

type AddressRepository interface {
	GetByID(id string) (*Address, error)
	QueryAutoComplete(input string) ([]AutocompletePrediction, error)
}

type CacheRepository interface {
	GetAddress(string) (*Address, error)
	SetAddress(string, *Address) error
}
