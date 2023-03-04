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
	Subpremise          string `json:"subpremise"`
	StreetNumber        string `json:"street_number"`
	Route               string `json:"route"`
	Locality            string `json:"locality"`
	AdministrativeArea3 string `json:"administrative_area_level_3"`
	AdministrativeArea2 string `json:"administrative_area_level_2"`
	AdministrativeArea1 string `json:"administrative_area_level_1"`
	Country             string `json:"country"`
	PostalCode          string `json:"postal_code"`
}
