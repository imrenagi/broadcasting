package server

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func (s *Server) routes() {
	// healthcheck
	s.Router.HandleFunc("/", s.healthcheckHandler)
	s.Router.HandleFunc("/healthz", s.healthcheckHandler)
	s.Router.HandleFunc("/readyz", s.readinessHandler)

	// serve api
	api := s.Router.PathPrefix("/api/v1/").Subrouter()
	api.Use(
		otelmux.Middleware(name),
	)
	api.Handle("/streams", otelhttp.WithRouteTag("/api/v1/streams",
		otelhttp.NewHandler(s.dummy(), name, otelhttp.WithHTTPRouteTag("/api/v1/streams"))))
}

func (s *Server) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("im alive"))
}

func (s *Server) readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("im ready to face the world"))
}

func (s *Server) dummy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}