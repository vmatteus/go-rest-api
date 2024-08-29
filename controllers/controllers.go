package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vmatteus/api-go-rest/database"
	"github.com/vmatteus/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var p models.Personality
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditPersolality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personality
	database.DB.First(&p, id)

	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
