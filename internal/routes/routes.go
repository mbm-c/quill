package routes

import (
	"fmt"
	"tweety/internal/ws"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	app.Use(recover.New(recover.Config{EnableStackTrace: true, StackTraceHandler: defaultStackTraceHandler}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use("/ws", ws.Get)
	app.Get("/ws/:id", websocket.New(ws.GetByID))

	return app
}

func defaultStackTraceHandler(c *fiber.Ctx, err interface{}) {
	fmt.Println(err.(error).Error())
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	c.Status(fiber.StatusInternalServerError)

	c.JSON(fiber.Map{
		"error": err.(error).Error(),
	})
}
