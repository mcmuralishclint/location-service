package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcmuralishclint/location-service/app/location/domain"
)

type server struct {
	handler http.Handler
}

func NewServer(service domain.LocationService) *server {
	router := mux.NewRouter()

	handler := NewLocationHandler(service)
	router.HandleFunc("/address/validate", handler.GetAddressByID).Methods("GET")
	router.HandleFunc("/address/search", handler.GetAddressSuggestionsByText).Methods("GET")

	return &server{
		handler: router,
	}
}

func (s *server) Start(addr string) error {
	return http.ListenAndServe(addr, s.handler)
}
