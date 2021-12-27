package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mertingen/go-samples/handlers"
	"github.com/mertingen/go-samples/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))
	databaseServ := services.NewDatabase("mysql", connInfo)
	db := databaseServ.Connect()

	// defer the close till after the main function has finished
	// executing
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Database close: %s\n", err)
		}
	}(db)

	studentService := services.NewStudent(db)
	studentHandler := handlers.NewStudent(studentService)

	r := mux.NewRouter()
	//specify endpoints, handler functions and HTTP method
	r.HandleFunc("/health", handlers.Health).Methods("GET")
	r.HandleFunc("/students", studentHandler.FetchAll).Methods("GET")
	r.HandleFunc("/students/{id}", studentHandler.FetchOne).Methods("GET")
	r.HandleFunc("/students", studentHandler.Insert).Methods("POST")
	r.HandleFunc("/students/{id}", studentHandler.Update).Methods("PUT")
	r.HandleFunc("/students/{id}", studentHandler.Delete).Methods("DELETE")
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("Server is up and running...")
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
