package databases

import (
	"Asm/moled"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

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
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, passwd, host, port, databases, charset)
	DB, err := gorm.Open(mysql.Open(args))
	if err != nil {
		fmt.Println("系统以退出，数据库连接异常")
		os.Exit(500)
	}
	var user moled.User
	DB.AutoMigrate(&user)
	return DB
}
