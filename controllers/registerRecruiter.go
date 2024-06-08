package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	db "karya-auth/config"
	"karya-auth/models"
	"karya-auth/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRecruiter handles the registration of a new recruiter
func RegisterRecruiter(c *fiber.Ctx) error {
	// Parse request body
	var recruiter models.Recruiter
	if err := c.BodyParser(&recruiter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate fields
	if recruiter.Username == "" || recruiter.Password == "" || recruiter.Profile.Name == "" || recruiter.Profile.Email == "" || recruiter.Profile.PhoneNo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(recruiter.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	recruiter.Password = string(hashedPassword)

	// Set last updated location
	recruiter.LastUpdatedLocation = time.Now()

	// Insert recruiter into database
	err = insertRecruiter(recruiter)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return utils.JSONResponse(c, http.StatusBadRequest, false, "Username already exists", nil)
		}
		return utils.JSONResponse(c, http.StatusInternalServerError, false, "Failed to register vendor", nil)
	}

	token, err := generateJWT(recruiter.Username, "recruiter")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})

}

func insertRecruiter(recruiter models.Recruiter) error {
	collection := db.GetCollection("recruiters")
	_, err := collection.InsertOne(context.Background(), recruiter)
	if err != nil {
		log.Printf("Failed to insert worker into database: %v", err)
		return err
	}
	return nil
}
