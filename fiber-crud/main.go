package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mertingen/go-samples/config"
	"github.com/mertingen/go-samples/handlers"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	err := config.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	app.Get("/students", handlers.GetStudents)
	app.Get("/students/:id", handlers.GetStudent)
	app.Post("/students", handlers.AddStudent)
	app.Put("/students/:id", handlers.UpdateStudent)
	app.Delete("/students/:id", handlers.RemoveStudent)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))

}
