package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"time"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Models for API Request

type Book struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Author    *Author             `json:"Author,omitempty" bson:"Author,omitempty"`
}

type Author struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}


// API methods

func getBooks(response http.ResponseWriter, request *http.Request) {
	// set response type in header
	response.Header().Set("Content-type", "application/json")
	// create result variable
	var books[] Book
	// get collection from db
	collection := client.Database("bookdb").Collection("books")
	// create context for Incoming request and outgoing calls to servers should accept a Context.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// passing empty object to Find to return all records from collection
	cursor, err := collection.Find(ctx, bson.M{})
	// Error Handling & cutomizing error response and status code
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	// closing the change stream and the underlying cursor 
    defer cursor.Close(ctx)
	// iterating and returning the next document in a cursor 
	// decoding value from the context
	// appending in the result var
	for cursor.Next(ctx) {
		var book Book
		cursor.Decode(&book)
		books = append(books, book)
	} 
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	}
    // finally returning the response
    json.NewEncoder(response).Encode(books)
}

func getBookById(response http.ResponseWriter, request *http.Request) {
    // set response type in header
	response.Header().Set("Content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	// create result variable
	var book Book
	// get collection from db
	collection := client.Database("bookdb").Collection("books")
	// create context for Incoming request and outgoing calls to servers should accept a Context.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Book{ID: id}).Decode(&book)
	// Error Handling & cutomizing error response and status code
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
    // finally returning the response
    json.NewEncoder(response).Encode(book)
}

func addBook(response http.ResponseWriter, request *http.Request) {
  response.Header().Set("Content-type", "application/json")
  var book Book
  json.NewDecoder(request.Body).Decode(&book)
  collection := client.Database("bookdb").Collection("books")
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  result, _ := collection.InsertOne(ctx, book)
  json.NewEncoder(response).Encode(result)
}

func updateBook(response http.ResponseWriter, request *http.Request, ) {
	// set response type in header
	response.Header().Set("Content-type", "application/json")

	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var book Book
	json.NewDecoder(request.Body).Decode(&book)
	
	// get collection from db
	collection := client.Database("bookdb").Collection("books")
	// create context for Incoming request and outgoing calls to servers should accept a Context.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, Book{ID: id},  bson.M{"$set": book})
	// Error Handling & cutomizing error response and status code
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
    // finally returning the response
    json.NewEncoder(response).Encode(result)
}

func deleteBook(response http.ResponseWriter, request *http.Request) {
    // set response type in header
	response.Header().Set("Content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	// get collection from db
	collection := client.Database("bookdb").Collection("books")
	// create context for Incoming request and outgoing calls to servers should accept a Context.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, Book{ID: id})
	// Error Handling & cutomizing error response and status code
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
    // finally returning the response
    json.NewEncoder(response).Encode(result)
}

func main() {
    fmt.Println("Book Rest API Application Starting .....")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
    clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.mlfbzbp.mongodb.net/?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx,clientOptions)

	//Initializing Router
	router := mux.NewRouter()

	// Route Handlers/Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBookById).Methods("GET")
	router.HandleFunc("/api/books", addBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}