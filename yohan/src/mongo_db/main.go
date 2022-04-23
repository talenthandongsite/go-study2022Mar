package main

import (
	"fmt"
	"time"
	"log"
	"encoding/json"

	"context"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	

)

type Person struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	// response.Header().Add("api-key", "kvfMNeHnbZrcEgRO7HlFtqrxNeHtx5QM3PMvcJzrX72s73v8R1jo3153fd9yTI45")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // define timeout
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
	fmt.Println(person)
}

// func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Add("content-type", "application/json")
// 	var people []Person
// 	collection := client.Database("thepolyglotdeveloper").Collection("people")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // define timeout
// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		response.WriteHeader(http.StatusinternalServerError)
// 		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cusrsor.Next(ctx) {
// 		var person Person
// 		cursor.Decode()
// 	}
// }


// func Create
func main() {
	fmt.Println("Starting the application...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        numOfBytesWritten, err := fmt.Fprintf(w, "Hello, World!")
        if err != nil {
            fmt.Fprintln(w, "Error occured!")
        }

        fmt.Fprintln(w, numOfBytesWritten, " bytes written.")
        http.ListenAndServe(":80", nil)
    })
	 http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello World"))
    })
 
    http.ListenAndServe(":5000", nil)


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ = mongo.Connect(ctx, "mongodb://localhost:27017") // tutorial 방식


	// 공식문서 방식 
	// https://www.mongodb.com/blog/post/quick-start-golang-mongodb-starting-and-setup
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://YohanSunn:hgt1234!@cluster0.mq9mn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	
	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	// postman으로 테스트 하기 // https://www.mongodb.com/developer/article/data-api-postman/
	http.ListenAndServe(":12345", router)
}
