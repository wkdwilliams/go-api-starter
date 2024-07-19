package main

import (
	"context"
	"fmt"
	"go-api-starter/api"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize the env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	server := api.New(fmt.Sprintf(":%s", os.Getenv("PORT")))

	go server.Start()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()

	// gracefully shutdown server
	if err := server.Stop(); err != nil {
		server.Echo.Logger.Fatal(err)
	}
}
