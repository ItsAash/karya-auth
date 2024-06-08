package controllers

import (
	"context"
	"fmt"

	db "karya-auth/config"
	"karya-auth/models"
	"karya-auth/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// LoginVendor handles vendor login
func LoginVendor(c *fiber.Ctx) error {

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	var vendor models.Vendor
	collection := db.GetCollection("vendors")
	err := collection.FindOne(context.Background(), bson.M{"username": request.Username}).Decode(&vendor)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(vendor.Password), []byte(request.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	token, err := generateJWT(vendor.Username, "vendor")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return utils.JSONResponse(c, fiber.StatusOK, true, "Successfully Logged in!", fiber.Map{"token": token})
}
