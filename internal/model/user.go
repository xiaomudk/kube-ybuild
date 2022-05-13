package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// UserModel User represents a registered user.
type UserModel struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Username  string    `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password  string    `json:"-" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Phone     int64     `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Sex       int       `gorm:"column:sex" json:"sex"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// UserInfo 对外暴露的结构体
type UserInfo struct {
	ID       uint64 `json:"id" example:"1"`
	Username string `json:"username" example:"张三"`
	Sex      int    `json:"sex"`
}

// TableName 表名
func (u *UserModel) TableName() string {
	return "user"
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
