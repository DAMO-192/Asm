package moled

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;" json:"name" form:"name"`
	Passwd    string `gorm:"not null;"json:"passwd"`
	Telephone string `gorm:"varchar(11);not null;unique;"json:"telephone"`
}
type ResourceType struct {
	gorm.Model
	UserName    string `gorm:"type:varchar(20);not null;"json:"user_name"`
	NetWorkType string `gorm:"type:varchar(20);not null;"json:"net_work_type"`
	PCType      string `gorm:"type:varchar(20);not null;"json:"pc_type"`
	PCModel     string `gorm:"type:varchar(20);not null;"json:"pc_model"`
	Address     string `gorm:"type:varchar(20);"json:"address"`
	MemmorySize int64
	DiskSize    int64
}
