package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const PORT string = "4000"
const DATABASE_NAME = "talent"
const PERSON_COLLECTION_NAME = "person"
const DB_URI = "mongodb+srv://<username>:<password>@hostandoptions"

type Person struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Println("CON ERR")
		return
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("PING ERR")
		return
	}

	repo := InitPersonRepo(client)
	handler := &PersonHandler{
		Repo: repo,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Handle)
	http.ListenAndServe(":"+PORT, mux)
}

func (h *PersonHandler) Handle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	w.Header().Set("content-type", "application/json")

	if r.Method == http.MethodGet {
		h.Read(w, r)
		return
	}

	if r.Method == http.MethodPost {
		h.Create(w, r)
		return
	}

	err := errors.New("method not allowed")
	log.Println(err)
	http.Error(w, err.Error(), http.StatusMethodNotAllowed)
}

type PersonHandler struct {
	Repo *PersonRepo
}

func (h *PersonHandler) Read(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	people, err := h.Repo.Read(ctx)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal struct to JSON
	jsonResp, err := json.Marshal(people)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write to resp
	w.Write(jsonResp)
}

func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var person Person
	err = json.Unmarshal(b, &person)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Repo.Create(ctx, person)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(id))
}

type PersonRepo struct {
	Coll *mongo.Collection
}

func InitPersonRepo(client *mongo.Client) *PersonRepo {
	database := client.Database(DATABASE_NAME)
	coinCollection := database.Collection(PERSON_COLLECTION_NAME)
	return &PersonRepo{
		Coll: coinCollection,
	}
}

func (repo *PersonRepo) Read(ctx context.Context) ([]Person, error) {
	var people []Person
	cursor, err := repo.Coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &people); err != nil {
		return nil, err
	}
	return people, nil
}

func (repo *PersonRepo) Create(ctx context.Context, person Person) (string, error) {

	insertResult, err := repo.Coll.InsertOne(ctx, person)
	if err != nil {
		return "", err
	}

	// convert inserted id into string
	objectId, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		errorStr := "something gone wrong while getting inserted id"
		err = errors.New(errorStr)
		return "", err
	}
	return objectId.Hex(), nil
}
