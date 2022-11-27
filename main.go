package main

import (
	"go-crud/daos"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", daos.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", daos.GetUser).Methods("GET")
	r.HandleFunc("/add-user", daos.CreateUser).Methods("POST")
	r.HandleFunc("/update-user/{id}", daos.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete-user/{id}", daos.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	InitializeRouter()
}	
