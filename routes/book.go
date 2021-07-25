package routes

import "api/controllers"

func createBookRoutes() {
	router.GET("/books", controllers.FindBooks)
	router.GET("/book/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
}
