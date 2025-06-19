package services

import (
	"context"
	"net/http"
)

type TemperatureService interface {
	Fahrenheit(celsius float32) (float32, error)
	Celsius(ctx context.Context, cep string, w http.ResponseWriter) (*Temperature, error)
	Kelvin(celsius float32) (float32, error)
}
