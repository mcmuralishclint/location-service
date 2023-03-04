package domain

type Address struct {
	Type              string
	FormattedAddress  string
	AddressComponents interface{}
}

type AutocompletePrediction struct {
	Description string `json:"description,omitempty"`
	PlaceID     string `json:"place_id,omitempty"`
}

type AddressComponents struct {
	StreetNumber string `json:"street_number"`
	StreetName   string `json:"street_name"`
	City         string `json:"city"`
	Country      string `json:"country"`
	PostalCode   string `json:"postal_code"`
}
