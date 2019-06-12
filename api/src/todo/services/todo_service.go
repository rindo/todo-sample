package todo_service

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"todo/db"
)

type createParams struct {
	Name string `json:"name"`
}

type updateParams struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
}

func Routes(route *gin.Engine) {
	todo := route.Group("/todo") {
		todo.GET("/", getTodos)
		todo.POST("/", createTodo)
		todo.PUT("/:id", updateTodo)
		todo.DELETE("/:id", deleteTodo)
	}
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
	var p createParams
	err := c.BindJSON(&p)
	if err != nil {
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
	var p updateParams
	err := c.BindJSON(&p)
	if err != nil {
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
