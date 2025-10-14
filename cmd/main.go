package main

import (
	"log"
	"os"

	"github.com/Oomat-Dzhumagulov/6SprintFinal/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.NewServer(logger)

	if err := srv.Start(); err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
