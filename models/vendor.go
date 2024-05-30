// models/vendor.go
package models

import "time"

// Vendor represents the schema for a vendor
type Vendor struct {
	ID                  string        `json:"id" bson:"_id,omitempty"`
	Username            string        `json:"username" bson:"username"`
	Password            string        `json:"password" bson:"password"`
	Profile             VendorProfile `json:"profile" bson:"profile"`
	LastUpdatedLocation time.Time     `json:"last_updated_location" bson:"last_updated_location"`
}

// VendorProfile represents the profile information of a vendor
type VendorProfile struct {
	ID       string    `json:"p_id" bson:"p_id,omitempty"`
	Name     string    `json:"name" bson:"name"`
	Email    string    `json:"email" bson:"email"`
	PhoneNo  string    `json:"phone_no" bson:"phone_no"`
	JoinedAt time.Time `json:"joined_at" bson:"joined_at"`
}
