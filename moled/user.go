package moled

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;"`
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
	MemmorySize int64
	DiskSize    int64
}
