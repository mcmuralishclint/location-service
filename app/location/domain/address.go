package domain

type Address struct {
	Type             string
	FormattedAddress string
}

type AutocompletePrediction struct {
	Description string `json:"description,omitempty"`
	PlaceID     string `json:"place_id,omitempty"`
}
