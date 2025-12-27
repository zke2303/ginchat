package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"column:id;type:uuid;primaryKey;"`
	Username    string    `gorm:"column:username;type:varchar(30);index:idx_username"`
	Password    string    `gorm:"column:password;type:varchar(64);not null"`
	Gender      string    `gorm:"column:gender;type:tinyint;default:0"`
	Age         string    `gorm:"column:age;type:int;"`
	Email       string    `gorm:"column:email;type:varchar(30);not null;index:idx_email"`
	LoginTime   time.Time `gorm:"column:login_time;type:datatime"`
	SignOutTime time.Time `gorm:"column:sign_out_time;type:datatime"`
}
