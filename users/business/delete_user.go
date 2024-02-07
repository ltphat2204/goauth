package business

import (
	"context" // i/o context
	"os" // get env
	"github.com/ltphat2204/goauth/users/entity" // get error
)

// Interface for delete new user in storage
type deleteUserStorage interface {
	DeleteUser(ctx context.Context, username string) error
}

// Business layer work with storage layer through an interface
type deleteUserBusiness struct {
	store deleteUserStorage
}

// Constructor for business layer
func DeleteUserBusiness(s deleteUserStorage) *deleteUserBusiness {
	return &deleteUserBusiness{store: s}
}

func (biz *deleteUserBusiness) DeleteUser(ctx context.Context, username string) error {
	// Prohibit from deleting admin user
	if username == os.Getenv("ADMIN_USERNAME") {
		return entity.ErrDelAdmin
	}

	// Delete user by username
	err := biz.store.DeleteUser(ctx, username)

	// Error of bellow layer(s), need not to be handle
	return err
}
