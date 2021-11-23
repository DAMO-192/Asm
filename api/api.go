package api

import (
	"Asm/conn"
	"Asm/middleware"
	"github.com/gin-gonic/gin"
)

func StartService() {
	e := gin.Default()
	e.Use(middleware.Cor())
	userV1 := e.Group("/api/user")
	userV1.POST("/register", conn.Register)
	userV1.POST("/login", conn.Login)
	userV1.GET("/Info", middleware.AuthMiddleware(), conn.Info)
	userV1.GET("/menulist", middleware.AuthMiddleware(), conn.MenuList)
	working := e.Group("/api/working")
	working.POST("/register", middleware.AuthMiddleware(), conn.AccsessReginst)
	working.GET("/info", middleware.AuthMiddleware(), conn.Accsessinfo)
	working.POST("/Personnel", middleware.AuthMiddleware(), conn.PersonnelRegistration)
	working.POST("/getPermission", middleware.AuthMiddleware(), conn.GetPermission)
	working.POST("/menulist", middleware.AuthMiddleware(), conn.Menulist)
	e.Run(":80")

}
