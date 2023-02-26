package main

import (
	"fmt"
	"github.com/mcmuralishclint/location-service/api/http"
	"github.com/mcmuralishclint/location-service/app/location/service"
	"github.com/mcmuralishclint/location-service/config"
	"strconv"
)

var Config config.ServiceConfig

func init() {
	config.LoadConfigs()
	Config = config.Config
}

func main() {
	//repository := adapter.NewMockRepository()
	repository, _ := config.LoadLocationConfig()
	locationService := service.NewLocationService(repository)

	server := http.NewServer(locationService)
	port := Config.Port

	fmt.Println("Server listening on :" + strconv.Itoa(port))
	if err := server.Start(":" + strconv.Itoa(port)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
