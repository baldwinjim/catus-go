package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	First    string             `json:"first" bson:"first"`
	Last     string             `json:"last" bson:"last"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
	Plan     int                `json:"plan" bson:"plan"`
	Org      string             `json:"org" bson:"org"`
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Verified bool				`json:"verified" bson:"verified"`
}

// UserLogin struct for login requests from UI
type UserLogin struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
