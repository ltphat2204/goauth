package transport

import (
	"net/http" // http status

	"github.com/gin-gonic/gin" // request
	"github.com/ltphat2204/goauth/common"
	"github.com/ltphat2204/goauth/users/business" // business layer
	"github.com/ltphat2204/goauth/users/storage"  // storage layer
	"gorm.io/gorm"                                // interact with mysql
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
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}
		if data == nil {
			c.JSON(http.StatusNotFound, common.NewErrorResponse(
				http.StatusNotFound, 
				"User not found", 
				"Username " + username.(string) + " is not existing",
			))
			return
		}

		// Get user successfully
		c.JSON(http.StatusFound, common.NewSimpleSuccessResponse("Valid token!"))
	}
}
