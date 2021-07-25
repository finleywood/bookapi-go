package main

import (
	"api/models"
	"api/routes"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func main() {
	//godotenv.Load()

	gin.SetMode(gin.ReleaseMode)

	routes.SetupRouter()

	models.ConnectDataBase()

	routes.R.Run()
}
