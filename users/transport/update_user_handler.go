package transport

import (
	"net/http" // http status

	"github.com/gin-gonic/gin"                    // request
	"github.com/ltphat2204/goauth/users/business" // business layer
	"github.com/ltphat2204/goauth/users/entity"   // entity layer
	"github.com/ltphat2204/goauth/users/storage"  // storage layer
	"gorm.io/gorm"                                // interact with mysql
)

func UpdatePasswordByUsername(db *gorm.DB) func(*gin.Context) {
	// gin handler
	return func(c *gin.Context) {
		// Get username
		username, usernameExist := c.Get("username")

		if !usernameExist {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Error in verify token!"})
			return
		}

		// Check username match
		if username != c.Param("username") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No right to change another's password"})
			return
		}

		// Input data
		var data entity.UserUpdate

		// Input lacks something
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// Get password
		password := data.Password

		// Create storage
		s := storage.NewMySqlStorage(db)

		// Initialize business layer with created storage
		biz := business.UpdatePasswordUserBusiness(s)

		// Update password
		err := biz.UpdatePasswordUser(c.Request.Context(), username.(string), password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// Get user successfully
		c.JSON(http.StatusOK, gin.H{"message": "Password update successfully!"})
	}
}
