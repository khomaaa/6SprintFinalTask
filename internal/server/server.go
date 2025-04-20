package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger
	http.Server
}

func NewServer(logger *log.Logger) *Server {

	router := chi.NewRouter()
	router.Get("/", handlers.RootHandler)
	router.Post("/upload", handlers.UploadHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

	// Настройка HTTP-сервера
	srv := &Server{
		Logger: logger,
		Server: http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	return srv
}
