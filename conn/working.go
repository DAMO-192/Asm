package conn

import (
	"Asm/databases"
	"Asm/moled"
	"github.com/gin-gonic/gin"
)

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

//登记资源注册
func Accsessinfo(c *gin.Context) {
	db := databases.InitDB()
	var resoucetype []moled.ResourceType
	//var resourceType moled.ResourceType
	db.Find(&resoucetype)
	//c.JSON(200, gin.H{"用户": resoucetype[4], "网络类型": resoucetype[5], "电脑类型": resoucetype[6], "电脑型号": resoucetype[7], "IP地址": resoucetype[8],
	//	"内存大小": resoucetype[9], "硬盘大小": resoucetype[10],
	//})
	user, _ := c.Get("user")
	m := user.(*moled.User)
	uri := c.Request.RequestURI
	if AccStatus(m, uri) != true {
		c.JSON(403, gin.H{"MSG": "权限不足"})
		return
	}
	c.JSON(200, resoucetype)
}

//查看资源
func GetPermission(c *gin.Context) {
	s1 := c.PostForm("name")
	s2 := c.PostForm("telephone")
	s3 := c.PostForm("url")
	user, _ := c.Get("user")
	m := user.(*moled.User)
	uri := c.Request.RequestURI
	if s1 == "" || s2 == "" || s3 == "" {
		c.JSON(200, gin.H{"MSG": "所有数据不能为空"})
		return
	}
	if AccStatus(m, uri) != true {
		c.JSON(403, gin.H{"MSG": "权限不足"})
		return
	}
	CasbinGetPermission(s1, s2, s3)
}

//授权账号
func PersonnelRegistration(c *gin.Context) {
	db := databases.InitDB()
	var Personnel moled.Personnel
	c.BindJSON(&Personnel)
	age := Personnel.Age
	name := Personnel.Name
	department := Personnel.Department
	gender := Personnel.Gender
	if age <= 0 || age > 100 {
		c.JSON(200, gin.H{"MSG": "年龄填写错误"})
		return
	}
	if name == "" {
		c.JSON(200, gin.H{"MSG": "姓名不能为空"})
		return
	}
	if department == "" {
		c.JSON(200, gin.H{"MSG": "部门不能为空"})
		return
	}
	if gender == "" && gender != "男" && gender != "女" {
		c.JSON(200, gin.H{"MSG": "性别只能为男女"})
		return
	}

	user, _ := c.Get("user")
	m := user.(*moled.User)
	uri := c.Request.RequestURI
	if AccStatus(m, uri) != true {
		c.JSON(403, gin.H{"MSG": "权限不足"})
		return
	}
	db.Create(&Personnel)
	c.JSON(200, gin.H{"MSG": "登记成功"})
}

//人员登记
func Menulist(c *gin.Context) {
	db := databases.InitDB()
	var menulist moled.Menulist
	c.BindJSON(&menulist)
	name := menulist.MenuName
	url := menulist.Url
	if name == " " || url == "" {
		c.JSON(200, gin.H{"msg": "数据不能为空"})
		return
	}
	db.Create(&menulist)
	c.JSON(200, gin.H{"msg": "成功"})
}
