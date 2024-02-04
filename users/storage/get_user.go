package storage

import (
	"context"                                   // i/o context
	"github.com/ltphat2204/goauth/users/entity" // entity object
)

func (s *mySqlStorage) GetUser(ctx context.Context, condition map[string]interface{}) (*entity.UserShow, error) {
	var user entity.UserShow

	if err := s.storage.Where(condition).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
