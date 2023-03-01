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

func (h *locationHandler) GetAddressByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	address, err := h.service.GetAddressByID(idParam)
	if err != nil {
		http.Error(w, "address not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(address)
}

func (h *locationHandler) GetAddressSuggestionsByText(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("q")
	addresses, err := h.service.GetQueryAutoCompleteByText(input)
	if err != nil {
		http.Error(w, "Please input a valid string", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(addresses)
}
