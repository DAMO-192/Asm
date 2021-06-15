package conn

import (
	"Asm/databases"
	"Asm/moled"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	db := databases.InitDB()
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
	if len(passwd) > 6 {
		c.JSON(http.StatusInternalServerError, gin.H{"Msg": "密码长度不足"})
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"msg": "密码加密错误"})
		return
	}
	if isTelephoneExisit(db, telephone) {
		c.JSON(203, gin.H{"msg": "账号以存在"})
		return
	}
	var user = moled.User{
		Model:     gorm.Model{},
		Name:      name,
		Passwd:    string(pwd),
		Telephone: telephone,
	}

	db.Create(&user)
	c.JSON(200, gin.H{"msg": "注册成功"})

}
func isTelephoneExisit(db *gorm.DB, telephone string) bool {
	var user moled.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
func Login(c *gin.Context) {
	db := databases.InitDB()
	telephone := c.PostForm("telephone")
	passwd := c.PostForm("passwd")
	if len(telephone) != 11 {
		c.JSON(http.StatusInternalServerError, gin.H{"Msg": "手机号不为11位"})
		return
	}
	if len(passwd) < 6 {
		c.JSON(http.StatusInternalServerError, gin.H{"Msg": "密码长度不足"})
		return
	}
	var user moled.User
	db.Where("telephone=?", telephone).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd))
	if err != nil {
		c.JSON(203, gin.H{"msg": "密码错误"})
	}
}
