package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mertingen/go-samples/config"
	"github.com/mertingen/go-samples/handlers"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := fiber.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	err := config.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	app.Get("/students", handlers.GetStudents)
	app.Get("/students/:id", handlers.GetStudent)
	app.Post("/students", handlers.AddStudent)
	app.Put("/students/:id", handlers.UpdateStudent)
	app.Delete("/students/:id", handlers.RemoveStudent)
	app.Post("/students/:id/lectures", handlers.AttachLectures)
	app.Get("/lectures", handlers.GetLectures)
	app.Get("/lectures/:id", handlers.GetLecture)
	app.Post("/lectures", handlers.AddLecture)
	app.Put("/lectures/:id", handlers.UpdateLecture)
	app.Delete("/lectures/:id", handlers.RemoveLecture)

	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Panic(err)
	}
}
