package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/baldwinjim/catus-go/controllers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tpl *template.Template

var mongoClient *mongo.Client

func init() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected")
	}
	// Parse templates
	tpl = template.Must(template.ParseFiles("templates/login.gohtml"))
}

func getClient() *mongo.Client {
	//Get Database connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func logmein(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	uc := controllers.NewUserController(mongoClient, "gotest", "users")
	r.HandleFunc("/", logmein).Methods("GET")
	r.HandleFunc("/user", uc.AddUser).Methods("POST")
	r.HandleFunc("/user/login", uc.LoginUser).Methods("POST")
	r.HandleFunc("/user/register", uc.RegisterUser).Methods("POST")
	r.HandleFunc("/user/{id}", uc.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", uc.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
