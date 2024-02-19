package transport

import (
	"net/http" // http status

	"github.com/gin-gonic/gin" // request
	"github.com/ltphat2204/goauth/common"
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
			c.JSON(http.StatusUnauthorized, common.NewErrorResponse(http.StatusUnauthorized, "Token error", "Can not validate token"))
			return
		}

		// Check username match
		if username != c.Param("username") {
			c.JSON(http.StatusUnauthorized, common.NewErrorResponse(http.StatusUnauthorized, "Token error", "No right to change another's password"))
			return
		}

		// Input data
		var data entity.UserUpdate

		// Input lacks something
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
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
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		// Get user successfully
		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(map[string]string{"message": "Password updated successfully"}))
	}
}
