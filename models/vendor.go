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
	ProfileID   string     `json:"p_id" bson:"p_id"`
	Name        Name       `json:"name" bson:"name"`
	Age         int        `json:"age" bson:"age"`
	Sex         string     `json:"sex" bson:"sex"`
	Email       string     `json:"email" bson:"email"`
	PhoneNumber string     `json:"phone_no" bson:"phone_no"`
	JoinedAt    time.Time  `json:"joined_at" bson:"joined_at"`
	Rating      float32    `json:"rating" bson:"rating"`
	Industries  []Industry `json:"industries" bson:"industries"`
}

type Name struct {
	FirstName  string `json:"first_name" bson:"first_name"`
	MiddleName string `json:"middle_name" bson:"middle_name"`
	LastName   string `json:"last_name" bson:"last_name"`
}
