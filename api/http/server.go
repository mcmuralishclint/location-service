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

	addressRouter := router.PathPrefix("/api/v1/address").Subrouter()
	addressRouter.HandleFunc("/validate", handler.GetAddressByID).Methods("GET")
	addressRouter.HandleFunc("/search", handler.GetAddressSuggestionsByText).Methods("GET")

	return &server{
		handler: router,
	}
}

func (s *server) Start(addr string) error {
	return http.ListenAndServe(addr, s.handler)
}
