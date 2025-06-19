package services

import (
	"context"
	"net/http"
)

type TemperatureService interface {
	GetTemp(ctx context.Context, cep string, w http.ResponseWriter) (*Temperature, error)
}
