package middlewares

import (
	"net/http" // http status
	"os"       // load .env
	"strings"  // upper and compare string
	"time"     // get now

	"github.com/gin-gonic/gin"     // request
	"github.com/golang-jwt/jwt/v5" // parse json web token
	"github.com/ltphat2204/goauth/common"
)

func AuthMiddleware(c *gin.Context) {
	// Get Authorization token
	tokenString := c.GetHeader("Authorization")

	// No Auth
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, common.NewErrorResponse(http.StatusUnauthorized, "Unkown user", "no token provided"))
		c.Abort()
		return
	}

	// Get Authorization method
	method := tokenString[:6]
	tokenString = tokenString[7:]

	// Validate token
	if strings.Compare(strings.ToUpper(method), "BEARER") != 0 || tokenString == "" {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("missing AUTH method, BEARER or BASIC"))
		c.Abort()
		return
	}

	// Parse token from secret key
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Validate token
	if err != nil || !token.Valid {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("invalid token"))
		c.Abort()
		return
	}

	// Claim token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("something went wrong in reading token"))
		c.Abort()
		return
	}

	// Get username
	username, exists := claims["username"].(string)
	if !exists {
		c.JSON(http.StatusBadRequest, common.NewSimpleErrorResponse("invalid token body"))
		c.Abort()
		return
	}

	// Token here has passed the verify token process
	// so no need to catch error anymore
	expire, _ := claims.GetExpirationTime()

	// Token expired
	if expire.Time.Unix() < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, common.NewErrorResponse(http.StatusUnauthorized, "Invalid token", "token expired"))
		c.Abort()
		return
	}

	// Pass username to next handler
	c.Set("username", username)
	c.Next()
}
