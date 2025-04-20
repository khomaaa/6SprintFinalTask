package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "INFO: ", log.LstdFlags)

	server := server.NewServer(logger)

	logger.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		logger.Fatalf("Could not start server: %s\n", err)
	}
}
