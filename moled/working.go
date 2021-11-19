package moled

import (
	"gorm.io/gorm"
)

type Personnel struct {
	gorm.Model
	Department string `gorm:"type:varchar(20);not null;" json:"department"` //部门名称可以自定义
	Name       string `gorm:"type:varchar(20);not null;" json:"name"`
	Gender     string `gorm:"type:varchar(20);not null;" json:"gender"`
	Age        int    `gorm:"type:int;not null;" json:"age"`
}
type Menulist struct {
	Id       int64  `gorm:"primarykey;" json:"id"`
	MenuName string `gorm:"varchar(20);not null" json:"menu"`
	Url      string `gorm:"type:varchar(20);not null;" json:"url"`
}
