package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Attention! ", log.LstdFlags)
	srv := server.CreateServer(logger)
	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
