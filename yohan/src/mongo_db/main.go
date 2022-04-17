package main

import (
	"fmt"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type Person struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client
func main() {
	fmt.Println("hello ex")
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	http>ListenAndServe(":12345", handler)
}
