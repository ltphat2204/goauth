package middlewares

import (
	"github.com/gin-gonic/gin" // request
	"os"                       // load .env
)

func AdminMiddleWare(c *gin.Context) {
	// Get username
	// Because this middleware come after the auth
	// We do not need to check its existing
	username, _ := c.Get("username")

	// Get admin username
	defaultAdminUsername := os.Getenv("ADMIN_USERNAME")
	if defaultAdminUsername == "" {
		defaultAdminUsername = "admin" //admin default username
	}

	// Check if admin is connecting
	if username == defaultAdminUsername {
		c.Next()
		return
	}

	c.Abort()
}
