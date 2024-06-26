package main

import (
	"net/http"
	"os"

	"github.com/subbbbbaru/first-sample/internal/handlers"
	"github.com/subbbbbaru/first-sample/internal/repositories"
	"github.com/subbbbbaru/first-sample/internal/services"
	myLog "github.com/subbbbbaru/first-sample/pkg/log"
)

func main() {
	// Инициализация логгеров
	myLog.InitLoggers(os.Stdout, os.Stderr)

	repo := repositories.NewInMemoryItemRepository()
	service := services.NewItemService(repo)
	handler := handlers.NewItemHandler(service)

	http.HandleFunc("/items", handler.HandleItems)

	// log.Println("Server is starting on port 8080...")
	myLog.Info().Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		myLog.Error().Fatalf("could not start server: %v\n", err)
	}
}
