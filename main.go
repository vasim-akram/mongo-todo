package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vasim-akram/mongo-todo/database"
	"github.com/vasim-akram/mongo-todo/todo"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	mongoDB := database.Connect()

	api := app.Group("/api")
	todo.Register(api, mongoDB)

	log.Fatal(app.Listen(":5000"))
}
