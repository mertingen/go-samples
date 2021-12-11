package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/mertingen/go-samples/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("Server is running...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C), SIGKILL, SIGQUIT.
	// SIGTERM (Ctrl+/) are not caught.
	signal.Notify(c, os.Interrupt, syscall.SIGKILL, syscall.SIGQUIT)

	// Block until we receive our signal.
	<-c
	log.Println("Shutdown request (Ctrl-C) caught")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Don't block if no connections, otherwise wait until the timeout deadline.
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Shutdown: %s\n", err)
	}
	// Optionally, could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Shutting down ...")
}
