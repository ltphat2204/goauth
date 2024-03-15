package business

import (
	"context"
	"log"
	"os"

	"github.com/ltphat2204/goauth/users/entity"
	"github.com/ltphat2204/goauth/users/storage"
	"gorm.io/gorm"
)

func CreateAdmin(database *gorm.DB) error {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	if adminUsername == "" {
		adminUsername = "admin"
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "admin"
	}

	s := storage.NewMySqlStorage(database)
	biz := CreateNewUserBusiness(s)

	user := entity.User {
		Username: adminUsername,
		Password: adminPassword,
	}

	err := biz.CreateNewUser(context.TODO(), &user)
	if err == entity.ErrExtUsername || err == nil {
		log.Println("Admin user already created!")
		return nil
	}
	
	return err
}