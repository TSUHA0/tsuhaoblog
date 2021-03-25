package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"Password"`
	Role     int    `gorm:"type:int" json:"Role"`
}
