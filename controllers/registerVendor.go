package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	db "karya-auth/config"
	"karya-auth/models"
	"karya-auth/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// RegisterVendor handles the registration of a new vendor
func RegisterVendor(c *fiber.Ctx) error {
	// Parse request body
	var vendor models.Vendor
	if err := c.BodyParser(&vendor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validate fields
	if vendor.Username == "" || vendor.Password == "" || vendor.Profile.Name.FirstName == "" || vendor.Profile.Name.LastName == "" || vendor.Profile.Email == "" || vendor.Profile.PhoneNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(vendor.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	vendor.Password = string(hashedPassword)

	// Set last updated location
	vendor.LastUpdatedLocation = time.Now()

	// Insert vendor into database
	err = insertVendor(vendor)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return utils.JSONResponse(c, http.StatusBadRequest, false, "Username already exists", nil)
		}
		return utils.JSONResponse(c, http.StatusInternalServerError, false, "Failed to register vendor", nil)
	}

	token, err := generateJWT(vendor.Username, "vendor")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}
	return c.JSON(fiber.Map{"token": token})

}

func insertVendor(vendor models.Vendor) error {
	collection := db.GetCollection("vendors")
	_, err := collection.InsertOne(context.Background(), vendor)
	if err != nil {
		return err
	}
	return nil
}
