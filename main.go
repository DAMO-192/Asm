package main

import (
	"Asm/conn"
	"Asm/databases"
	"Asm/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	databases.InitDB()

	e := gin.Default()
	e.POST("/api/path/register", conn.Register)
	e.POST("/api/path/login", conn.Login)
	e.GET("/api/Info", middleware.AuthMiddleware(), conn.Info)
	e.Run(":8080")
}
