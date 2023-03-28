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
		users.POST("/create", controllers.CreateBook)
		users.GET("/", controllers.GetBooks)
		users.GET("/:id", controllers.GetBookById)
		users.PUT("/update/:id", controllers.UpdateBook)
		users.DELETE("/delete/:id", controllers.DeleteBook)
	}

	r.Run(PORT)
}
