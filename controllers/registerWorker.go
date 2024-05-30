package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	db "karya-auth/config"
	"karya-auth/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterWorker handles the registration of a new worker
func RegisterWorker(c *fiber.Ctx) error {
	// Parse request body
	fmt.Println("Testtttttt")
	var worker models.Worker
	if err := c.BodyParser(&worker); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate fields
	if worker.Username == "" || worker.Password == "" || worker.Profile.Name == "" || worker.Profile.Email == "" || worker.Profile.PhoneNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(worker.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	worker.Password = string(hashedPassword)

	// Set last updated location
	worker.LastUpdatedLocation = time.Now()

	// Insert worker into database
	err = insertWorker(worker)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register worker"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Worker registered successfully"})
}

func insertWorker(worker models.Worker) error {
	collection := db.GetCollection("workers")
	_, err := collection.InsertOne(context.Background(), worker)
	if err != nil {
		log.Printf("Failed to insert worker into database: %v", err)
		return err
	}
	return nil
}
