package storage

import (
	"context" // i/o context

	"github.com/ltphat2204/goauth/users/entity" // entity object
)

func (s *mySqlStorage) DeleteUser(ctx context.Context, username string) error {
	user := entity.UserShow {
		Username: username,
	}
	
	err := s.storage.Table("users").Delete(&user).Error

	return err
}
