package http

import (
	"github.com/mcmuralishclint/location-service/api/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	_ "github.com/mcmuralishclint/location-service/docs"
)

type server struct {
	handler http.Handler
}

func NewServer(service domain.LocationService) *server {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	handler := NewLocationHandler(service)

	// address routes
	addressRouter := router.PathPrefix("/api/v1/address").Subrouter()
	addressRouter.HandleFunc("/validate", handler.GetAddressByID).Methods("GET")
	addressRouter.HandleFunc("/search", handler.GetAddressSuggestionsByText).Methods("GET")

	// swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// server
	return &server{
		handler: router,
	}
}

func (s *server) Start(addr string) error {
	return http.ListenAndServe(addr, s.handler)
}
