package db

import (
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInMemMockRepo(t *testing.T) {
	repo := NewInMemRepo()
	assert.NotNil(t, repo)
}

func TestNewInMemRepo_GetAddressWithoutSetting(t *testing.T) {
	repo := NewInMemRepo()
	address, error := repo.GetAddress("TEST")
	assert.Nil(t, error)
	assert.Nil(t, address)
}

func TestNewInMemRepo_GetAddressWithSetting(t *testing.T) {
	repo := NewInMemRepo()
	address := domain.Address{FormattedAddress: "TEST", Type: "TEST", AddressComponents: domain.AddressComponents{}}
	key := "address_id"
	repo.SetAddress(key, &address)
	addressFromGet, err := repo.GetAddress(key)
	assert.Nil(t, err)
	address.AddressComponents = map[string]interface{}{"city": "", "country": "", "postal_code": "", "street_name": "", "street_number": ""}
	assert.Equal(t, *addressFromGet, address)
}

func TestNewInMemRepo_SetAddress(t *testing.T) {
	repo := NewInMemRepo()
	address := domain.Address{FormattedAddress: "TEST", Type: "TEST", AddressComponents: domain.AddressComponents{}}
	key := "address_id"
	err := repo.SetAddress(key, &address)
	assert.Nil(t, err)
}
