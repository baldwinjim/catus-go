package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/baldwinjim/catus-go/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var t models.User
var loginRequest models.UserLogin

type JwtToken struct {
	Token string `json:"token"`
}

// UserController for user controller
type UserController struct {
	userCollection *mongo.Collection
}

// NewUserController returns a new user controller
func NewUserController(c *mongo.Client, database string, collection string) *UserController {
	col := c.Database(database).Collection(collection)
	return &UserController{col}
}

// LoginUser allows the user to login
func (uc UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&loginRequest)
	err := uc.userCollection.FindOne(context.TODO(), bson.D{{"email", loginRequest.Email}}).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if loginRequest.Password != t.Password {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    t.Email,
		"password": t.Password,
		"org": t.Org,
		"role": t.Role,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

// GetUsers returns all users
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var t models.User
	//collection := uc.client.Database("gotest").Collection("users")
	err := uc.userCollection.FindOne(context.TODO(), bson.D{{"email", "jim.baldwin@gmail.com"}}).Decode(&t)
	if err != nil {
		//log.Fatal("Error:" err)
		io.WriteString(w, "No record found!")
	} else {
		err = json.NewEncoder(w).Encode(t)
	}

	//io.WriteString(w, "Get all users")
}

// func findByCredentials(email string, password string) {
// 	//collection := uc.Client.Database("gotest").Collection("users")
// 	err := collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&t)
// }

// AddUser adds a single user
func (uc UserController) AddUser(w http.ResponseWriter, r *http.Request) {
	//collection := uc.client.Database("gotest").Collection("users")
	decoder := json.NewDecoder(r.Body)
	var t models.User
	err := decoder.Decode(&t)
	if err != nil {
		io.WriteString(w, "Error")
	}
	res, err := uc.userCollection.InsertOne(context.Background(), t)
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
