package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vmatteus/api-go-rest/controllers"
)

func LoadRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personality/{id}", controllers.GetPersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
