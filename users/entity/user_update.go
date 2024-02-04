package entity

import "github.com/ltphat2204/goauth/common" // common model

// User entity
type UserUpdate struct {
	common.MySqlModel
	Password string `json:"password" gorm:"column:password;not null"`
}

// Table name of entity
func (UserUpdate) TableName() string { return "users" }
