package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/create-user", CreateUser).Methods("POST")
	r.HandleFunc("/update-user/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/delete-user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initialMigration()
	initializeRouter()

}
