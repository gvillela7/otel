package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	config "github.com/gvillela7/temperature/service_b/configs"
	"github.com/gvillela7/temperature/service_b/internal/handler"
	"github.com/gvillela7/temperature/service_b/util"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Run() {
	cfg := config.GetAPIConfig()
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/temperature", otelhttp.NewHandler(http.HandlerFunc(handler.GetCep), "endpoint-service-b").ServeHTTP)
	})

	util.Log(true, false, "info", "Server running on", "Port", cfg.Port)
	http.ListenAndServe(cfg.Port, r)
}
