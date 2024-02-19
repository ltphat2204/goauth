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

func CreateUser(db *gorm.DB) func(*gin.Context) {
	// gin handler
	return func(c *gin.Context) {
		// Input data
		var data entity.User

		// Input lacks something
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		// Create storage
		s := storage.NewMySqlStorage(db)

		// Initialize business layer with created storage
		biz := business.CreateNewUserBusiness(s)

		// Get user token
		token, err := common.GenerateToken(map[string]interface{}{"username": data.Username})
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		// Can not create new user
		if err := biz.CreateNewUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse(err.Error()))
			return
		}

		// Create new user successfully
		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(map[string]string{
			"message": "User " + data.Username + " created successfully!", 
			"token": token,
		}))
	}
}
