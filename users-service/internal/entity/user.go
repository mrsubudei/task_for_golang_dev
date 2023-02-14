// Package entity defines main entities for business logic (services)
// HTTP response objects if suitable.
package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// User -.
type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	Salt     string             `json:"salt" bson:"salt"`
	Password string             `json:"password" bson:"password"`
}
