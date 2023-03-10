package main

import (
	"fmt"
	"github.com/mcmuralishclint/location-service/api/http"
	"github.com/mcmuralishclint/location-service/app/location/service"
	"github.com/mcmuralishclint/location-service/config"
	"strconv"
)

var ConfigManager *config.ConfigManager

func init() {
	ConfigManager = config.NewConfigManager()
	ConfigManager.LoadConfigs()
}

// @title Location Service
// @description Location Service
// @version 1
// @host localhost:3000
// @BasePath /
func main() {
	cacheRepository, err := ConfigManager.LoadCacheConfig()
	if err != nil {
		panic(err)
	}
	locationRepository, err := ConfigManager.LoadLocationConfig(&cacheRepository)
	if err != nil {
		panic(err)
	}
	locationService := service.NewLocationService(locationRepository)

	server := http.NewServer(locationService)
	port := ConfigManager.Config.Port

	fmt.Println("Server listening on :" + strconv.Itoa(port))
	if err := server.Start(":" + strconv.Itoa(port)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
