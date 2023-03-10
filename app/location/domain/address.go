package domain

// Address represents a physical address.
// swagger:model
type Address struct {
	// The type of the address, such as "google" or "geoscape".
	Type string `json:"type"`
	// The formatted address string.
	FormattedAddress string `json:"formattedAddress"`
	// The components of the address, such as street number, street name, city, state, etc.
	AddressComponents interface{} `json:"addressComponents" ref:"#definitions/AddressComponents"`
}

// AutocompletePrediction represents a prediction response
// swagger:model
type AutocompletePrediction struct {
	// Formatted Address.
	Description string `json:"description,omitempty"`
	// address id.
	PlaceID string `json:"place_id,omitempty"`
}

// AddressComponents represents a physical address' fields.
// swagger:model
type AddressComponents struct {
	StreetNumber string `json:"street_number"`
	StreetName   string `json:"street_name"`
	City         string `json:"city"`
	Country      string `json:"country"`
	PostalCode   string `json:"postal_code"`
}
