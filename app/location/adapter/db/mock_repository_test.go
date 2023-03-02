package db

import (
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMockRepo(t *testing.T) {
	repo := NewMockRepo("", "", 0)
	assert.Equal(t, repo, &mockRepo{})
}

func TestMockRepo_GetAddress(t *testing.T) {
	repo := NewMockRepo("", "", 0)

	// invalid address
	addr, err := repo.GetAddress("ADDR1")
	assert.NoError(t, err)
	assert.Nil(t, addr)

	// valid address
	addr, err = repo.GetAddress("CACHED_ADDRESS")
	assert.NoError(t, err)
	assert.Equal(t, addr, &domain.Address{Type: "TestType", FormattedAddress: "TestFormattedAddr"})
}

func TestRedisRepo_SetAddress(t *testing.T) {
	repo := NewMockRepo("", "", 0)
	address := domain.Address{Type: "TYPE", FormattedAddress: "FORMATTED_ADDRESS"}
	err := repo.SetAddress("ADDR1", &address)
	assert.NoError(t, err)
}
