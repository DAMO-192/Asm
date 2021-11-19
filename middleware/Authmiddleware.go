package middleware

import (
	"Asm/conn"
	"Asm/databases"
	"Asm/moled"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenstring := c.GetHeader("Authorization")
		if tokenstring == "" || !strings.HasPrefix(tokenstring, "Bearer") {
			c.JSON(200, gin.H{"msg": "权限不足1"})
			c.Abort()
			return
		}
		tokenstring = tokenstring[7:]
		tokenjs, claims, err := conn.Tokenjs(tokenstring)
		if err != nil || !tokenjs.Valid {
			c.JSON(200, gin.H{"msg": "权限不足2"})
			c.Abort()
			return
		}
		userid := claims.Userid
		db := databases.InitDB()
		var user moled.User
		db.First(&user, userid)
		if user.ID == 0 {
			c.JSON(200, gin.H{"code": "401", "msg": "权限不足3"})
			c.Abort()
			return
		}
		c.Set("user", &user)
		c.Next()
	}
}
