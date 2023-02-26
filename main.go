package main

import (
	"fmt"

	"github.com/mcmuralishclint/location-service/api/http"
	"github.com/mcmuralishclint/location-service/app/location/adapter"
	"github.com/mcmuralishclint/location-service/app/location/service"
)

func main() {
	repository := adapter.NewMockRepository()
	repository, _ = adapter.NewGoogleMapsRepository("RV5g8MbDtVVb4LVMZqKhHq2gV1cudjAB")
	locationService := service.NewLocationService(repository)

	server := http.NewServer(locationService)

	fmt.Println("Server listening on :8080")
	if err := server.Start(":8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
