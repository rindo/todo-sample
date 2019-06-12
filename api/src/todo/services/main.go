package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"todo/db"
	"todo/services/todo_service"
)

func main() {
	defer db.GetInstance().Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))
	
	todo_service.Routes(r)

	r.Run(8080)
}
