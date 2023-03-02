package third_party

import (
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGoogleMapsRepository(t *testing.T) {
	repo := NewMockRepository()
	if repo == nil {
		t.Fatalf("expected repo to not be nil")
	}
}

func TestMockRepository_GetByID(t *testing.T) {
	repo := NewMockRepository()

	id := "TEST_ID"
	addr, err := repo.GetByID(id)
	assert.Equal(t, err.Error(), "address not found")
	assert.Nil(t, addr)

	id = "GA123"
	addr, err = repo.GetByID(id)
	assert.NoError(t, err)
	assert.Equal(t, addr, &domain.Address{Type: "TestType", FormattedAddress: "123"})
}

func TestMockRepository_QueryAutoComplete(t *testing.T) {
	repo := NewMockRepository()

	input := "TEST_ID"
	addr, err := repo.QueryAutoComplete(input)
	assert.Equal(t, err.Error(), "No suggestions found")
	assert.Nil(t, addr)

	input = "test_addr"
	addr, err = repo.QueryAutoComplete(input)
	assert.NoError(t, err)
	expectedAddresses := []domain.AutocompletePrediction{
		{Description: "Address 1", PlaceID: "GA123"},
		{Description: "Address 2", PlaceID: "GA234"},
	}
	assert.Equal(t, addr, expectedAddresses)
}
