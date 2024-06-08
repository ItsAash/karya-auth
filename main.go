package main

import (
	db "karya-auth/config"
	"karya-auth/controllers"
	"karya-auth/middlewares"

	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	db.Connect()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World"})
	})

	app.Post("/recruiter/register", controllers.RegisterRecruiter)
	app.Post("/vendor/register", controllers.RegisterVendor)
	app.Post("/vendor/login", controllers.LoginVendor)
	app.Post("/recruiter/login", controllers.LoginRecruiter)

	protected := app.Group("/me", middlewares.JWTProtected())
	protected.Get("/vendor", controllers.GetVendorProfile)
	protected.Get("/recruiter", controllers.GetRecruiterProfile)

	log.Fatal(app.Listen(":8080"))
}
