package http

import (
	"encoding/json"
	"net/http"

	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type locationHandler struct {
	service domain.LocationService
}

func NewLocationHandler(service domain.LocationService) *locationHandler {
	return &locationHandler{service}
}

// GetAddressByID retrieves an address by ID.
// @Summary Get address by ID
// @Description Get an address by its ID.
// @Tags Addresses
// @ID get-address-by-id
// @Produce json
// @Param id query string true "Address ID"
// @Success 200 {object} domain.Address
// @Failure 404 {string} string "address not found"
// @Router /api/v1/address/validate [get]
func (h *locationHandler) GetAddressByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	address, err := h.service.GetAddressByID(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(address)
}

// GetAddressSuggestionsByText retrieves address suggestions by string.
// @Summary Get address suggestions by string
// @Description Get addresses by string.
// @Tags Addresses
// @ID get-addresses-by-string
// @Produce json
// @Param q query string true "Address ID"
// @Success 200 {array} domain.AutocompletePrediction
// @Failure 404 {string} string "Please input a valid string"
// @Router /api/v1/address/search [get]
func (h *locationHandler) GetAddressSuggestionsByText(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("q")
	addresses, err := h.service.GetQueryAutoCompleteByText(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(addresses)
}
