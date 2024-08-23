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
	r.HandleFunc("/api/personalities", controllers.GetPersonalities)
	log.Fatal(http.ListenAndServe(":8080", r))
}
