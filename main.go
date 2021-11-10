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
	userV1 := e.Group("/api/user")
	userV1.Use(middleware.Cor())
	userV1.POST("/register", conn.Register)
	userV1.POST("/login", conn.Login)
	userV1.GET("/Info", middleware.AuthMiddleware(), conn.Info)
	working := e.Group("/api/working")
	working.POST("/register", middleware.AuthMiddleware(), conn.AccsessReginst)
	working.GET("/info", middleware.AuthMiddleware(), conn.Accsessinfo)
	working.POST("/Personnel",middleware.AuthMiddleware(),conn.PersonnelRegistration)
	working.POST("/getPermission",middleware.AuthMiddleware(),conn.GetPermission)
	e.Run(":80")
}
