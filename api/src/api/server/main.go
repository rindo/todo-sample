package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"api/db"
)

type createParams struct {
	Name string `json:"name"`
}

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
	r.Run()
}

func getTodos(c *gin.Context) {
	todos, err := db.Todos()

	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{ "todos": todos })
}

func createTodo(c *gin.Context) {
	var p createParams
	c.BindJSON(&p)

	if p.Name == "" {
		c.Status(400)
		return	
	}

	err := db.CreateTodo(p.Name)
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}

func updateTodo(c *gin.Context) {
	c.Status(200)
}

func deleteTodo(c *gin.Context) {
	c.Status(200)
}
