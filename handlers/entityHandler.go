package handlers

import (
	"encoding/json"
	"fmt"
	"jsrdriguez/crud-go/models"
	"jsrdriguez/crud-go/repositories"
	"jsrdriguez/crud-go/requests"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetEntities(w http.ResponseWriter, r *http.Request) {
	entities := repositories.GetEntities()

	payload, err := json.Marshal(entities)

	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}

func GetEntity(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	entity := repositories.GetEntity(id)

	payload, err := json.Marshal(entity)

	if err != nil {
		log.Println(err)
	}

	w.Write(payload)
}

func StoreEntity(w http.ResponseWriter, r *http.Request) {
	var entity requests.EntityRequestInsert

	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repositories.SaveEntity(&entity)

	fmt.Fprintf(w, "ENTITY ADD")
}

func UpdateEntity(w http.ResponseWriter, r *http.Request) {
	var entity models.Entity
	var params = mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := json.NewDecoder(r.Body).Decode(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repositories.UpdateEntity(id, &entity)

	fmt.Fprintf(w, "ENTITY UPDATE")
}

func DeleteEntity(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	repositories.DeleteEntity(id)

	fmt.Fprintf(w, "ENTITY DELETE")
}
