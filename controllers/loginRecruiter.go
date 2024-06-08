package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	db "karya-auth/config"
	"karya-auth/models"
	"karya-auth/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// LoginRecruiter handles recruiter login
func LoginRecruiter(c *fiber.Ctx) error {

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	var recruiter models.Recruiter
	collection := db.GetCollection("recruiters")
	err := collection.FindOne(context.Background(), bson.M{"username": request.Username}).Decode(&recruiter)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(recruiter.Password), []byte(request.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	token, err := generateJWT(recruiter.Username, "recruiter")
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ok": false, "error": "Failed to generate token"})
	}

	return utils.JSONResponse(c, fiber.StatusOK, true, "Successfully Logged in!", fiber.Map{"token": token})
}

func generateJWT(username string, userType string) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"type":     userType,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 72 hours
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
