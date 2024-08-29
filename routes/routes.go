package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vmatteus/api-go-rest/controllers"
	"github.com/vmatteus/api-go-rest/middleware"
)

func LoadRoutes() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/personality/{id}", controllers.GetPersonality).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personality/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personality/{id}", controllers.EditPersolality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
