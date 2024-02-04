package transport

import (
	"github.com/gin-gonic/gin"                    // request
	"github.com/ltphat2204/goauth/users/business" // business layer
	"github.com/ltphat2204/goauth/users/storage"  // storage layer
	"gorm.io/gorm"                                // interact with mysql
	"net/http"                                    // http status
)

func VerifyToken(db *gorm.DB) func(*gin.Context) {
	// gin handler
	return func(c *gin.Context) {
		// Get username
		// Because this middleware come after the auth
		// We do not need to check its existing
		username, _ := c.Get("username")

		// Create storage
		s := storage.NewMySqlStorage(db)

		// Initialize business layer with created storage
		biz := business.GetUserByUsernameBusiness(s)

		// Can not find user
		data, err := biz.GetUserByUsername(c.Request.Context(), username.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found!"})
			return
		}

		// Get user successfully
		c.JSON(http.StatusFound, gin.H{"message": "Valid token!"})
	}
}
