package db

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewMockRedisRepo(t *testing.T) {
	db, mock := redismock.NewClientMock()
	assert.NotNil(t, db, nil)
	assert.NotNil(t, mock, nil)
}

func TestRedisRepo_GetAddress(t *testing.T) {
	client, mockRedis := redismock.NewClientMock()
	key := "TEST_KEY"
	redisRepo := &redisRepo{client: client}
	defer client.Close()
	sampleVal := `{
        "type": "test",
        "formattedAddress": "TEST_FORMATTED_ADDRESS",
        "addressComponents": {
            "streetNumber": "TEST_STREET",
            "route": "TEST_ROUTE",
            "locality": "TEST_LOCALITY",
            "administrativeAreaLevel1": "TEST",
            "postalCode": "TEST_ZIP",
            "country": "TEST_COUNTRY"
        }
    }`

	// Key present in redis
	mockRedis.ExpectGet(key).SetVal(sampleVal)
	address, err := redisRepo.GetAddress(key)
	assert.Nil(t, err)
	expectedAddress := domain.Address{}
	json.Unmarshal([]byte(sampleVal), &expectedAddress)
	assert.Equal(t, address, &expectedAddress)

	// Key not present in redis
	mockRedis.ExpectGet(key).RedisNil()
	address, err = redisRepo.GetAddress(key)
	assert.Equal(t, address, (*domain.Address)(nil))
}

func TestMockRepo_SetAddress(t *testing.T) {
	sampleVal := `{
        "type": "test",
        "formattedAddress": "TEST_FORMATTED_ADDRESS",
        "addressComponents": {
            "streetNumber": "TEST_STREET",
            "route": "TEST_ROUTE",
            "locality": "TEST_LOCALITY",
            "administrativeAreaLevel1": "TEST",
            "postalCode": "TEST_ZIP",
            "country": "TEST_COUNTRY"
        }
    }`
	address := domain.Address{}
	json.Unmarshal([]byte(sampleVal), &address)

	client, mockRedis := redismock.NewClientMock()
	key := "TEST_KEY"
	redisRepo := &redisRepo{client: client}
	defer client.Close()
	err := redisRepo.SetAddress(key, &address)
	assert.NotNil(t, err)

	mockRedis.ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))
	err = redisRepo.SetAddress(key, &address)
	assert.NotNil(t, err)
}
