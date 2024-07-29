package main

import (
	"context"
	"fmt"
	"go-api-starter/internal/api"
	"os"
	"os/signal"

	_ "go-api-starter/docs"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo and gorm.
// @host localhost:3000
// @BasePath /
func main() {
	// Initialize the env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err)
	}

	if enviro := os.Getenv("ENVIRONMENT"); enviro != "production" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	server := api.New(fmt.Sprintf(":%s", os.Getenv("PORT")))

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	// gracefully shutdown server
	if err := server.Stop(); err != nil {
		log.Fatal().Err(err)
	}
}
