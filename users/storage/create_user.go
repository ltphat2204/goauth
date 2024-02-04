package storage

import (
	"context" 									// i/o context
	"github.com/ltphat2204/goauth/users/entity" // entity object
)

func (s *mySqlStorage) CreateUser(ctx context.Context, data *entity.User) error {
	if err := s.storage.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
