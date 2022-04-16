package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/filipeandrade6/iti-digital-desafio-backend/app/handlers"
)

// var build = "develop"
var porta = ":8000"
var serverShutdownTime = 5 * time.Second

func main() {
	logger := log.New(os.Stdout, "PASSVALIDATION", 0)

	if err := run(logger); err != nil {
		logger.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	// ========== GOMAXPROCS
	// ========== CONFIGURATION
	// ========== APP STARTING
	// ========== AUTH
	// ========== DB
	// ========== TRACING
	// ========== DEBUG
	// ========== API SERVICE
	logger.Print("startup", "status", "initializing V1 API support")

	// Make a channel to listen for INTERRUPT/TERMINATE signal from user.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Construct the API to service the requests.
	api := handlers.NewAPI()

	// Construct the server to the service the rpi
	srv := &http.Server{
		Addr:    porta,
		Handler: api,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)
	go func() {
		logger.Printf("status: api router started, host: %s", porta) // * puxar das configurações
		serverErrors <- srv.ListenAndServe()
	}()

	// ========== SHUTDOWN

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		logger.Printf("shutdown, status, shutdown started, signal %s", sig)
		defer logger.Printf("shutdown, status, shutdown complete, signal %s", sig)

		ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTime) // * puxar das configurações
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
