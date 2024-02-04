package common

import "time"

// Common model for most situations
type MySqlModel struct {
	CreatedAt    *time.Time `json:"created_at"     gorm:"column:created_at"`
	LastAccessAt *time.Time `json:"last_access_at" gorm:"column:last_access_at"`
}
