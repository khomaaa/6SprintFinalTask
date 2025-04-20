package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	http.Server
}

func NewServer(logger *log.Logger) *Server {

	router := http.NewServeMux()

	// Register the root handler to serve index.html.
	router.HandleFunc("GET /", handlers.RootHandler)
	router.HandleFunc("POST /upload", handlers.UploadHandler)

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
