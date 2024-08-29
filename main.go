package main

import (
	"fmt"

	"github.com/vmatteus/api-go-rest/database"
	"github.com/vmatteus/api-go-rest/routes"
)

func main() {
	database.Connect()
	fmt.Println("Start...")
	routes.LoadRoutes()
}
