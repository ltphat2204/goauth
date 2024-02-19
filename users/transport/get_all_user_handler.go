package transport

import (
	"github.com/gin-gonic/gin"                    // request
	"github.com/ltphat2204/goauth/users/business" // business layer
	"github.com/ltphat2204/goauth/users/storage"  // storage layer
	"gorm.io/gorm"                                // interact with mysql
	"net/http"                                    // http status
	"github.com/ltphat2204/goauth/common"
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
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		// Find user successfully
		c.JSON(http.StatusFound, common.NewSimpleSuccessResponse(data))
	}
}
