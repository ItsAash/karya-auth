package main

import (
	db "karya-auth/config"
	"karya-auth/controllers"

	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	db.Connect()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World"})
	})

	app.Post("/worker/register", controllers.RegisterWorker)
	app.Post("/vendor/register", controllers.RegisterVendor)

	log.Fatal(app.Listen(":8080"))
}
