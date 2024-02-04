package business

import (
	"context"                                   // i/o context
	"github.com/ltphat2204/goauth/users/entity" // entity object
)

// Interface for get new user in storage
type getUserStorage interface {
	GetUser(ctx context.Context, condition map[string]interface{}) (*entity.UserShow, error)
}

// Business layer work with storage layer through an interface
type getUserBusiness struct {
	store getUserStorage
}

// Constructor for business layer
func GetUserByUsernameBusiness(s getUserStorage) *getUserBusiness {
	return &getUserBusiness{store: s}
}

func (biz *getUserBusiness) GetUserByUsername(ctx context.Context, username string) (*entity.UserShow, error) {
	// Get user
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"username": username})

	// Error of bellow layer(s), need not to be handle
	return user, err
}
