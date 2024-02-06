package main

import (
	"log"
	"log/slog"

	"github.com/Drakmyth/feedgator/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("No .env file found")
	} else {
		slog.Info("Found .env file")
	}

	log.Fatal(server.ListenAndServe())
}
