package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"jsrdriguez/crud-go/handlers"
	"jsrdriguez/crud-go/models"
)

func main() {
	models.AutoMigrate()

	r := mux.NewRouter()

	r.HandleFunc("/entities", handlers.GetEntities).Methods("GET")
	r.HandleFunc("/entity", handlers.StoreEntity).Methods("POST")
	r.HandleFunc("/entity/{id}", handlers.GetEntity).Methods("GET")
	r.HandleFunc("/entity/{id}", handlers.DeleteEntity).Methods("DELETE")
	r.HandleFunc("/entity/{id}", handlers.UpdateEntity).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", r))
}
