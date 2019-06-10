package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"api/db"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"*"},
    }))

	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)
	r.Run(8080)
}

func getTodos(c *gin.Context) {
	todos, err := db.Todos()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error occured",
		})
		return
	}

	c.JSON(200, gin.H{ "todos": todos })
}

func createTodo(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return	
	}

	err := db.CreateTodo(name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error occured",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func updateTodo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello",
	})
}

func deleteTodo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello",
	})
}
