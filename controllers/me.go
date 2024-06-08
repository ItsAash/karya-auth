package controllers

import (
	"context"
	db "karya-auth/config"
	"karya-auth/models"
	"karya-auth/utils"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
)

// GetVendorProfile retrieves the profile of the authenticated vendor
func GetVendorProfile(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	userType := c.Locals("type").(string)

	if userType != "vendor" {
		return utils.JSONResponse(c, fiber.StatusUnauthorized, false, "Token not valid", nil)
	}

	var vendor models.Vendor
	collection := db.GetCollection("vendors")
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&vendor)
	if err != nil {
		return utils.JSONResponse(c, fiber.StatusInternalServerError, false, "Failed to retrieve vendor profile", nil)
	}

	vendorResponse := fiber.Map{
		"id":                    vendor.ID,
		"username":              vendor.Username,
		"profile":               vendor.Profile,
		"last_updated_location": vendor.LastUpdatedLocation,
	}

	return utils.JSONResponse(c, fiber.StatusOK, true, "Successfully retrieved vendor profile", fiber.Map{
		"user": vendorResponse,
	})
}

// GetRecruiterProfile retrieves the profile of the authenticated recruiter
func GetRecruiterProfile(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	userType := c.Locals("type").(string)

	if userType != "recruiter" {
		return utils.JSONResponse(c, fiber.StatusUnauthorized, false, "Token not valid", nil)
	}

	var recruiter models.Recruiter
	collection := db.GetCollection("recruiters")
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&recruiter)
	if err != nil {
		return utils.JSONResponse(c, fiber.StatusInternalServerError, false, "Failed to retrieve recruiter profile", nil)
	}

	recruiterResponse := fiber.Map{
		"id":                    recruiter.ID,
		"username":              recruiter.Username,
		"profile":               recruiter.Profile,
		"last_updated_location": recruiter.LastUpdatedLocation,
	}

	return utils.JSONResponse(c, fiber.StatusOK, true, "Successfully retrieved recruiter profile", fiber.Map{
		"user": recruiterResponse,
	})
}
