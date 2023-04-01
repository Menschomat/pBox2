package api

import (
	"log"
	"net/http"

	m "github.com/Menschomat/pBox2/model"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Route defines a valid endpoint with the type of action supported on it
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func NewBasicRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(GetCors())
	r.Use(middleware.Logger)
	return r
}

// NewRouter returns a router handle loaded with all the supported routes
func NewApiRouter(cfg *m.Configuration, mqttClient mqtt.Client) *chi.Mux {
	r := NewBasicRouter()
	var routes = []Route{
		{
			Method:      "GET",
			Path:        "/enclosure",
			HandlerFunc: GetEnclosure(cfg),
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/lights/{lightId}",
			HandlerFunc: GetLight(cfg),
		},
		{
			Method:      "POST",
			Path:        "/{boxId}/lights/{lightId}",
			HandlerFunc: UpdateLight(cfg, mqttClient),
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/fans/{fanId}",
			HandlerFunc: GetFan(cfg),
		},
		{
			Method:      "POST",
			Path:        "/{boxId}/fans/{fanId}",
			HandlerFunc: UpdateFan(cfg, mqttClient),
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/sensors/{sensorId}",
			HandlerFunc: GetSensor(cfg),
		},
		{
			Method:      "GET",
			Path:        "/{boxId}/sensors/{sensorId}/data",
			HandlerFunc: GetSensorData(cfg),
		},
	}

	for _, route := range routes {
		r.Method(route.Method, route.Path, route.HandlerFunc)
		log.Printf("Route added: %#v\n", route)
	}
	return r
}

func GetCors() func(http.Handler) http.Handler {
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}).Handler
	return handler
}
