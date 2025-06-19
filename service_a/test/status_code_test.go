package test

import (
	"github.com/gvillela7/temperature/configs"
	handler2 "github.com/gvillela7/temperature/internal/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus200(t *testing.T) {
	err := config.Load("../")
	assert.NoError(t, err, "Load config fail")
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/v1/temperature?cep=79801905", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler2.GetCep)
	handler.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestStatus422(t *testing.T) {
	err := config.Load("../")
	assert.NoError(t, err, "Load config fail")
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/v1/temperature?cep=79801", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler2.GetCep)
	handler.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	jsonString := rr.Body.String()
	assert.Contains(t, jsonString, "invalid zipcode.", "Response should contain the expected message")
}

func TestStatus404(t *testing.T) {
	err := config.Load("../")
	assert.NoError(t, err, "Load config fail")
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8000/v1/temperature?cep=11111111", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler2.GetCep)
	handler.ServeHTTP(rr, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	jsonString := rr.Body.String()
	assert.Contains(t, jsonString, "zipcode not found.", "Response should contain the expected message")
}
