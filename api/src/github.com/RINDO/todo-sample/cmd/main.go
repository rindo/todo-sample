package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/RINDO/todo-sample/pkg/db"
)

func main() {
	defer db.GetInstance().Close()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	todo := route.Group("/todo")
	{
		todo.GET("/", getTodos)
		todo.POST("/", createTodo)
		todo.PUT("/:id", updateTodo)
		todo.DELETE("/:id", deleteTodo)
	}

	r.Run()
}

func getTodos(c *gin.Context) {
	todos, err := db.GetTodos()
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, todos)
}

func createTodo(c *gin.Context) {
	var p struct {
		Name string `json:"name"`
	}
	if c.BindJSON(&p) != nil {
		c.Status(400)
		return	
	}

	err = db.CreateTodo(p.Name)
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}

func updateTodo(c *gin.Context) {
	var p struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Done bool `json:"done"`
	}
	if c.BindJSON(&p) != nil {
		c.Status(400)
	}

	err = db.UpdateTodo(p.Id, p.Name, p.Done)
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}

func deleteTodo(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		c.Status(400)
		return
	}

	id, err := strconv.Atoi(p)
	if err != nil {
		c.Status(400)
		return
	}

	err = db.DeleteTodo(id)
	if err != nil {
		c.Status(500)
		return
	}

	c.Status(200)
}
