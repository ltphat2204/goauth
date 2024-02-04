package business

import (
	"context"                                   // i/o context
	"time"                                      // get time now
	"github.com/ltphat2204/goauth/users/entity" // entity object
	"golang.org/x/crypto/bcrypt"                // hashing
)

// Interface for create new user in storage
type createUserStorage interface {
	CreateUser(ctx context.Context, data *entity.User) error
	GetUser(ctx context.Context, condition map[string]interface{}) (*entity.UserShow, error)
}

// Business layer work with storage layer through an interface
type createUserBusiness struct {
	store createUserStorage
}

// Constructor for business layer
func CreateNewUserBusiness(s createUserStorage) *createUserBusiness {
	return &createUserBusiness{ store: s }
}

func (biz *createUserBusiness) CreateNewUser(ctx context.Context, data *entity.User) error {
	// Make sure no existing username
	if oldUser, _ := biz.store.GetUser(ctx, map[string]interface{}{"username": data.Username}); oldUser != nil {
		return entity.ErrExtUsername
	}

	// Hash password before pushing to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.ErrPswdCnntHash
	}

	// Assign new password
	data.Password = string(hashedPassword)

	// Assign created time and last access time
	now := time.Now()
	data.CreatedAt = &now
	data.LastAccessAt = &now

	// Create user
	err = biz.store.CreateUser(ctx, data)

	// Error of bellow layer(s), need not to be handle
	return err
}
