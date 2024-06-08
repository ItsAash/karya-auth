package models

type Industry struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
