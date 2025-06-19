package services

import (
	"context"
	"encoding/json"
	"errors"
	config "github.com/gvillela7/temperature/service_b/configs"
	"github.com/gvillela7/temperature/service_b/internal/data/response"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"io"
	"net/http"
	"strings"
)

type TemperatureCelsius struct {
	Temp string `json:"temp"`
}

type Temperature struct {
	State string  `json:"state"`
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}

type ViaCep struct {
	Estado string `json:"estado"`
	UF     string `json:"uf"`
	Erro   string `json:"erro,omitempty"`
}
type WeatherResponse struct {
	Current Current `json:"current"`
}
type Current struct {
	TempC float32 `json:"temp_c"`
}

func NewTemperature() *Temperature {
	return &Temperature{
		State: "",
		TempC: 0.0,
		TempF: 0.0,
		TempK: 0.0,
	}
}

var tracer = otel.Tracer("service-b")

func (t *Temperature) Celsius(ctx context.Context, cep string, w http.ResponseWriter) (*Temperature, error) {
	cfg := config.GetWeatherAPI()
	ctx, span := tracer.Start(ctx, "GetViaCEP")
	defer span.End()

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "error creating request for viacep.", nil)
		return nil, err
	}
	span.SetAttributes(attribute.String("cep", cep))
	res, err := client.Do(req)
	if err != nil {
		response.HttpResponse(w, http.StatusBadRequest, "request error.", nil)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "error read response viacep", nil)
		return nil, err
	}

	var viacep ViaCep
	if err := json.Unmarshal(body, &viacep); err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "error Unmarchal json viacep", nil)
		return nil, err
	}
	if viacep.Erro == "true" {
		response.HttpResponse(w, http.StatusNotFound, "zipcode not found.", nil)
		return nil, errors.New("zipcode not found")
	}
	if viacep.UF == "SP" {
		viacep.Estado = "Sao_Paulo"
	}

	state := strings.ReplaceAll(viacep.Estado, " ", "+")

	ctx, span = tracer.Start(ctx, "GetWeatherAPI")
	defer span.End()

	span.SetAttributes(attribute.String("estado", viacep.Estado), attribute.String("service.name", "WeatherAPI"))
	reqWeather, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://api.weatherapi.com/v1/current.json?key="+cfg.Key+"&q="+state+"&aqi=no",
		nil,
	)

	if err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "erro ao criar requisição para o weatherapi", nil)
		return nil, err
	}

	resWeather, err := client.Do(reqWeather)
	if err != nil {
		response.HttpResponse(w, http.StatusNotFound, "não foi possível encontrar informações meteorológicas.", nil)
		return nil, err
	}
	defer resWeather.Body.Close()

	bodyWeather, err := io.ReadAll(resWeather.Body)
	if err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "erro ao ler resposta do weatherapi", nil)
		return nil, err
	}

	var weather WeatherResponse
	if err := json.Unmarshal(bodyWeather, &weather); err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "erro ao decodificar resposta do weatherapi", nil)
		return nil, err
	}

	//tempC, _ := strconv.ParseFloat(weather.Current.TempC, 32)
	t.State = state
	t.TempC = weather.Current.TempC
	t.TempF, _ = t.Fahrenheit(t.TempC)
	t.TempK, _ = t.Kelvin(t.TempC)

	return t, nil
}

func (t *Temperature) Fahrenheit(celsius float32) (float32, error) {
	return celsius*1.8 + 32, nil
}

func (t *Temperature) Kelvin(celsius float32) (float32, error) {
	return celsius + 273.15, nil
}
