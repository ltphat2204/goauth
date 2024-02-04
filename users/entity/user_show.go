package entity

import "github.com/ltphat2204/goauth/common" // common model

// User entity
type UserShow struct {
	Username string `json:"username" gorm:"column:username;size:16;primaryKey;index:idx_username;not null"`
	common.MySqlModel
}

// Table name of entity
func (UserShow) TableName() string { return "users" }
