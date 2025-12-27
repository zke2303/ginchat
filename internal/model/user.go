package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          string    `gorm:"column:id;type:varchat(64);primaryKey;"`
	Username    string    `gorm:"column:username;type:varchar(30);index:idx_username"`
	Password    string    `gorm:"column:password;type:varchar(64);not null"`
	Gender      string    `gorm:"column:gender;type:tinyint;default:0"`
	Age         string    `gorm:"column:age;type:int;"`
	Email       string    `gorm:"column:email;type:varchar(30);not null;index:idx_email"`
	LoginTime   time.Time `gorm:"column:login_time;type:datatime"`
	SignOutTime time.Time `gorm:"column:sign_out_time;type:datatime"`
}
