package server

import (
	"6sprintFinal/internal/handlers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

func NewServer(logger *log.Logger) *Server {
	r := chi.NewRouter()

	r.Get("/", handlers.MainHandle)
	r.Post("/upload", handlers.UploadHandle)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}

func (s *Server) Start() error {
	s.Logger.Println("Сервер запущен")
	return s.HTTP.ListenAndServe()
}
