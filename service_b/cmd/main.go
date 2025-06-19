package main

import (
	"context"
	config "github.com/gvillela7/temperature/service_b/configs"
	"github.com/gvillela7/temperature/service_b/internal/route"
	"github.com/gvillela7/temperature/service_b/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
	"log"
)

func InitTracer() func(context.Context) error {
	err := config.Load(".")
	cfg := config.GetZipkinConfig()
	if err != nil {
		util.Log(true, false, "error", "failed to initialize environment variables:", "error", err)
		panic(err)
	}
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("service-b"),
		semconv.ServiceVersionKey.String("1.0.0"),
		semconv.ServiceInstanceIDKey.String("b1"),
	)
	exporter, err := zipkin.New(
		cfg.Endpoint,
	)
	if err != nil {
		log.Fatalf("failed to create zipkin exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
func main() {
	shutdown := InitTracer()
	defer shutdown(context.Background())
	route.Run()
}
