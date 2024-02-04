package storage

import (
	"gorm.io/gorm" // loading mysql
)

// Wrap storage in a struct for future use
type mySqlStorage struct {
	storage *gorm.DB
}

// Storage constructor
func NewMySqlStorage(db *gorm.DB) *mySqlStorage {
	return &mySqlStorage{storage: db}
}
