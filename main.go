package main

import (
	"fmt"

	"github.com/vmatteus/api-go-rest/routes"
)

func main() {
	fmt.Println("Start...")
	routes.LoadRoutes()
}
