package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ltphat2204/goauth/common"
	"github.com/ltphat2204/goauth/middlewares"
	"github.com/ltphat2204/goauth/users/business"
	"github.com/ltphat2204/goauth/users/transport"
	"gorm.io/gorm"
	"log"
	"os"
)

var database *gorm.DB

func init() {
	// Get the service mode
	mode := os.Getenv("GIN_MODE")

	// Load the local environment
	if mode != "release" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("No setting in .env.local found in Debug mode!")
		}
	}

	// Connect to MySql
	database = common.ConnectDB()

	// Create `users` table in MySql
	if err := business.CreateTableUser(database); err != nil {
		log.Fatal(err.Error())
	}

	if err := business.CreateAdmin(database); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	service := gin.Default()

	service.Group("/")
	{
		service.GET("", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, transport.GetAllUsers(database))
		service.POST("", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, transport.CreateUser(database))

		service.GET("/:username", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, transport.GetUserByUsername(database))
		service.PATCH("/:username", middlewares.AuthMiddleware, transport.UpdatePasswordByUsername(database))
		service.DELETE("/:username", middlewares.AuthMiddleware, middlewares.AdminMiddleWare, transport.DeleteUser(database))

		service.GET("/token", transport.GetTokenFromUser(database))
		service.GET("/verify", middlewares.AuthMiddleware, transport.VerifyToken(database))
	}

	service.Run()
}
