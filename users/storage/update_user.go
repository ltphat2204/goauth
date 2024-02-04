package storage

import (
	"context" // i/o context
	"time"

	"github.com/ltphat2204/goauth/users/entity" // entity object
)

func (s *mySqlStorage) UpdateUser(ctx context.Context, data *entity.User) error {
	if err := s.storage.Table("users").Where("username = ?", data.Username).Update("password", data.Password).Update("last_access_at", time.Now().UTC()).Error; err != nil {
		return err
	}

	return nil
}
