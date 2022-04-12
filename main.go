package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Cfg struct {
	Addr string
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Make a channel to listen for INTERRUPT/TERMINATE signal from user.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/", passValidation)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("status: api router started, host: %s", ":8000")
		serverErrors <- http.ListenAndServe(":8080", nil)
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Printf("shutdown, status, shutdown started, signal %s", sig)
		defer log.Printf("shutdown, status, shutdown complete, signal %s", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}

func passValidation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	pass := struct {
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&pass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	valido, err := validar(pass.Password)
	w.Write([]byte())
}

func validar(pass string) bool {
	return false
}
