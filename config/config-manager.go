package config

import (
	"errors"
	"fmt"
	"github.com/mcmuralishclint/location-service/app/location/adapter/db"
	"github.com/mcmuralishclint/location-service/app/location/adapter/third-party"
	"github.com/mcmuralishclint/location-service/app/location/adapter/third-party/geoscape"
	"github.com/mcmuralishclint/location-service/app/location/adapter/third-party/google"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"gopkg.in/yaml.v2"
	"os"
)

type ServiceConfig struct {
	Country         string          `yaml:"country"`
	Port            int             `yaml:"port"`
	Cache           cacheConfig     `yaml:"cache"`
	LocationPartner locationPartner `yaml:"location_partner"`
}

type locationPartner struct {
	Google struct {
		MapsApiKey string `yaml:"maps_api_key"`
	} `yaml:"google"`
	Geoscape struct {
		MapsApiKey string `yaml:"maps_api_key"`
	} `yaml:"geoscape"`
	Test struct {
		MapsApiKey string `yaml:"maps_api_key"`
	} `yaml:"test"`
	LocationProvider string `yaml:"location_provider"`
}

type cacheConfig struct {
	Redis    redisConfig `yaml:"redis"`
	InMemory inMemConfig `yaml:"in_memory"`
	CacheDB  string      `yaml:"cache_db"`
}

type inMemConfig struct{}
type redisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type ConfigManager struct {
	Config ServiceConfig
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}

type ConfigInterface interface {
	LoadConfigs() *ServiceConfig
	LoadLocationConfig() (domain.AddressRepository, error)
	LoadCacheConfig() (domain.CacheRepository, error)
}

func (c *ConfigManager) LoadConfigs() {
	// Read the contents of the YAML file
	data, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &c.Config)
	if err != nil {
		panic(err)
	}
}
func (c *ConfigManager) LoadLocationConfig(cacheRepository *domain.CacheRepository) (domain.AddressRepository, error) {
	switch c.Config.LocationPartner.LocationProvider {
	case "google":
		fmt.Println("Using Google Configs")
		return google.NewGoogleMapsRepository(c.Config.LocationPartner.Google.MapsApiKey, c.Config.Country, *cacheRepository)
	case "geoscape":
		fmt.Println("Using Geoscape Configs")
		return geoscape.NewGeoscapeRepository(c.Config.LocationPartner.Geoscape.MapsApiKey, c.Config.Country, *cacheRepository)
	case "test":
		fmt.Println("Using Mock Configs")
		return third_party.NewMockRepository(), nil
	}
	return nil, errors.New("Wrong Config")
}

func (c *ConfigManager) LoadCacheConfig() (domain.CacheRepository, error) {
	switch c.Config.Cache.CacheDB {
	case "redis":
		fmt.Println("Redis DB was chosen as the cache layer")
		return db.NewRedisRepo(c.Config.Cache.Redis.Host, c.Config.Cache.Redis.Password, c.Config.Cache.Redis.DB), nil
	case "in_memory":
		fmt.Println("In Memory Cache was chosen as the cache layer")
		return db.NewInMemRepo(), nil
	case "test":
		fmt.Println("Redis DB was chosen as the cache layer")
		return db.NewMockRepo("", "", 0), nil
	}
	return db.NewMockRepo("", "", 0), nil
}
