package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	myHTTP "github.com/mcmuralishclint/location-service/api/http"
	"github.com/mcmuralishclint/location-service/app/location/domain"
	"github.com/stretchr/testify/assert"
)

type mockService struct{}

func (s *mockService) GetAddressByID(id string) (*domain.Address, error) {
	if id == "123" {
		return &domain.Address{
			Type:             "GOOGLE",
			FormattedAddress: "TEST",
		}, nil
	}
	return nil, errors.New("ERROR")
}

func (s *mockService) GetQueryAutoCompleteByText(input string) ([]domain.AutocompletePrediction, error) {
	if input == "Lo" {
		return []domain.AutocompletePrediction{{Description: "Los Angeles, CA", PlaceID: "ID1"}}, nil
	}
	return nil, errors.New("ERROR")
}

func TestLocationHandler_GetRightAddressByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/address?id=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := myHTTP.NewLocationHandler(&mockService{})
	http.HandlerFunc(handler.GetAddressByID).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := &domain.Address{
		Type:             "GOOGLE",
		FormattedAddress: "TEST",
	}

	var actual domain.Address
	err = json.NewDecoder(rr.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, &actual)
}

func TestLocationHandler_GetWrongAddressByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/address?id=456", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := myHTTP.NewLocationHandler(&mockService{})
	http.HandlerFunc(handler.GetAddressByID).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var actual domain.Address
	err = json.NewDecoder(rr.Body).Decode(&actual)

	assert.Error(t, err)
	assert.Equal(t, actual, domain.Address{})
}

func TestLocationHandler_GetRightAddressSuggestionsByText(t *testing.T) {
	req, err := http.NewRequest("GET", "/suggestions?q=Lo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := myHTTP.NewLocationHandler(&mockService{})
	http.HandlerFunc(handler.GetAddressSuggestionsByText).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []domain.AutocompletePrediction{{"Los Angeles, CA", "ID1"}}

	var actual []domain.AutocompletePrediction
	err = json.NewDecoder(rr.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestLocationHandler_GetWrongAddressSuggestionsByText(t *testing.T) {
	req, err := http.NewRequest("GET", "/suggestions?q=Loss", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := myHTTP.NewLocationHandler(&mockService{})
	http.HandlerFunc(handler.GetAddressSuggestionsByText).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var actual []domain.AutocompletePrediction
	err = json.NewDecoder(rr.Body).Decode(&actual)
	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestNewLocationHandler(t *testing.T) {
	service := mockService{}
	handler := myHTTP.NewLocationHandler(&service)
	assert.NotNil(t, handler)
}
