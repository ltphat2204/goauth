package business

import (
	"context" // i/o context
	"github.com/ltphat2204/goauth/common"
	"github.com/ltphat2204/goauth/users/entity" // entity object
	"golang.org/x/crypto/bcrypt"                // compare hashed password
	"time"
)

// Interface for get token from user in storage
type getTokenStorage interface {
	LoginUser(ctx context.Context, condition map[string]interface{}) (*entity.User, error)
	UpdateLastAccess(ctx context.Context, condition map[string]interface{}) error
}

// Business layer work with storage layer through an interface
type getTokenBusiness struct {
	store getTokenStorage
}

// Constructor for business layer
func GetTokenFromUserBusiness(s getTokenStorage) *getTokenBusiness {
	return &getTokenBusiness{store: s}
}

func (biz *getTokenBusiness) GetTokenFromUser(ctx context.Context, username string, password string) (string, error) {
	// Get user
	user, err := biz.store.LoginUser(ctx, map[string]interface{}{"username": username})

	// Error in getting user
	if err != nil {
		return "", err
	}

	// User not found
	if user == nil {
		return "", entity.ErrNotFoundUsername
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", entity.ErrPswdNotMatch
	}

	// Assign old data to user
	var userNewAccess entity.UserAccess

	// Update last access time
	now := time.Now()
	userNewAccess.LastAccessAt = &now

	// Call the storage layer to update
	updateErr := biz.store.UpdateLastAccess(ctx, map[string]interface{}{"username": username})
	if updateErr != nil {
		return "", updateErr
	}

	// Get token
	token, tokenErr := common.GenerateToken(map[string]interface{}{"username": username})

	// Error of bellow layer(s), need not to be handle
	return token, tokenErr
}
