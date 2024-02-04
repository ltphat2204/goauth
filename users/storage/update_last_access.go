package storage

import (
	"context" // i/o context
	"time"
)

func (s *mySqlStorage) UpdateLastAccess(ctx context.Context, condition map[string]interface{}) error {
	if err := s.storage.Table("users").Where(condition).Update("last_access_at", time.Now().UTC()).Error; err != nil {
		return err
	}

	return nil
}
