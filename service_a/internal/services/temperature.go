package services

import (
	"context"
	"encoding/json"
	"github.com/gvillela7/temperature/internal/data/response"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type TemperatureCelsius struct {
	Temp string `json:"temp"`
}
type TemperatureResponse struct {
	Data Data `json:"data"`
}
type Data struct {
	State string  `json:"state"`
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}
type Temperature struct {
	State string  `json:"state"`
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
	TempK float32 `json:"temp_k"`
}

func NewTemperature() *Temperature {
	return &Temperature{
		State: "",
		TempC: 0.0,
		TempF: 0.0,
		TempK: 0.0,
	}
}

var tracer = otel.Tracer("service-a")

func (t *Temperature) GetTemp(ctx context.Context, cep string, w http.ResponseWriter) (*Temperature, error) {
	ctx, span := tracer.Start(ctx, "callServiceB")
	defer span.End()

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://otel-service_b-1:8001/v1/temperature?cep="+cep, nil)
	if err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "error creating request for viacep.", nil)
		return nil, err
	}
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
	var temperatureResponse TemperatureResponse
	if err := json.Unmarshal(body, &temperatureResponse); err != nil {
		response.HttpResponse(w, http.StatusInternalServerError, "error Unmarchal json viacep", nil)
		return nil, err
	}

	t.State = temperatureResponse.Data.State
	t.TempC = temperatureResponse.Data.TempC
	t.TempF = temperatureResponse.Data.TempF
	t.TempK = temperatureResponse.Data.TempK

	return t, nil
}
