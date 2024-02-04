package transport

import (
	"github.com/gin-gonic/gin"                    // request
	"github.com/ltphat2204/goauth/users/business" // business layer
	"github.com/ltphat2204/goauth/users/storage"  // storage layer
	"gorm.io/gorm"                                // interact with mysql
	"net/http"                                    // http status
)

func GetAllUsers(db *gorm.DB) func(*gin.Context) {
	// gin handler
	return func(c *gin.Context) {
		// Create storage
		s := storage.NewMySqlStorage(db)

		// Initialize business layer with created storage
		biz := business.FindAllUsersBusiness(s)

		// Can not find user
		data, err := biz.FindAllUsers(c.Request.Context(), map[string]interface{}{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "No account found!"})
			return
		}

		// Find user successfully
		c.JSON(http.StatusFound, gin.H{"users": data})
	}
}
