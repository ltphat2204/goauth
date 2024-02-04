package entity

import "time" // time data type

// User entity
type UserAccess struct {
	LastAccessAt *time.Time `json:"last_access_at" gorm:"column:last_access_at"`
}

// Table name of entity
func (UserAccess) TableName() string { return "users" }
