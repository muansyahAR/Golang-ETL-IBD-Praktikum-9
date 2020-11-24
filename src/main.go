package main

import (
	s "src/controllers"

	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Berhasil")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", hello)
	// app.Get("/ambil", s.Ambil)
	app.Post("/uploadFile", s.UploadFile)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
