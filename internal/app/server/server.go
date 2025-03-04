package server

import (
	"fmt"
	"net/http"

	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/api/handlers"
	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/config"
	"github.com/DmitriiSvarovskii/booking-service-go/internal/app/logger"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func ShortenerRouter(cfg *config.AppConfig) *Server {
	handler := handlers.NewHandler(cfg)

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return logger.RequestLogger(next.ServeHTTP)
	})

	fmt.Println("Setting up route for shortURL")

	r.Get("/", handler.BaseURL)
	r.MethodNotAllowed(handler.MethodNotAllowedHandle)

	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.ServiceURL,
			Handler: r,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
