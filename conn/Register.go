package conn

import (
	"Asm/databases"
	"Asm/dto"
	"Asm/moled"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	db := databases.InitDB()
	var req moled.User
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	//name:=req.Name
	//telephone := req.Telephone
	c.ShouldBind(&req)
	passwd := c.PostForm("passwd")

	if len(name) == 0 {
		c.JSON(http.StatusOK, gin.H{"Msg": "请输入用户名"})
		return
	}
	if len(telephone) != 11 {
		c.JSON(http.StatusOK, gin.H{"Msg": "手机号不为11位"})
		return
	}
	if len(passwd) < 6 {
		c.JSON(http.StatusOK, gin.H{"Msg": "密码长度不足"})
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"msg": "密码加密错误"})
		return
	}
	if isTelephoneExisit(db, telephone) {
		c.JSON(200, gin.H{"msg": "账号以存在"})
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
	_ = db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
func Login(c *gin.Context) {
	db := databases.InitDB()
	var req = moled.User{}
	c.Bind(&req)
	telephone := req.Telephone
	passwd := req.Passwd
	if len(telephone) != 11 {
		c.JSON(204, gin.H{"Msg": "手机号不为11位"})
		return
	}
	if len(passwd) < 6 {
		c.JSON(204, gin.H{"Msg": "密码长度不足"})
		return
	}
	var user moled.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(203, gin.H{"Msg": "密码错误"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd)); err != nil {
		c.JSON(203, gin.H{"Msg": "密码错误"})
		return
	}
	token, e := tokenst(user)
	if e != nil {
		c.JSON(303, "系统异常")
		return
	}

	c.JSON(200, gin.H{"Msg": "登陆成功",
		"token": token,
		"data":  " ",
	})

}
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{"msg": dto.ToUserDto(user.(*moled.User))})
}
func AccsessReginst(c *gin.Context) {
	db := databases.InitDB()
	var pcty moled.ResourceType
	c.Bind(&pcty)
	username := pcty.UserName
	networktype := pcty.NetWorkType
	pcytpe := pcty.PCType
	pcmoled := pcty.PCModel
	ip := pcty.Address
	disksize := pcty.DiskSize
	memmorysize := pcty.MemmorySize
	if len(username) == 0 || len(networktype) == 0 || len(pcytpe) == 0 || len(pcmoled) == 0 {
		c.JSON(200, gin.H{
			"user_name":     "",
			"net_work_type": "",
			"pc_type":       "",
			"pc_model":      "",
			"address":       "",
			"MSG":           "请根据相关类型填写",
		})

		return
	}
	if networktype != "WAN" && networktype != "LAN" {

		c.JSON(200, gin.H{"msg": "网络类型只能为LAN或者WAN", "c": networktype})
		return
	}
	if pcytpe != "台式" && pcytpe != "笔记本" && pcytpe != "云桌面" {
		c.JSON(200, gin.H{"msg": "计算机类型只能为 台式 笔记本 云桌面"})
		return
	}
	var resourceType = moled.ResourceType{
		UserName:    username,
		PCModel:     pcmoled,
		PCType:      pcytpe,
		NetWorkType: networktype,
		Address:     ip,
		DiskSize:    disksize,
		MemmorySize: memmorysize,
	}
	db.Create(&resourceType)
	c.JSON(200, gin.H{"msg": "资源登记成功"})
}
func Accsessinfo(c *gin.Context) {
	db := databases.InitDB()
	var resoucetype []moled.ResourceType
	//var resourceType moled.ResourceType
	db.Find(&resoucetype)
	//c.JSON(200, gin.H{"用户": resoucetype[4], "网络类型": resoucetype[5], "电脑类型": resoucetype[6], "电脑型号": resoucetype[7], "IP地址": resoucetype[8],
	//	"内存大小": resoucetype[9], "硬盘大小": resoucetype[10],
	//})
	c.JSON(200, resoucetype)
}
