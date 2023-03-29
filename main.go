package main

import (
	"fmt"
	"project-02/controllers"
	"project-02/database"

	"github.com/gin-gonic/gin"
)

const PORT = ":3000"

func main() {
	fmt.Println("Starting server.......")

	db := database.StartDB()
	//Passing
	controllers := controllers.New(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Users
	users := r.Group("/books")
	{
		users.POST("/", controllers.CreateBook)
		users.GET("/", controllers.GetBooks)
		users.GET("/:id", controllers.GetBookById)
		users.PUT("/:id", controllers.UpdateBook)
		users.DELETE("/:id", controllers.DeleteBook)
	}

	r.Run(PORT)
}
