package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;"`
	Passwd    string `gorm:"size:255;not null;unique;"`
	Telephone string `gorm:"varchar(11);not null;"`
}

func main() {
	db := InitDB()
	e := gin.Default()
	e.POST("/api/path/register", func(c *gin.Context) {
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		passwd := c.PostForm("password")

		if len(name) == 0 {
			c.JSON(200, gin.H{"Msg": "请输入用户名"})
			return
		}
		if len(telephone) != 11 {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "手机号不为11位"})
			return
		}
		if len(passwd) < 6 {
			c.JSON(http.StatusInternalServerError, gin.H{"Msg": "密码长度不足"})
			return
		}
		var user = User{
			Model:     gorm.Model{},
			Name:      name,
			Passwd:    passwd,
			Telephone: telephone,
		}
		db.AutoMigrate(&user)
		db.Create(&user)

	})
	e.Run(":8080")
}
func InitDB() *gorm.DB {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("conf")
	v.AddConfigPath(".")
	v.ReadInConfig()
	host := v.Get("db.host")
	username := v.Get("db.username")
	port := v.Get("db.port")
	passwd := v.Get("db.passwd")
	charset := v.Get("db.charset")
	databases := v.Get("db.databases")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, passwd, host, port, databases, charset)
	DB, err := gorm.Open(mysql.Open(args))
	if err != nil {
		fmt.Println("系统以退出，数据库连接异常")
		os.Exit(500)
	}
	return DB
}
