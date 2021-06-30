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
	e.Use(middleware.Cor())
	e.POST("/api/path/register", conn.Register)
	e.POST("/api/path/login", conn.Login)
	e.GET("/api/Info", middleware.AuthMiddleware(), conn.Info)
	e.POST("/api/accsessreginst", middleware.AuthMiddleware(), conn.AccsessReginst)
	e.GET("/api/accsessinfo", middleware.AuthMiddleware(), conn.Accsessinfo)
	e.Run(":8081")
}
