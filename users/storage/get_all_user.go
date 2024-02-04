package storage

import (
	"context"                                   // i/o context
	"github.com/ltphat2204/goauth/users/entity" // entity object
)

func (s *mySqlStorage) GetAllUsers(ctx context.Context, condition map[string]interface{}) ([]entity.UserShow, error) {
	var result []entity.UserShow

	if err := s.storage.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
