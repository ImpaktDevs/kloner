package server

import (
	"log"
	"strings"

	"main/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func StartServer() {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./public")

	app.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello world",
		})
	})

	app.Get("/auth", func(c *fiber.Ctx) error {
		return c.Render("auth", fiber.Map{
			"Title": "Access kloner",
		})
	})

	port := config.GetConfig().Port

	log.Fatal(app.Listen(strings.Join([]string{":", port}, "")))
}
