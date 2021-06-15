package moled

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;"`
	Passwd    string `gorm:"size:255;not null;unique;"`
	Telephone string `gorm:"varchar(11);not null;"`
}
