package main

import (
	"Asm/conn"
	"Asm/databases"
	"github.com/gin-gonic/gin"
)

func main() {
	databases.InitDB()

	e := gin.Default()
	e.POST("/api/path/register", conn.Register)
	e.POST("/api/path/login", conn.Login)

	e.Run(":8080")
}
