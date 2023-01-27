package server

import (
	"log"
	"strings"

	"main/config"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	app.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello world",
		})
	})

	log.Fatal(app.Listen(strings.Join([]string{":", config.GetConfig().Port}, "")))
}
