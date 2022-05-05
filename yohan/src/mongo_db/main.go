package main


import (
	"context"
	"fmt"
	// "os"
	"time"
	"log"
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type Person struct {
	ID 		  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string 		     `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string 			 `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client


func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	database := client.Database("techcody")
	peopleCollection := database.Collection("training")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // define timeout
	result, _ := peopleCollection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
	fmt.Println(person)
}

func main() {
	fmt.Println("Starting the application...")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     numOfBytesWritten, err := fmt.Fprintf(w, "Hello, World!")
    //     if err != nil {
    //         fmt.Fprintln(w, "Error occured!")
    //     }

    //     fmt.Fprintln(w, " / " , numOfBytesWritten, " bytes written..!!!")
    //     http.ListenAndServe(":80", nil)
    // })
	
	// http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
    //     w.Write([]byte("Hello World..!!!"))
    // })
    // http.ListenAndServe(":5000", nil)


	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://YohanSunn:hgt1234!@cluster0.mq9mn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	// podcast := Podcast{
    // 	Title:  "Lets develop Go server",
    // 	Author: "Yohan Sun",
    // 	Tags:   []string{"development", "programming", "coding"},
	// }

	
	
	// database := client.Database("techcody")
	// podcastsCollection := database.Collection("training")

	// insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println(insertResult.InsertedID)

	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	http.ListenAndServe(":8080", router)
	// postman으로 테스트 하기 // https://www.mongodb.com/developer/article/data-api-postman/

	// [5.5] router, postman 관련된게 안된다. atlas에 데이터 insert하는것 까지는 된다.  

}
