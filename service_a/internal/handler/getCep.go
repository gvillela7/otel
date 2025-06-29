package handler

import (
	"encoding/json"
	"github.com/gvillela7/temperature/internal/data/response"
	"github.com/gvillela7/temperature/internal/services"
	"net/http"
	"strings"
)

type Cep struct {
	Cep string `json:"cep"`
}

func GetCep(w http.ResponseWriter, r *http.Request) {
	var cepDecode Cep
	err := json.NewDecoder(r.Body).Decode(&cepDecode)
	if err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "Unable to decode json.", nil)
		return
	}
	if cepDecode.Cep == "" {
		response.HttpResponse(w, http.StatusBadRequest, "Cep required.", nil)
		return
	}
	cep := strings.ReplaceAll(cepDecode.Cep, "-", "")
	if len(cep) != 8 {
		response.HttpResponse(w, http.StatusUnprocessableEntity, "invalid zipcode.", nil)
		return
	}

	service := services.NewTemperature()
	temperature, err := service.GetTemp(r.Context(), cep, w)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			response.HttpResponse(w, http.StatusNotFound, "Zipcode not found.", nil)
			return
		}
		response.HttpResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.HttpResponse(w, http.StatusOK, "success", temperature)
	return

}
