package entity

import "github.com/ltphat2204/goauth/common" // common model

// User entity
type User struct {
	common.MySqlModel
	Username string `json:"username" gorm:"column:username;size:16;primaryKey;index:idx_username;not null"`
	Password string `json:"password" gorm:"column:password;not null"`
}

// Table name of entity
func (User) TableName() string { return "users" }

func CopyInformationFrom(b UserShow) (User) {
	var a User
	a.Username = b.Username
	a.CreatedAt = b.CreatedAt
	return a
}
