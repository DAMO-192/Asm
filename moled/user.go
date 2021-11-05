package moled

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;" json:"name" form:"name"`
<<<<<<< HEAD
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
=======
	Passwd    string `gorm:"not null;"`
	Telephone string `gorm:"varchar(11);not null;unique;"`
}
type ResourceType struct {
	gorm.Model
	UserName    string `gorm:"type:varchar(20);not null;"`
	NetWorkType string `gorm:"type:varchar(20);not null;"`
	PCType      string `gorm:"type:varchar(20);not null;"`
	PCModel     string `gorm:"type:varchar(20);not null;"`
	Address     string `gorm:"type:varchar(20);"`
>>>>>>> 2a9ad1e (修正router组)
	MemmorySize int64
	DiskSize    int64
}
