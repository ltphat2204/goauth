package business

import (
	"context"                                   // i/o context
	"github.com/ltphat2204/goauth/users/entity" // entity object
)

// Interface for get new user in storage
type findUserStorage interface {
	GetAllUsers(ctx context.Context, condition map[string]interface{}) ([]entity.UserShow, error)
}

// Business layer work with storage layer through an interface
type findUserBusiness struct {
	store findUserStorage
}

// Constructor for business layer
func FindAllUsersBusiness(s findUserStorage) *findUserBusiness {
	return &findUserBusiness{store: s}
}

func (biz *findUserBusiness) FindAllUsers(ctx context.Context, condition map[string]interface{}) ([]entity.UserShow, error) {
	// Get user
	user, err := biz.store.GetAllUsers(ctx, condition)

	// Error of bellow layer(s), need not to be handle
	return user, err
}
