package config

import (
	"errors"
	"fmt"
	"github.com/mcmuralishclint/location-service/app/location/adapter/db"
	"github.com/mcmuralishclint/location-service/app/location/adapter/third-party"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"gopkg.in/yaml.v2"
	"os"
)

var Config ServiceConfig

type ServiceConfig struct {
	Port   int `yaml:"port"`
	Google struct {
		MapsApiKey string `yaml:"maps_api_key"`
	} `yaml:"google"`
	AddressProvider string      `yaml:"address_provider"`
	Cache           cacheConfig `yaml:"cache"`
}

type cacheConfig struct {
	Redis   redisConfig `yaml:"redis"`
	CacheDB string      `yaml:"cache_db"`
}

type redisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func LoadConfigs() {
	// Read the contents of the YAML file
	data, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
}
func LoadLocationConfig() (domain.AddressRepository, error) {
	switch Config.AddressProvider {
	case "google":
		fmt.Println("Using Google Configs")
		return third_party.NewGoogleMapsRepository(Config.Google.MapsApiKey)
	case "test":
		fmt.Println("Using Mock Configs")
		return third_party.NewMockRepository(), nil
	}
	return nil, errors.New("Wrong Config")
}

func LoadCacheConfig() (domain.CacheRepository, error) {
	switch Config.Cache.CacheDB {
	case "redis":
		fmt.Println("Redis DB was chosen as the cache layer")
		return db.NewRedisRepo(Config.Cache.Redis.Host, Config.Cache.Redis.Password, Config.Cache.Redis.DB), nil
	case "test":
		fmt.Println("Redis DB was chosen as the cache layer")
		return db.NewMockRepo("", "", 0), nil
	}
	return db.NewMockRepo("", "", 0), nil
}
