package main

import (
	"fmt"

	"github.com/vmatteus/api-go-rest/models"
	"github.com/vmatteus/api-go-rest/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{ID: 1, Name: "Vitor", History: "Vitor is a software engineer."},
		{ID: 2, Name: "Matteus", History: "Matteus is a software engineer."},
	}

	fmt.Println("Start...")
	routes.LoadRoutes()
}
