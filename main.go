package main

import (
	"os"

	"tweety/internal/routes"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := routes.Init()

	app.Listen(":" + port)
}
