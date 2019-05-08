package controllers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/baldwinjim/catus-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserController for user controller
type UserController struct {
	client *mongo.Client
}

// NewUserController returns a new user controller
func NewUserController(c *mongo.Client) *UserController {
	return &UserController{c}
}

// GetUsers returns all users
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var t models.User
	collection := uc.client.Database("gotest").Collection("users")
	err := collection.FindOne(context.TODO(), bson.D{{"email", "jim.baldwin@gmail"}}).Decode(&t)
	if err != nil {
		//log.Fatal("Error:" err)
		io.WriteString(w, "No record found!")
	} else {
		err = json.NewEncoder(w).Encode(t)
	}

	//io.WriteString(w, "Get all users")
}

// AddUser adds a single user
func (uc UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	collection := uc.client.Database("gotest").Collection("users")
	decoder := json.NewDecoder(r.Body)
	var t models.User
	err := decoder.Decode(&t)
	if err != nil {
		io.WriteString(w, "Error")
	}
	res, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(res)
	//io.WriteString(w, "Add user here")
}

// DeleteUser will delete a single user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Delete User here")
}
