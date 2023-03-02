package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mcmuralishclint/location-service/api/http"
	"github.com/mcmuralishclint/location-service/app/location/adapter/db"
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
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	cacheRepository := db.NewRedisRepo(redisClient)
	locationRepository, err := config.LoadLocationConfig()
	if err != nil {
		panic(err)
	}
	locationService := service.NewLocationService(locationRepository, cacheRepository)

	server := http.NewServer(locationService)
	port := Config.Port

	fmt.Println("Server listening on :" + strconv.Itoa(port))
	if err := server.Start(":" + strconv.Itoa(port)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
