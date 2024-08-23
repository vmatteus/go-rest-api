package routes

import (
	"log"
	"net/http"

	"github.com/vmatteus/api-go-rest/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
