package business

import (
	"context" // i/o context
	"time" // get time now

	"github.com/ltphat2204/goauth/users/entity" // entity object
	"golang.org/x/crypto/bcrypt"                // hashing
)

// Interface for update new user in storage
type updateUserStorage interface {
	UpdateUser(ctx context.Context, data *entity.User) error
	GetUser(ctx context.Context, condition map[string]interface{}) (*entity.UserShow, error)
}

// Business layer work with storage layer through an interface
type updateUserBusiness struct {
	store updateUserStorage
}

// Constructor for business layer
func UpdatePasswordUserBusiness(s updateUserStorage) *updateUserBusiness {
	return &updateUserBusiness{store: s}
}

func (biz *updateUserBusiness) UpdatePasswordUser(ctx context.Context, username string, password string) error {
	// Find user
	user, err := biz.store.GetUser(ctx, map[string]interface{}{"username": username})

	// Error of below layer(s), no need to handle
	if err != nil {
		return err
	}

	// User not found
	if user == nil {
		return entity.ErrNotFoundUsername
	}

	// Hash password before pushing to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.ErrPswdCnntHash
	}

	// Assign new password to user
	var userWithNewPassword entity.User
	userWithNewPassword = userWithNewPassword.Copy(*user)
	userWithNewPassword.Password = string(hashedPassword)

	// Update last access time
	now := time.Now()
	userWithNewPassword.LastAccessAt = &now

	// Update user
	err = biz.store.UpdateUser(ctx, &userWithNewPassword)

	// Error of bellow layer(s), need not to be handle
	return err
}
