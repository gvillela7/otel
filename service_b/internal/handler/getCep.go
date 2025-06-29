package handler

import (
	"github.com/gvillela7/temperature/service_b/internal/data/response"
	"github.com/gvillela7/temperature/service_b/internal/services"
	"net/http"
	"strings"
)

func GetCep(w http.ResponseWriter, r *http.Request) {
	cepRequest := r.URL.Query().Get("cep")
	cep := strings.ReplaceAll(cepRequest, "-", "")
	if len(cep) != 8 {
		response.HttpResponse(w, http.StatusUnprocessableEntity, "invalid zipcode.", nil)
		return
	}

	service := services.NewTemperature()
	temperature, err := service.Celsius(r.Context(), cep, w)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			response.HttpResponse(w, http.StatusNotFound, "Zipcode not found.", nil)
			return
		}
		response.HttpResponse(w, http.StatusInternalServerError, "error processing request.", nil)
		return
	}

	if err == nil {
		response.HttpResponse(w, http.StatusOK, "success", temperature)
		return
	}
}
