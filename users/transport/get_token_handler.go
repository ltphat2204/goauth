package transport

import (
	"github.com/gin-gonic/gin"                    // request
	"github.com/ltphat2204/goauth/users/business" // business layer
	"github.com/ltphat2204/goauth/users/storage"  // storage layer
	"gorm.io/gorm"                                // interact with mysql
	"net/http"                                    // http status
)

func GetTokenFromUser(db *gorm.DB) func(*gin.Context) {
	// gin handler
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()

		if !hasAuth {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "No account found!"})
			return
		}

		// Create storage
		s := storage.NewMySqlStorage(db)

		// Initialize business layer with created storage
		biz := business.GetTokenFromUserBusiness(s)

		// Can not find user
		token, err := biz.GetTokenFromUser(c.Request.Context(), username, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// Get user successfully
		c.JSON(http.StatusFound, gin.H{"token": token})
	}
}
