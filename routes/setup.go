package routes

import (
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

var router *gin.Engine

func SetupRouter() {
	router = gin.Default()
	R = router

	setupRoutes()
}

func setupRoutes() {
	createBookRoutes()
}
