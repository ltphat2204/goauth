package storage

import (
	"context"                                   // i/o context
	"github.com/ltphat2204/goauth/users/entity" // entity object
)

func (s *mySqlStorage) LoginUser(ctx context.Context, condition map[string]interface{}) (*entity.User, error) {
	var user entity.User
	if err := s.storage.Where(condition).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
