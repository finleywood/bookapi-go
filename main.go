package main

import (
	"api/models"
	"api/routes"
)

func main() {
	routes.SetupRouter()

	models.ConnectDataBase()

	routes.R.Run()
}
