package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"column:id;type:varchar(36);primaryKey;"`
	Username    string     `gorm:"column:username;type:varchar(30);index:idx_username"`
	Password    string     `gorm:"column:password;type:varchar(64);not null" json:"-"`
	Gender      uint8      `gorm:"column:gender;type:tinyint;default:0"`
	Age         uint8      `gorm:"column:age;type:int;"`
	Email       string     `gorm:"column:email;type:varchar(30);not null;index:idx_email"`
	LoginTime   *time.Time `gorm:"column:login_time;type:datetime"`
	SignOutTime *time.Time `gorm:"column:sign_out_time;type:datetime"`
}

func (User) TableName() string {
	return "users"
}
