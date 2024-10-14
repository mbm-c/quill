package main

import (
	"errors"
	"os"

	"tweety/internal/routes"

	"github.com/gofiber/fiber/v2/log"
)

// NewApp initializes the application and starts listening on the provided port
func NewApp(port string, listenFunc func(string) error) error {
	if port == "" {
		return errors.New("$PORT must be set")
	}

	err := listenFunc(":" + port)
	if err != nil {
		return err
	}

	return nil
}

// notest
func main() {
	port := os.Getenv("PORT")

	// Call NewApp with the actual app.Listen method
	err := NewApp(port, routes.Init().Listen)
	if err != nil {
		log.Fatal(err)
	}
}
